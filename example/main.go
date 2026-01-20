package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhangga/go_hotreload"
	"github.com/zhangga/go_hotreload/example/internal"
)

func main() {
	// 启动热更新WebUI，监听8080端口
	if err := go_hotreload.StartWebUI(8080, pingHandle); err != nil {
		panic(err)
	}
}

// pingHandle 测试更新后是否生效
func pingHandle(r *gin.Engine) {
	r.GET("/example/ping", func(c *gin.Context) {
		err := internal.DoSomething()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": -1, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "pong"})
	})
}
