package main

import (
	"github.com/patdhlk/cli/cli"
	"os"
)

func main() {
	os.Exit(cli.Run(os.Args[1:]))
}
