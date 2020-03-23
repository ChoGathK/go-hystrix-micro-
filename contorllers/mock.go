package contorllers

import (
	"github.com/gin-gonic/gin"
)

// Hello -
func Hello(c *gin.Context) {
	c.String(200, "hello")
}

// Error -
func Error(c *gin.Context) {
	// 获取 Get 参数
	panic("i am panic")
}
