package eacnet

import (
	"github.com/YoshihikoAbe/avsproperty"
	"github.com/YoshihikoAbe/eacclient/internal"
)

const KonasuteRoot = "eacnet"

type KonasuteConfig struct {
	Game    string
	Version string
	Token   string
}

type LowLevelKonasuteClient struct {
	KonasuteConfig
	baseClient
}

func NewLowLevelKonasuteClient(config KonasuteConfig) *LowLevelKonasuteClient {
	cl := &LowLevelKonasuteClient{
		KonasuteConfig: config,
	}
	cl.baseClient.LowLevelClient = cl
	cl.baseClient.HTTP.Transport = transport
	cl.baseClient.proto = KonasuteProtocol
	return cl
}

func (cl *LowLevelKonasuteClient) preprocess(root *avsproperty.Node, form map[string]string) error {
	if err := internal.CompareNodeName(KonasuteRoot, root); err != nil {
		return err
	}

	info := internal.NewUniqueNode(root, "info")
	// game_id and soft_version must be present, even if they're empty
	if cl.Game != "" || info.SearchChild("game_id") == nil {
		internal.SetChildValue(info, "game_id", cl.Game)
	}
	if cl.Version != "" || info.SearchChild("soft_version") == nil {
		internal.SetChildValue(info, "soft_version", cl.Version)
	}
	if cl.Token != "" {
		internal.SetChildValue(info, "token", cl.Token)
	}
	internal.SetChildValue(info, "retry_count", int32(0))

	return nil
}
