package cmd

import (
	"os"

	"github.com/YoshihikoAbe/avsproperty"
	"github.com/spf13/cobra"
)

var sendCmd = &cobra.Command{
	Use:   "send FILENAME",
	Short: "Send a request to an eacnet server using a property file",
	Args:  cobra.MinimumNArgs(1),

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagsOneRequired("url")
	},

	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		prop := avsproperty.Property{}
		if err := prop.ReadFile(filename); err != nil {
			Fatal("read property:", err)
		}

		resp, err := Client().Send(Context(), GlobalFlags.URL, prop.Root)
		if err != nil {
			Fatal("send:", err)
		}

		prop.Settings.Format = avsproperty.FormatPrettyXML
		prop.Root = resp
		prop.Write(os.Stdout)
	},
}

func init() {
	RootCmd.AddCommand(sendCmd)
}
