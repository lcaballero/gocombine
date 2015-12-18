package run

import (
	"bytes"
	"fmt"
	"text/template"
	"io/ioutil"

	cmd "github.com/codegangsta/cli"
)

func Combine(cmd *cmd.Context) {
	conf := NewCombineConf(cmd)

	if conf.IsValid() {
		run(conf)
	} else {
		panic(fmt.Errorf("Missing required flags."))
	}
}

func read(name string) string {
	bb, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return string(bb)
}

func run(conf *CombineConf) {
	conf.GoCode = read(conf.GoFile)
	conf.CppCode = read(conf.CppFile)

	name := fmt.Sprintf("Combine: %s", conf.CppFile)

	t, err := template.New(name).Parse(conf.GoCode)
	t.Option()
	if err != nil {
		panic(err)
	}

	buf := bytes.NewBuffer([]byte{})

	err = t.Execute(buf, conf)
	if err != nil {
		panic(err)
	}

	if conf.Stdout {
		fmt.Println(buf.String())
		return
	}

	err = ioutil.WriteFile(conf.Out, buf.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}
