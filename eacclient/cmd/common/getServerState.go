package konasute

import (
	"fmt"

	"github.com/YoshihikoAbe/eacclient/eacclient/cmd"
	"github.com/spf13/cobra"
)

var getServerStateCmd = &cobra.Command{
	Use:   "getServerState",
	Short: "getServerState service client",

	Run: func(c *cobra.Command, args []string) {
		cmd.SetServiceUrl(&client.URLs.GetServerState)
		state, err := client.GetServerState(cmd.Context())
		if err != nil {
			cmd.Fatal(err)
		}

		fmt.Println("State             = ", state.State)
		fmt.Println("Maintenance Start = ", state.MainteStart)
		fmt.Println("Maintenance End   = ", state.MainteEnd)
	},
}

func init() {
	commonCmd.AddCommand(getServerStateCmd)
}
