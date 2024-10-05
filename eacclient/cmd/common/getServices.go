package konasute

import (
	"fmt"

	"github.com/YoshihikoAbe/eacclient/eacclient/cmd"
	"github.com/spf13/cobra"
)

var getServicesCmd = &cobra.Command{
	Use:   "getServices",
	Short: "getServices service client",

	Run: func(c *cobra.Command, args []string) {
		cmd.SetServiceUrl(&client.URLs.GetServices)
		services, err := client.GetServices(cmd.Context())
		if err != nil {
			cmd.Fatal(err)
		}

		longest := 0
		for _, svc := range services {
			if len := len(svc.Name); len > longest {
				longest = len
			}
		}
		for _, svc := range services {
			fmt.Print(svc.Name)
			for i := len(svc.Name); i < longest; i++ {
				fmt.Print(" ")
			}
			fmt.Print(" = ")
			fmt.Println(svc.URL)
		}
	},
}

func init() {
	commonCmd.AddCommand(getServicesCmd)
}
