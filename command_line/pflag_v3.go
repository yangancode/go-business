package main

import (
	flag "github.com/spf13/pflag"
	"os"
)
import (
	"fmt"
)

var (
	help       bool
	version    bool
	cfgPath    string
	cfgPort    int
	serverPort int // 使用 cfgPort 代替 serverPort
)

func usage() {
	fmt.Fprintf(os.Stderr, `a simple http server create by yangan

Options:
`)
	flag.PrintDefaults()
}

func init() {
	flag.BoolVarP(&help, "help", "h", false, "the usage of http server")
	flag.BoolVarP(&version, "version", "v", false, "the version of http server")
	flag.StringVarP(&cfgPath, "cfgPath", "c", "conf.yaml", "the config path of http server")
	flag.IntVarP(&cfgPort, "cfgPort", "p", 8001, "the port of http server")
	flag.IntVarP(&serverPort, "serverPort", "s", 8001, "the port of http server")

	// 重写 usage
	flag.Usage = usage
}

func main() {
	// 把 serverPort 参数标记为即将废弃的，请用户使用 cfgPort 参数
	flag.CommandLine.MarkDeprecated("serverPort", "please use cfgPort instead")

	// 把 serverPort 参数的 shorthand 标记为即将废弃的，请用户使用 cfgPort 的 shorthand 参数
	flag.CommandLine.MarkShorthandDeprecated("serverPort", "please use -p instead")

	// 禁止排序
	flag.CommandLine.SortFlags = false

	// 解析参数
	flag.Parse()
	if help {
		flag.Usage()
		return
	}

	fmt.Printf("server start, port is: %v", cfgPort)
}