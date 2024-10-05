package konasute

import (
	"github.com/YoshihikoAbe/eacclient/eacclient/cmd"
	"github.com/YoshihikoAbe/eacclient/eacnet"
	"github.com/YoshihikoAbe/eacclient/eacnet/eacservice"
	"github.com/spf13/cobra"
)

var client *eacservice.CommonServiceClient

var commonCmd = &cobra.Command{
	Use:   "common",
	Short: "Common service client",

	PersistentPreRunE: func(c *cobra.Command, args []string) error {
		lowLevel := cmd.Client()
		var urls *eacservice.CommonServiceURLs
		if _, ok := lowLevel.(*eacnet.LowLevelInfinitasClient); ok {
			urls = eacservice.InfinitasCommonServices
		} else if _, ok := lowLevel.(*eacnet.LowLevelKonasuteClient); ok {
			urls = eacservice.KonasuteCommonServices
		}
		client = eacservice.NewCommonServiceClient(lowLevel, urls)
		return nil
	},
}

func init() {
	cmd.RootCmd.AddCommand(commonCmd)
}
