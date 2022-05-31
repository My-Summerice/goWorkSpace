package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行处理器函数
	// gin.Context，封装了request和response
	// 当客户端以GET方法请求'/'路径时，会自动执行后面的匿名函数
	r.GET("/",func(c *gin.Context) {
		// 返回字符串格式的数据
		c.String(http.StatusOK, "hello world!")
		/* 返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
		*/
	})

	r.GET("/user/:name/*action", func(c *gin.Context) {
		// Param函数获取URL的路径参数
		name := c.Param("name")
		action := c.Param("action")
		// 截取action中首个‘/’
		action = strings.Trim(action, "/")
		
		c.String(http.StatusOK, " username=%s\n action=%s", name, action)
	})

	
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	// r.Run()
	r.Run(":8000")
}