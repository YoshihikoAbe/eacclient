package eacnet

import (
	"github.com/YoshihikoAbe/avsproperty"
	"github.com/YoshihikoAbe/eacclient/internal"
)

const InfinitasRoot = "p2d"

type InfinitasConfig struct {
	Version string
	ID      string
	Token   string
}

type LowLevelInfinitasClient struct {
	InfinitasConfig
	baseClient
}

func NewLowLevelInfinitasClient(config InfinitasConfig) *LowLevelInfinitasClient {
	cl := &LowLevelInfinitasClient{
		InfinitasConfig: config,
	}
	cl.baseClient.LowLevelClient = cl
	cl.baseClient.HTTP.Transport = transport
	cl.baseClient.proto = InfinitasProtocol
	return cl
}

func (cl *LowLevelInfinitasClient) preprocess(root *avsproperty.Node, form map[string]string) error {
	if err := internal.CompareNodeName(InfinitasRoot, root); err != nil {
		return err
	}

	form["client_version"] = cl.Version
	if cl.ID != "" {
		form["infinitas_id"] = cl.ID
	}
	if cl.Token != "" {
		form["p2d_token"] = cl.Token
	}
	form["retry_count"] = "0"
	return nil
}
