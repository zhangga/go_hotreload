package go_hotreload

import (
	"embed"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

//go:embed webui/*
var uiStatic embed.FS

func StartWebUI(port int, extraHandlers ...func(r *gin.Engine)) error {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.StaticFS("/static", http.FS(uiStatic))

	// 内置路由
	r.GET("/", index)
	r.GET("/api/handle-path", handlePath)
	// 额外路由
	for _, handler := range extraHandlers {
		handler(r)
	}

	log.Printf("WebUI listening: http://%s:%d", localIP(), port)
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

	// 调用加载补丁功能
	err := LoadPatch(c, filePath)

	// ===================== 【重点】在这里写你的自定义功能 =====================
	// 示例：根据文件名/路径执行你的业务逻辑
	// 比如：读取指定路径的文件、解析文件内容、调用其他函数等
	// 注意：如果是本地文件路径，需要保证后端程序有对应文件的读取权限

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  err.Error()})
		return
	}
	// 返回处理结果给前端
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  fmt.Sprintf("已接收路径参数：%s，自定义功能执行完成！", filePath),
	})
}

func localIP() string {
	// 优先获取运维配置机器的环境变量
	hostIP := os.Getenv("HOST_IP")
	if len(hostIP) > 0 {
		return hostIP
	}

	// 优先获取外网ip
	if conn, err := net.Dial("udp", "8.8.8.8:53"); err == nil {
		laddr := conn.LocalAddr().(*net.UDPAddr)
		ip := strings.Split(laddr.String(), ":")[0]
		return ip
	}

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "127.0.0.1"
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	return "127.0.0.1"
}
