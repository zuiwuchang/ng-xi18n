package main

import (
	"github.com/zuiwuchang/ng-xi18n/cmd"
	"log"
)

func main() {
	// 設置 版本 信息
	cmd.Version = Version
	// 執行命令
	if e := cmd.Execute(); e != nil {
		log.Fatalln(e)
	}
}
