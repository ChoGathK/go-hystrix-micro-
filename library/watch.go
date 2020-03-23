package library

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Watch - 异常捕捉 (中间件)
func Watch(ctx *gin.Context) {

	defer func() {
		if r := recover(); r != nil {
			// 异常日志
			Panic(fmt.Sprintf("%s", r))
			// 返回异常提示
			ctx.JSON(http.StatusBadGateway, "BadGateway")
		}
	}()

	// 先执行
	ctx.Next()

}
