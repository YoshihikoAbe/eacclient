package eacservice

import (
	"context"

	"github.com/YoshihikoAbe/avsproperty"
	"github.com/YoshihikoAbe/eacclient/eacnet"
)

type InfinitasServiceURLs struct {
	GetResourceInfo string
}

const infinitasUrl = "https://p.eagate.573.jp/game/eac2dx/infinitas/APIEX/servletex2/"

var defaultInfinitasServices = &InfinitasServiceURLs{
	GetResourceInfo: infinitasUrl + "needlessAuth/GetFile",
}

type InfinitasServiceClient struct {
	URLs     *InfinitasServiceURLs
	lowLevel *eacnet.LowLevelInfinitasClient
}

func NewInfinitasServiceClient(client *eacnet.LowLevelInfinitasClient) *InfinitasServiceClient {
	return &InfinitasServiceClient{
		URLs:     defaultInfinitasServices,
		lowLevel: client,
	}
}

type InfinitasResourceInfo struct {
	Hash string
	Node *avsproperty.Node
}

func (cl *InfinitasServiceClient) GetResourceInfo(ctx context.Context, full bool) (*InfinitasResourceInfo, error) {
	root, data := cl.lowLevel.MakeRequestProperty("getResourceInfo")
	data.NewNodeWithValue("full", avsproperty.BoolValue(full))
	resp, err := cl.lowLevel.Send(ctx, cl.URLs.GetResourceInfo, root)
	if err != nil {
		return nil, err
	}

	ri := &InfinitasResourceInfo{}
	ri.Hash, _ = resp.ChildValue("hash").(string)
	ri.Node = resp.SearchChild("resource_info")
	return ri, nil
}
