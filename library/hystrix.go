package library

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
)

// InitHystrix 根据路由路径，注册熔断器
func InitHystrix(path string) {

	hystrix.ConfigureCommand(path, hystrix.CommandConfig{
		Timeout:                int(time.Second * 10),  // 超时时间
		MaxConcurrentRequests:  100,                    // 最大并发数，超过并发返回错误
		SleepWindow:            int(time.Second * 100), // 熔断尝试恢复时间，单位毫秒
		RequestVolumeThreshold: 2,                      // 熔断探测前的调用次数
		ErrorPercentThreshold:  50,                     // 错误率阀值，百分比。达到阀值，启动熔断
	})

}

// HystrixManager 熔断处理器 (中间件)
func HystrixManager(ctx *gin.Context) {

	defer func() {

		err := hystrix.Do(ctx.Request.RequestURI, func() error {
			if ctx.Writer.Status() >= http.StatusBadRequest {
				return errors.New("TEST--")
			}
			return nil
		}, nil)

		// 熔断器处理异常后抛出，通知 watch 捕捉
		if err != nil {
			// circuit open 开启断路器模式时会报错
			log.Println(err)
		}

	}()

	// 先执行
	ctx.Next()

}
