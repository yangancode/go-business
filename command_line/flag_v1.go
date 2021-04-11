package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// 使用 flag 定制一个 HTTP 服务器

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
