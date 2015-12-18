package cli

import (
	"gocombine/run"

	cmd "github.com/codegangsta/cli"
)

func NewCli() *cmd.App {
	c := cmd.NewApp()
	c.Flags = []cmd.Flag{
		cmd.StringFlag{
			Name:  "cc",
			Usage: "C/C++ code file to include when combining.",
		},
		cmd.StringFlag{
			Name:  "go",
			Usage: "Go code file to combine.",
		},
		cmd.StringFlag{
			Name:  "out",
			Usage: "Output go file produced.",
		},
		cmd.BoolFlag{
			Name:  "stdout",
			Usage: "Dumps to stdout instead of file.",
		},
	}
	c.Action = run.Combine

	return c
}
