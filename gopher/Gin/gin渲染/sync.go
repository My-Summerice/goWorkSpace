package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

/*
    goroutine机制可以方便地实现异步处理
    另外，在启动新的goroutine时，不应该使用原始上下文，必须使用它的只读副本
*/

func main() {
	// 创建路由
	// 默认使用两个中间件Logger(),Recovery()
	r := gin.Default()
	// 1.异步async
	r.GET("/async", func(c *gin.Context) {
		// 拷贝一个只读副本用于开启协程
		copyContext := c.Copy()
		// 异步处理
		go func() {
			time.Sleep(3 * time.Second)
			log.Println("异步执行：" + copyContext.Request.URL.Path)
		}()
	})
	// 2.同步
	r.GET("/sync", func(c *gin.Context) {
		time.Sleep(3000)
		log.Println("同步执行：" + c.Request.URL.Path)
	})

	r.Run(":12345")
}