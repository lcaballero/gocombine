package main

import (
	"gocombine/cli"
	"os"
)

func main() {
	cli.NewCli().Run(os.Args)
}
