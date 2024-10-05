package konasute

import (
	"fmt"

	"github.com/YoshihikoAbe/eacclient/eacclient/cmd"
	"github.com/spf13/cobra"
)

var getResourceInfoCmd = &cobra.Command{
	Use:   "getResourceInfo",
	Short: "getResourceInfo service client",

	Run: func(c *cobra.Command, args []string) {
		cmd.SetServiceUrl(&client.URLs.GetResourceInfo)
		info, err := client.GetResourceInfo(cmd.Context())
		if err != nil {
			cmd.Fatal(err)
		}

		fmt.Println("Hash =", info.Hash)
		fmt.Println("URL  =", info.URL)
	},
}

func init() {
	konasuteCmd.AddCommand(getResourceInfoCmd)
}
