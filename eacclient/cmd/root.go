package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/YoshihikoAbe/eacclient/eacnet"
	"github.com/spf13/cobra"
)

var GlobalFlags struct {
	Timeout     time.Duration
	Protocol    string
	Version     string
	Token       string
	Game        string
	InfinitasID string
	URL         string
}

var RootCmd = &cobra.Command{
	Use:   "eacclient",
	Short: "eacnet clients and tools",
}

func init() {
	flags := RootCmd.PersistentFlags()
	flags.DurationVarP(&GlobalFlags.Timeout, "timeout", "", time.Second*20, "Timeout")
	flags.StringVarP(&GlobalFlags.Protocol, "protocol", "p", "konasute", "\"konasute\" or \"infinitas\"")
	flags.StringVarP(&GlobalFlags.Version, "version", "v", "", "Software version")
	flags.StringVarP(&GlobalFlags.Token, "token", "t", "", "Authentication token")
	flags.StringVarP(&GlobalFlags.Game, "game", "g", "", "Konasute game ID")
	flags.StringVarP(&GlobalFlags.InfinitasID, "iid", "i", "", "Infinitas ID")
	flags.StringVarP(&GlobalFlags.URL, "url", "u", "", "Service URL")

}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func Client() eacnet.LowLevelClient {
	name := GlobalFlags.Protocol
	if strings.EqualFold(name, "konasute") {
		return eacnet.NewLowLevelKonasuteClient(eacnet.KonasuteConfig{
			Game:    GlobalFlags.Game,
			Version: GlobalFlags.Version,
			Token:   GlobalFlags.Token,
		})
	} else if strings.EqualFold(name, "infinitas") {
		return eacnet.NewLowLevelInfinitasClient(eacnet.InfinitasConfig{
			ID:      GlobalFlags.InfinitasID,
			Version: GlobalFlags.Version,
			Token:   GlobalFlags.Token,
		})
	} else {
		Fatal("invalid client:", name)
	}
	return nil
}

func Context() (ctx context.Context) {
	ctx = context.Background()
	if to := GlobalFlags.Timeout; to > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithDeadline(ctx, time.Now().Add(to))
		_ = cancel
	}
	return
}

func SetServiceUrl(target *string) {
	if url := GlobalFlags.URL; url != "" {
		*target = url
	}
}

func Fatal(v ...any) {
	fmt.Fprintln(os.Stderr, v...)
	os.Exit(1)
}
