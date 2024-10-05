package eacservice

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/YoshihikoAbe/avsproperty"
	"github.com/YoshihikoAbe/eacclient/eacnet"
	"github.com/YoshihikoAbe/eacclient/internal"
)

type KonasuteServiceURLs struct {
	GetResourceInfo string
	GetUserIDs      string
	AcRelay         string
}

const konasuteUrl = "https://p.eagate.573.jp/game/konasteapp/APIEX/servletex/"

var defaultKonasuteServices = &KonasuteServiceURLs{
	GetResourceInfo: konasuteUrl + "needless_auth/GetResourceInfo",
	GetUserIDs:      konasuteUrl + "need_auth/GetUserId",
	AcRelay:         konasuteUrl + "need_auth/AcRelay",
}

type KonasuteServiceClient struct {
	URLs     *KonasuteServiceURLs
	lowLevel *eacnet.LowLevelKonasuteClient
}

func NewKonasuteServiceClient(client *eacnet.LowLevelKonasuteClient) *KonasuteServiceClient {
	return &KonasuteServiceClient{
		URLs:     defaultKonasuteServices,
		lowLevel: client,
	}
}

type KonasuteResourceInfo struct {
	Hash string
	URL  string
}

func (cl *KonasuteServiceClient) GetResourceInfo(ctx context.Context) (*KonasuteResourceInfo, error) {
	root, _ := cl.lowLevel.MakeRequestProperty("getResourceInfo")
	resp, err := cl.lowLevel.Send(ctx, cl.URLs.GetResourceInfo, root)
	if err != nil {
		return nil, err
	}

	ri := &KonasuteResourceInfo{}
	ri.Hash, _ = resp.ChildValue("hash").(string)
	ri.URL, _ = resp.ChildValue("url").(string)
	return ri, nil
}

type UserIDs struct {
	CardID string
	RefID  string
	DataID string
	SnsID  string
}

func (cl *KonasuteServiceClient) GetUserIDs(ctx context.Context) (*UserIDs, error) {
	root, _ := cl.lowLevel.MakeRequestProperty("getUserIDs")
	resp, err := cl.lowLevel.Send(ctx, cl.URLs.GetUserIDs, root)
	if err != nil {
		return nil, err
	}

	ids := &UserIDs{}
	ids.CardID, _ = resp.ChildValue("card_num").(string)
	ids.RefID, _ = resp.ChildValue("ref_id").(string)
	ids.DataID, _ = resp.ChildValue("data_id").(string)
	ids.SnsID, _ = resp.ChildValue("sns_id").(string)
	return ids, nil
}

type AcRelayRequest struct {
	Service string
	Module  string
	Method  string
	Data    *avsproperty.Node
}

type RelayStatusError int32

func (err RelayStatusError) Error() string {
	return "relay status code: " + strconv.Itoa(int(err))
}

type RelayFaultError int32

func (err RelayFaultError) Error() string {
	return "relay fault code: " + strconv.Itoa(int(err))
}

func (cl *KonasuteServiceClient) AcRelay(ctx context.Context, arr AcRelayRequest) (*avsproperty.Node, error) {
	data := arr.Data
	if err := internal.CompareNodeName("data", data); err != nil {
		return nil, err
	}

	root, _ := avsproperty.NewNode(eacnet.KonasuteRoot)
	req, _ := root.NewNode("request")
	req.NewNodeWithValue("service", arr.Service)
	req.NewNodeWithValue("module", arr.Module)
	req.NewNodeWithValue("method", arr.Method)
	req.AppendChild(data.ShallowCopy())

	resp, err := cl.lowLevel.Send(ctx, cl.URLs.AcRelay, root)
	if err != nil {
		if resp != nil {
			root = resp.Parent()
			if status, _ := root.ChildValue("xrpc_status_code").(int32); status != 0 {
				err = errors.Join(err, RelayStatusError(status))
			} else if fault, _ := root.ChildValue("xrpc_fault_code").(int32); fault != 0 {
				err = errors.Join(err, RelayFaultError(fault))
			}
		}
		return nil, err
	}

	children := resp.Children()
	if count := len(resp.Children()); count != 1 {
		return nil, fmt.Errorf("response node contains an invalid number of children: %d != 1", count)
	}
	return children[0], nil
}
