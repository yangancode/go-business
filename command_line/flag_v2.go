package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// flag 定义变量的第二种方式

var (
	help    *bool
	version *bool
	cfgPath *string
	cfgPort *int
)

func usage() {
	fmt.Fprintf(os.Stderr, `a simple http server create by yangan

Options:
`)
	flag.PrintDefaults()
}

func init() {
	help = flag.Bool("help", false, "the usage of http server")
	version = flag.Bool("version", false, "the version of http server")
	cfgPath = flag.String("cfgPath", "conf.yaml", "the config path of http server")
	cfgPort = flag.Int("cfgPort", 8001, "the port of http server")

	// 重写 usage
	flag.Usage = usage
}

func main() {
	// 解析传入的命令行参数
	flag.Parse()

	if *help {
		flag.Usage()
		return
	}

	if *version {
		_, err := fmt.Fprint(os.Stdout, "http server version: v1.1.0")
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	// 启动服务器
	fmt.Printf("server start, port is: %v", *cfgPort)
}
