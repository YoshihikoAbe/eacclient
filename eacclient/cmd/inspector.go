package cmd

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/YoshihikoAbe/eacclient/eacnet"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var inspectorFlags struct {
	filename string
	addr     string
	upstream string
	replace  string
}

var inspectorCmd = &cobra.Command{
	Use:   "inspector",
	Short: "Start the traffic inspector",
	Run: func(c *cobra.Command, args []string) {
		upstream, err := url.Parse(inspectorFlags.upstream)
		if err != nil {
			Fatal("parse upstream url:", err)
		}
		replace, err := url.Parse(inspectorFlags.replace)
		if err != nil {
			Fatal("parse replacement url:", err)
		}

		var proto *eacnet.Protocol
		name := strings.ToLower(GlobalFlags.Protocol)
		if strings.EqualFold(name, "konasute") {
			proto = eacnet.KonasuteProtocol
		} else if strings.EqualFold(name, "infinitas") {
			proto = eacnet.InfinitasProtocol
		} else {
			Fatal("invalid protocol:", name)
		}

		out := os.Stdout
		if name := inspectorFlags.filename; name != "-" {
			f, err := os.Create(name)
			if err != nil {
				Fatal(err)
			}
			out = f
		}

		log.Fatalln(http.ListenAndServe(inspectorFlags.addr, eacnet.NewInspectorReverseProxy(upstream, replace, proto, out)))
	},
}

func init() {
	RootCmd.AddCommand(inspectorCmd)
	flags := inspectorCmd.Flags()
	inspectorCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		cmd.InheritedFlags().VisitAll(func(f *pflag.Flag) {
			if f.Name != "protocol" {
				// hide global client flags
				f.Hidden = true
			}
		})
		cmd.Parent().HelpFunc()(cmd, args)
	})
	flags.StringVar(&inspectorFlags.filename, "out", "-", "Output filename")
	flags.StringVar(&inspectorFlags.addr, "addr", "127.0.0.1:80", "Listen address")
	flags.StringVar(&inspectorFlags.upstream, "upstream", "https://p.eagate.573.jp/", "URL of upstream server")
	flags.StringVar(&inspectorFlags.replace, "replace", "http://127.0.0.1/", "Service replacement URL")
}
