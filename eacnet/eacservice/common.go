package eacservice

import (
	"context"

	"github.com/YoshihikoAbe/eacclient/eacnet"
)

type CommonServiceURLs struct {
	GetServices    string
	GetServerState string
	GetServerClock string
}

var (
	KonasuteCommonServices = &CommonServiceURLs{
		GetServices:    konasuteUrl + "pre_process/GetServices",
		GetServerState: konasuteUrl + "pre_process/GetServerStatus",
		GetServerClock: konasuteUrl + "pre_process/GetServerClock",
	}

	InfinitasCommonServices = &CommonServiceURLs{
		GetServices:    infinitasUrl + "preProcess/GetServices",
		GetServerState: infinitasUrl + "preProcess/GetServerState",
		GetServerClock: infinitasUrl + "preProcess/GetServerClock",
	}
)

type CommonServiceClient struct {
	URLs     *CommonServiceURLs
	lowLevel eacnet.LowLevelClient
}

func NewCommonServiceClient(client eacnet.LowLevelClient, urls *CommonServiceURLs) *CommonServiceClient {
	return &CommonServiceClient{
		URLs:     urls,
		lowLevel: client,
	}
}

type ServiceEntry struct {
	URL  string
	Name string
}

func (cl *CommonServiceClient) GetServices(ctx context.Context) ([]ServiceEntry, error) {
	root, _ := cl.lowLevel.MakeRequestProperty("getServices")
	resp, err := cl.lowLevel.Send(ctx, cl.URLs.GetServices, root)
	if err != nil {
		return nil, err
	}

	serviceNodes := resp.SearchChildren("service")
	services := make([]ServiceEntry, len(serviceNodes))
	for i, serviceNode := range serviceNodes {
		services[i].Name, _ = serviceNode.ChildValue("service_name").(string)
		services[i].URL, _ = serviceNode.ChildValue("url").(string)
	}
	return services, nil
}

type ServerState struct {
	State       int32
	MainteStart uint64
	MainteEnd   uint64
}

func (cl *CommonServiceClient) GetServerState(ctx context.Context) (*ServerState, error) {
	root, _ := cl.lowLevel.MakeRequestProperty("getServerState")
	resp, err := cl.lowLevel.Send(ctx, cl.URLs.GetServerState, root)
	if err != nil {
		return nil, err
	}

	state := &ServerState{}
	state.State, _ = resp.ChildValue("server_state").(int32)
	state.MainteStart, _ = resp.ChildValue("mainte_start_clock").(uint64)
	state.MainteEnd, _ = resp.ChildValue("mainte_end_clock").(uint64)
	return state, nil
}

type ServerClock uint64

func (cl *CommonServiceClient) GetServerClock(ctx context.Context) (ServerClock, error) {
	root, _ := cl.lowLevel.MakeRequestProperty("getServerClock")
	resp, err := cl.lowLevel.Send(ctx, cl.URLs.GetServerClock, root)
	if err != nil {
		return 0, err
	}

	clock, _ := resp.ChildValue("server_clock").(uint64)
	return ServerClock(clock), nil
}
