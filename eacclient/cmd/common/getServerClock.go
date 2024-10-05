package konasute

import (
	"fmt"

	"github.com/YoshihikoAbe/eacclient/eacclient/cmd"
	"github.com/spf13/cobra"
)

var getServerClockCmd = &cobra.Command{
	Use:   "getServerClock",
	Short: "getServerClock service client",

	Run: func(c *cobra.Command, args []string) {
		cmd.SetServiceUrl(&client.URLs.GetServerClock)
		clock, err := client.GetServerClock(cmd.Context())
		if err != nil {
			cmd.Fatal(err)
		}

		fmt.Println("Clock = ", clock)
	},
}

func init() {
	commonCmd.AddCommand(getServerClockCmd)
}
