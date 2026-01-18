package main

import "github.com/zhangga/go_hotreload"

func main() {
	if err := go_hotreload.StartWebUI(8080); err != nil {
		panic(err)
	}
}
