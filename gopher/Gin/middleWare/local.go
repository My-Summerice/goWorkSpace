package main

import (
	"fmt"
	"time"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定义中间件
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t1 := time.Now()
		fmt.Println("中间件开始执行了", t1)
		// 设置变量到Context的key中
		c.Set("request", "中间件")
		// c.Next() 之前的操作是在 Handler 执行之前就执行；
		// c.Next() 之后的操作是在 Handler 执行之后再执行；
		// 之前的操作一般用来做验证处理，访问是否允许之类的。
		// 之后的操作一般是用来做总结处理，比如格式化输出、响应结束时间，响应时长计算之类的。
		c.Next()
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t1)
		fmt.Println("time", t2)
	}
}

func main() {
	// 创建路由
	// 默认使用2个中间件Logger(),Recovery()
	r := gin.Default()
	// 局部中间件使用
	r.GET("/local", MiddleWare(), func(c *gin.Context) {
		// 取值
		req, _ := c.Get("request")
		fmt.Println("request:", req)
		// 页面接收
		c.JSON(http.StatusOK, gin.H{"request": req})
	})

	r.Run()
}