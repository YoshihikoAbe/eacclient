package infinitas

import (
	"os"

	"github.com/YoshihikoAbe/avsproperty"
	"github.com/YoshihikoAbe/eacclient/eacclient/cmd"
	"github.com/spf13/cobra"
)

var getResourceInfoCmd = &cobra.Command{
	Use:   "getResourceInfo",
	Short: "getResourceInfo service client",

	Run: func(c *cobra.Command, args []string) {
		cmd.SetServiceUrl(&client.URLs.GetResourceInfo)
		info, err := client.GetResourceInfo(cmd.Context(), true)
		if err != nil {
			cmd.Fatal(err)
		}

		prop := avsproperty.Property{
			Root: info.Node,
			Settings: avsproperty.PropertySettings{
				Format: avsproperty.FormatPrettyXML,
			},
		}
		if err := prop.Write(os.Stdout); err != nil {
			cmd.Fatal(err)
		}
	},
}

func init() {
	infinitasCmd.AddCommand(getResourceInfoCmd)
}
