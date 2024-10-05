package eacnet

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/YoshihikoAbe/avslz"
	"github.com/YoshihikoAbe/avsproperty"
	"github.com/YoshihikoAbe/eacclient/internal"
)

type Protocol struct {
	version            string
	clientSalt         Salt
	serverSalt         Salt
	base64             *base64.Encoding
	responseHeaderSize int
	errorNodeName      string
	requestNodeName    string
	resultNodeName     string
	dataNodeName       string
	rootNodeName       string
}

var (
	KonasuteProtocol = &Protocol{
		version:            "2020090800",
		clientSalt:         KonasuteClientSalt,
		serverSalt:         KonasuteServerSalt,
		base64:             base64.RawURLEncoding,
		responseHeaderSize: 42,
		errorNodeName:      "error_code",
		requestNodeName:    "request",
		resultNodeName:     "response",
		dataNodeName:       "data",
		rootNodeName:       KonasuteRoot,
	}

	InfinitasProtocol = &Protocol{
		version:            "P2D:2015091800",
		clientSalt:         InfinitasClientSalt,
		serverSalt:         InfinitasServerSalt,
		base64:             base64.StdEncoding,
		responseHeaderSize: 46,
		errorNodeName:      "error",
		requestNodeName:    "",
		resultNodeName:     "result",
		dataNodeName:       "params",
		rootNodeName:       InfinitasRoot,
	}
)

func (p *Protocol) Version() string {
	return p.version
}

func (p *Protocol) ClientSalt() Salt {
	return p.clientSalt
}

func (p *Protocol) ServerSalt() Salt {
	return p.serverSalt
}

func (p *Protocol) Base64() *base64.Encoding {
	return p.base64
}

func (p *Protocol) ResponseHeaderSize() int {
	return p.responseHeaderSize
}

func (p *Protocol) ErrorNodeName() string {
	return p.errorNodeName
}

func (p *Protocol) RequestNodeName() string {
	return p.requestNodeName
}

func (p *Protocol) ResultNodeName() string {
	return p.resultNodeName
}

func (p *Protocol) DataNodeName() string {
	return p.dataNodeName
}

func (p *Protocol) RootNodeName() string {
	return p.rootNodeName
}

func (p *Protocol) BuildForm(out *bytes.Buffer, prop *avsproperty.Property, form map[string]string) (err error) {
	lz := avslz.NewWriter(out)
	if err = prop.Write(lz); err != nil {
		return
	}
	lz.Close()
	data := out.Bytes()

	form["protocol_version"] = p.version
	form["signature"] = hex.EncodeToString(p.clientSalt.Sign(data))
	form["request"] = p.base64.EncodeToString(data)
	out.Reset()
	buildForm(out, form)
	return
}

type StatusCodeError int

func (err StatusCodeError) Error() string {
	return "status code: " + strconv.Itoa(int(err))
}

type ErrorCodeError int

func (err ErrorCodeError) Error() string {
	return "error code: " + strconv.Itoa(int(err))
}

var errMissingResult = errors.New("missing response/result node")

func (p *Protocol) ReadResponse(rd io.Reader, prop *avsproperty.Property) (result *avsproperty.Node, err error) {
	vlen := len(p.version)
	head := make([]byte, p.responseHeaderSize)
	if _, err = io.ReadFull(rd, head); err != nil {
		return
	}
	if got := head[:vlen]; string(got) != p.version {
		return nil, fmt.Errorf("invalid protocol version: %v != %v", got, []byte(p.version))
	}

	if err = prop.Read(avslz.NewReader(rd)); err != nil {
		return
	}
	root := prop.Root
	if err = internal.CompareNodeName(p.rootNodeName, root); err != nil {
		return
	}

	// for StatusCodeError and ErrorCodeError, we return the result node
	// as well as the error
	if code, _ := root.ChildValue("status").(int32); code != 0 {
		err = StatusCodeError(code)
	}
	if code, _ := root.ChildValue(p.errorNodeName).(int32); code != 0 {
		err = errors.Join(err, ErrorCodeError(code))
	}
	if result = root.SearchChild(p.resultNodeName); result == nil && err == nil {
		return nil, errMissingResult
	}
	return
}

func buildForm(out *bytes.Buffer, form map[string]string) {
	first := true
	for k, v := range form {
		if first {
			first = false
		} else {
			out.WriteByte('&')
		}
		out.WriteString(k)
		out.WriteByte('=')
		out.WriteString(v)
	}
}
