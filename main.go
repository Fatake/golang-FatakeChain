package main

import (
	"os"

	"github.com/fatake/golang-FatakeChain/cli"
)

func main() {
	defer os.Exit(0)

	cmd := cli.CommandLine{}
	cmd.Run()
}
