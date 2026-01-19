package main

import "github.com/zhangga/go_hotreload"

func main() {
	// 启动热更新WebUI，监听8080端口
	if err := go_hotreload.StartWebUI(8080); err != nil {
		panic(err)
	}
}
