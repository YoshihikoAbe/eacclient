package eacnet

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/YoshihikoAbe/avsproperty"
)

const UserAgent = "e-AMUSEMENT CLOUD AGENT"

// common transport with compression disabled
var transport = &http.Transport{
	DisableCompression: true,
}

type LowLevelClient interface {
	Send(ctx context.Context, url string, root *avsproperty.Node) (result *avsproperty.Node, err error)
	MakeRequestProperty(method string) (root *avsproperty.Node, data *avsproperty.Node)

	preprocess(root *avsproperty.Node, form map[string]string) error
}

type baseClient struct {
	LowLevelClient
	HTTP  http.Client
	proto *Protocol
}

func (cl *baseClient) Send(ctx context.Context, url string, root *avsproperty.Node) (result *avsproperty.Node, err error) {
	prop := &avsproperty.Property{
		Settings: avsproperty.PropertySettings{
			Encoding: avsproperty.EncodingUTF8,
		},
		Root: root.ShallowCopy(),
	}
	req, err := cl.makeRequest(ctx, url, prop)
	if err != nil {
		return nil, fmt.Errorf("request: %v", err)
	}

	resp, err := cl.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New("HTTP status: " + resp.Status)
	}

	var s StatusCodeError
	if result, err = cl.proto.ReadResponse(resp.Body, prop); err != nil && !errors.As(err, &s) {
		err = fmt.Errorf("response: %v", err)
	}
	return
}

func (cl *baseClient) MakeRequestProperty(method string) (root *avsproperty.Node, data *avsproperty.Node) {
	root, _ = avsproperty.NewNode(cl.proto.rootNodeName)
	inner := root
	if name := cl.proto.requestNodeName; name != "" {
		inner, _ = root.NewNode(name)
	}
	inner.NewNodeWithValue("method", method)
	data, _ = inner.NewNode(cl.proto.dataNodeName)
	return
}

func (cl *baseClient) makeRequest(ctx context.Context, url string, prop *avsproperty.Property) (*http.Request, error) {
	form := map[string]string{}
	if err := cl.preprocess(prop.Root, form); err != nil {
		return nil, err
	}

	out := bytes.NewBuffer(nil)
	if err := cl.proto.BuildForm(out, prop, form); err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, out)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return req, nil
}
