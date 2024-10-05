package konasute

import (
	"os"

	"github.com/YoshihikoAbe/avsproperty"
	"github.com/YoshihikoAbe/eacclient/eacclient/cmd"
	"github.com/YoshihikoAbe/eacclient/eacnet/eacservice"
	"github.com/spf13/cobra"
)

var acRelayCmd = &cobra.Command{
	Use:   "acRelay SERVICE MODULE METHOD FILENAME",
	Short: "acRelay service client",
	Args:  cobra.MinimumNArgs(4),

	Run: func(c *cobra.Command, args []string) {
		cmd.SetServiceUrl(&client.URLs.AcRelay)

		prop := avsproperty.Property{}
		if err := prop.ReadFile(args[3]); err != nil {
			cmd.Fatal(err)
		}
		req := eacservice.AcRelayRequest{
			Service: args[0],
			Module:  args[1],
			Method:  args[2],
			Data:    prop.Root,
		}
		resp, err := client.AcRelay(cmd.Context(), req)
		if err != nil {
			cmd.Fatal(err)
		}

		prop = avsproperty.Property{
			Root: resp,
			Settings: avsproperty.PropertySettings{
				Format:   avsproperty.FormatPrettyXML,
				Encoding: avsproperty.EncodingUTF8,
			},
		}
		if err := prop.Write(os.Stdout); err != nil {
			cmd.Fatal(err)
		}
	},
}

func init() {
	konasuteCmd.AddCommand(acRelayCmd)
}
