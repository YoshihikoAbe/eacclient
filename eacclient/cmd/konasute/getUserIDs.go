package konasute

import (
	"fmt"

	"github.com/YoshihikoAbe/eacclient/eacclient/cmd"
	"github.com/spf13/cobra"
)

var getUserIDsCmd = &cobra.Command{
	Use:   "getUserIDs",
	Short: "getUserIDs service client",

	Run: func(c *cobra.Command, args []string) {
		cmd.SetServiceUrl(&client.URLs.GetUserIDs)
		ids, err := client.GetUserIDs(cmd.Context())
		if err != nil {
			cmd.Fatal(err)
		}

		fmt.Println("CardID =", ids.CardID)
		fmt.Println("RefID  =", ids.RefID)
		fmt.Println("DataID =", ids.DataID)
		fmt.Println("DataID =", ids.DataID)
		fmt.Println("SnsID  =", ids.SnsID)
	},
}

func init() {
	konasuteCmd.AddCommand(getUserIDsCmd)
}
