package main

// repo: https://github.com/spf13/pflag

// 使用 pflag 可以无缝替代 flag

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"log"
	"os"
)

var (
	help    bool
	version bool
	cfgPath string
	cfgPort int
)

func usage() {
	fmt.Fprintf(os.Stderr, `a simple http server create by yangan

Options:
`)
	flag.PrintDefaults()
}

func init() {
	flag.BoolVar(&help, "help", false, "the usage of http server")
	flag.BoolVar(&version, "version", false, "the version of http server")
	flag.StringVar(&cfgPath, "cfgPath", "conf.yaml", "the config path of http server")
	flag.IntVar(&cfgPort, "cfgPort", 8001, "the port of http server")

	// 重写 usage
	flag.Usage = usage
}

func main() {
	// 解析传入的命令行参数
	flag.Parse()

	if help {
		flag.Usage()
		return
	}

	if version {
		_, err := fmt.Fprint(os.Stdout, "http server version: v1.0.0")
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	// 启动服务器
	fmt.Printf("server start, port is: %v", cfgPort)
}
