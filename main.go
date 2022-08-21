package main

import (
	"os"

	"github.com/ryancswallace/gaar/cmd"
)

func main() {
	cmd.SetUsage()

	dispConf := cmd.ParseCmd()

	code, err := cmd.Run(dispConf)
	if err != nil {
		os.Exit(code)
	}
}
