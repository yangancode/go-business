package main

// repo: https://github.com/jessevdk/go-flags

import (
	"fmt"
	flags "github.com/jessevdk/go-flags"
)

type Cmd struct {
	Help    bool     `short:"h" long:"help" description:"the usage of http server"`
	Version bool     `short:"v" long:"version" description:"the version of http server"`
	CfgPath string   `short:"c" long:"cfgPath" description:"the config path of http server"`
	CfgPort int      `short:"p" long:"cfgPort" description:"the port of http server"`
	Names   []string `short:"n" long:"names" description:"just for test"`
}

func main() {
	var cmd Cmd
	_, err := flags.Parse(&cmd)
	if err != nil {
		fmt.Println("Parse error:", err)
		return
	}

	fmt.Println("cmd: ", cmd.Help, cmd.Version, cmd.CfgPath, cmd.CfgPort, cmd.Names)
}
