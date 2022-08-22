package main

import (
	"os"

	"github.com/ryancswallace/gaar/gaar"
)

func main() {
	gaar.SetUsage()

	dispConf := gaar.ParseCmd()

	code, err := gaar.Run(dispConf)
	if err != nil {
		os.Exit(code)
	}
}
