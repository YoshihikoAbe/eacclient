package infinitas

import (
	"github.com/YoshihikoAbe/eacclient/eacclient/cmd"
	"github.com/YoshihikoAbe/eacclient/eacnet"
	"github.com/YoshihikoAbe/eacclient/eacnet/eacservice"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var client *eacservice.InfinitasServiceClient

var infinitasCmd = &cobra.Command{
	Use:   "infinitas",
	Short: "Infinitas service client",

	PersistentPreRun: func(c *cobra.Command, args []string) {
		client = eacservice.NewInfinitasServiceClient(eacnet.NewLowLevelInfinitasClient(eacnet.InfinitasConfig{
			ID:      cmd.GlobalFlags.InfinitasID,
			Version: cmd.GlobalFlags.Version,
			Token:   cmd.GlobalFlags.Token,
		}))
	},
}

func init() {
	cmd.RootCmd.AddCommand(infinitasCmd)
	infinitasCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		cmd.InheritedFlags().VisitAll(func(f *pflag.Flag) {
			if f.Name == "game" || f.Name == "client" {
				f.Hidden = true
			}
		})
		cmd.Parent().HelpFunc()(cmd, args)
	})
}
