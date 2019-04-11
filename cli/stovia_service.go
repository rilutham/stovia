package main

import (
	"os"
	"rilutham/stovia/cli/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
