package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	__l "main.go/library"
)

func main() {
	// 创建服务实例
	app := gin.New()

	// 打印请求日志
	app.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
		)
	}))

	// 使用熔断器
	app.Use(__l.HystrixManager)

	// 使用异常捕捉
	app.Use(__l.Watch)

	// 加载路由
	__l.LoadRouter(app)

	// 启动
	app.Run(":4396")
}
