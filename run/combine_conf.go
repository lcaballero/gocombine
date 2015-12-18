package run

import (
	cmd "github.com/codegangsta/cli"
)

type CombineConf struct {
	CppCode string
	CppFile string
	GoCode  string
	GoFile  string
	Out     string
	Stdout  bool
}

func NewCombineConf(cmd *cmd.Context) *CombineConf {
	return &CombineConf{
		CppFile: cmd.String("cc"),
		GoFile:  cmd.String("go"),
		Out:     cmd.String("out"),
		Stdout:  cmd.Bool("stdout"),
	}
}

func (c *CombineConf) IsValid() bool {
	hasCpp := c.CppFile != ""
	hasGo := c.GoFile != ""
	hasOut := c.Out != ""

	return hasCpp && hasGo && (hasOut || c.Stdout)
}