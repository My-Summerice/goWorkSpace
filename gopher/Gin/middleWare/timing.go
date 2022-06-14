package main

import (
	"fmt"
	_"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func myTime(c *gin.Context) {
	start := time.Now()
	c.Next()
	totalTime := time.Since(start)
	fmt.Println("总共用时：", totalTime)
}

func main() {
	// 创建路由
	// 默认使用两个中间件Logger(),Recovery()
	r := gin.Default()
	// 注册中间件
	r.Use(myTime)
	// 注册路由组
	testGroup := r.Group("/test")
	{
		//testGroup.GET("/test1", myTime, testTime1)
		testGroup.GET("/test1", testTime1)
		testGroup.GET("/test2", testTime2)
	}

	r.Run()
}

func testTime1(c *gin.Context) {
	time.Sleep(5 * time.Second)
}

func testTime2(c *gin.Context) {
	time.Sleep(3 * time.Second)
}