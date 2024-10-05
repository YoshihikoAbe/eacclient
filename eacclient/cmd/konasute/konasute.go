package konasute

import (
	"github.com/YoshihikoAbe/eacclient/eacclient/cmd"
	"github.com/YoshihikoAbe/eacclient/eacnet"
	"github.com/YoshihikoAbe/eacclient/eacnet/eacservice"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var client *eacservice.KonasuteServiceClient

var konasuteCmd = &cobra.Command{
	Use:   "konasute",
	Short: "Konasute service client",

	PersistentPreRun: func(c *cobra.Command, args []string) {
		client = eacservice.NewKonasuteServiceClient(eacnet.NewLowLevelKonasuteClient(eacnet.KonasuteConfig{
			Game:    cmd.GlobalFlags.Game,
			Version: cmd.GlobalFlags.Version,
			Token:   cmd.GlobalFlags.Token,
		}))
	},
}

func init() {
	cmd.RootCmd.AddCommand(konasuteCmd)
	konasuteCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		cmd.InheritedFlags().VisitAll(func(f *pflag.Flag) {
			if f.Name == "iid" || f.Name == "client" {
				f.Hidden = true
			}
		})
		cmd.Parent().HelpFunc()(cmd, args)
	})
}
