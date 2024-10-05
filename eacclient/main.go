package main

import (
	"github.com/YoshihikoAbe/eacclient/eacclient/cmd"
	_ "github.com/YoshihikoAbe/eacclient/eacclient/cmd/common"
	_ "github.com/YoshihikoAbe/eacclient/eacclient/cmd/infinitas"
	_ "github.com/YoshihikoAbe/eacclient/eacclient/cmd/konasute"
)

func main() {
	cmd.Execute()
}
