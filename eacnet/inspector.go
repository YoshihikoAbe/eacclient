package eacnet

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
	"sync"

	"github.com/YoshihikoAbe/avslz"
	"github.com/YoshihikoAbe/avsproperty"
)

type inspector struct {
	upstream    *url.URL
	replacement *url.URL
	proto       *Protocol
	out         *bufio.Writer
	outMu       sync.Mutex
}

func NewInspectorReverseProxy(upstream, replacement *url.URL, proto *Protocol, out io.Writer) http.Handler {
	bio, ok := out.(*bufio.Writer)
	if !ok {
		bio = bufio.NewWriter(out)
	}

	ins := &inspector{
		replacement: replacement,
		upstream:    upstream,
		proto:       proto,
		out:         bio,
	}
	rp := &httputil.ReverseProxy{
		Rewrite:        ins.rewrite,
		ModifyResponse: ins.modifyResponse,
	}
	return rp
}

func (ins *inspector) modifyResponse(resp *http.Response) error {
	ins.outMu.Lock()
	defer ins.outMu.Unlock()

	fmt.Fprintln(ins.out, "<", resp.Proto, resp.Status)
	ins.dumpHeaders("<", resp.Header)
	if err := ins.out.Flush(); err != nil {
		return fmt.Errorf("dump response headers: %v", err)
	}

	prop, buf, err := ins.readResponseProperty(resp.Body)
	if err != nil {
		return fmt.Errorf("read response property: %v", err)
	}
	resp.Body = io.NopCloser(buf)

	if strings.HasSuffix(resp.Request.URL.String(), "GetServices") {
		if err := ins.modifyServices(resp, buf, prop); err != nil {
			return fmt.Errorf("rewrite services: %v", err)
		}
	}

	if err := ins.dumpProperty(prop); err != nil {
		return fmt.Errorf("dump response property: %v", err)
	}
	return nil
}

func (ins *inspector) readResponseProperty(body io.Reader) (*avsproperty.Property, *bytes.Buffer, error) {
	data, err := io.ReadAll(body)
	if err != nil {
		return nil, nil, err
	}
	headerSize := ins.proto.ResponseHeaderSize()
	if size := len(data); size <= headerSize {
		return nil, nil, fmt.Errorf("invalid response size: %d <= %d", size, headerSize)
	}

	prop := &avsproperty.Property{}
	lz := avslz.NewReader(bytes.NewReader(data[headerSize:]))
	if err := prop.Read(lz); err != nil {
		return nil, nil, err
	}
	return prop, bytes.NewBuffer(data), nil
}

var urlNodeName, _ = avsproperty.NewNodeName("url")

func (ins *inspector) modifyServices(resp *http.Response, buf *bytes.Buffer, prop *avsproperty.Property) error {
	prop.Root.Traverse(func(n *avsproperty.Node) error {
		if !n.Name().Equals(urlNodeName) {
			return nil
		}

		url, err := url.Parse(n.StringValue())
		if err != nil {
			log.Println(err)
			return nil
		}
		url.Host = ins.replacement.Host
		url.Scheme = ins.replacement.Scheme
		n.SetValue(url.String())
		return nil
	}, nil)

	buf.Truncate(ins.proto.responseHeaderSize)
	lz := avslz.NewWriter(buf)
	if err := prop.Write(lz); err != nil {
		return err
	}
	lz.Close()
	size := buf.Len()
	resp.ContentLength = int64(size)
	resp.Header.Set("Content-Length", strconv.Itoa(size))
	return nil
}

func (ins *inspector) rewrite(req *httputil.ProxyRequest) {
	if req.Out.Body != nil {
		if err := ins.rewriteE(req); err != nil {
			log.Println("rewrite error:", err)
		}
	}
	req.SetURL(ins.upstream)
}

func (ins *inspector) rewriteE(req *httputil.ProxyRequest) error {
	prop, rd, err := ins.readRequestProperty(req.Out.Body)
	if err != nil {
		return fmt.Errorf("read request property: %v", err)
	}
	if err := ins.dumpRequest(req.Out, prop); err != nil {
		return fmt.Errorf("dump request property: %v", err)
	}
	req.Out.Body = rd
	return nil
}

func (ins *inspector) readRequestProperty(body io.Reader) (*avsproperty.Property, io.ReadCloser, error) {
	form, _ := io.ReadAll(body)
	b64, err := formValue(form, "request")
	if err != nil {
		return nil, nil, err
	}
	raw, err := ins.proto.Base64().DecodeString(b64)
	if err != nil {
		return nil, nil, err
	}

	lz := avslz.NewReader(bytes.NewReader(raw))
	prop := &avsproperty.Property{}
	if err := prop.Read(lz); err != nil {
		return nil, nil, err
	}
	return prop, io.NopCloser(bytes.NewReader(form)), nil
}

func (ins *inspector) dumpRequest(req *http.Request, prop *avsproperty.Property) error {
	ins.outMu.Lock()
	defer ins.outMu.Unlock()

	fmt.Fprintln(ins.out, ">", req.Proto, req.Method, req.URL.Path)
	ins.dumpHeaders(">", req.Header)
	return ins.dumpProperty(prop)
}

func (ins *inspector) dumpProperty(prop *avsproperty.Property) error {
	prop.Settings.Format = avsproperty.FormatPrettyXML
	if err := prop.Write(ins.out); err != nil {
		return err
	}
	prop.Settings.Format = avsproperty.FormatBinary
	fmt.Fprintln(ins.out)
	return ins.out.Flush()
}

func (ins *inspector) dumpHeaders(direction string, headers http.Header) {
	for k, v := range headers {
		for _, v := range v {
			fmt.Fprintln(ins.out, direction, k+":", v)
		}
	}
	fmt.Fprintln(ins.out, direction)
}

func formValue(b []byte, key string) (string, error) {
	for {
		curKey, after, found := bytes.Cut(b, []byte{'='})
		if !found {
			break
		}
		before, after, found := bytes.Cut(after, []byte{'&'})
		if string(curKey) == key {
			return string(before), nil
		}
		if !found {
			break
		}
		b = after
	}
	return "", errors.New("form field \"" + key + "\" cannot be found")
}
