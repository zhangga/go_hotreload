package go_hotreload

import (
	"embed"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed webui/*
var uiStatic embed.FS

func StartWebUI(port int) error {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.StaticFS("/static", http.FS(uiStatic))

	r.GET("/", index)
	r.GET("/api/handle-path", handlePath)

	log.Printf("WebUI listening: http://127.0.0.1:%d", port)
	// 启动服务，监听端口
	err := r.Run(fmt.Sprintf(":%d", port))
	return err
}

func index(c *gin.Context) {
	file, err := uiStatic.ReadFile("webui/index.html")
	if err != nil {
		c.String(http.StatusInternalServerError, "加载首页失败：%v", err)
		return
	}
	// 设置响应头为html格式，返回首页内容
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(http.StatusOK, string(file))
}

func handlePath(c *gin.Context) {
	// 从前端获取传递的文件路径/文件名参数
	filePath := c.Query("filePath")
	if filePath == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, "msg": "文件路径参数不能为空"})
		return
	}

	fmt.Printf("✅ 触发功能：接收到文件路径参数 -> %s\n", filePath)

	// ===================== 【重点】在这里写你的自定义功能 =====================
	// 示例：根据文件名/路径执行你的业务逻辑
	// 比如：读取指定路径的文件、解析文件内容、调用其他函数等
	// 注意：如果是本地文件路径，需要保证后端程序有对应文件的读取权限

	// 返回处理结果给前端
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  fmt.Sprintf("已接收路径参数：%s，自定义功能执行完成！", filePath),
	})
}
