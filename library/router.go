package library

import (
	"github.com/gin-gonic/gin"
	__c "main.go/contorllers"
)

// LoadRouter - 加载路由
func LoadRouter(app *gin.Engine) {

	router := app.Group("/micro")

	// 初始化熔断器
	InitHystrix("/micro/test/get/error")
	InitHystrix("/micro/test/get/hello")

	// 放置 contorllers
	{
		router.GET("/test/get/error", __c.Error)
		router.GET("/test/get/hello", __c.Hello)
	}

}
