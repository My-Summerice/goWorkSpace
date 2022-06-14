package main

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
)

/*
	模拟实现权限验证中间件

		有2个路由，login和home
		login用于设置cookie
		home是访问查看信息的请求
		在请求home之前，先跑中间件代码，检验是否存在cookie

*/

/*	这样写中间件只能使用r.Use()全局调用，必须按照局部中间件格式书写
func checkMiddleWare(c *gin.Context) {
	if cookie, err := c.Cookie("cookie"); err == nil {
		if cookie == "myCookie" {
			// 在 Handler 执行之前进行验证
			c.Next()
			return 
		}
	}
	// 没有对应cookie时返回错误
	c.JSON(http.StatusUnauthorized, gin.H{"error": "err"})
	// 若验证不通过，不再调用后续函数处理
	c.Abort()
	return
}
*/

func checkMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Cookie("cookie"); err == nil {
			if cookie == "myCookie" {
				// 使 Handler 在Next()执行后才执行
				c.Next()
				return
			}
		}
		// 没有对应cookie时返回错误:状态未授权
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未持有正确cookie"})
		// 若验证不通过，不再调用后续函数处理，不加这个会导致JSON数据冲突报错
		c.Abort()
	}
}

func main() {
	r := gin.Default()

	// 用于设置cookie
	r.GET("/login", func(c *gin.Context) {
		// 设置cookie
		c.SetCookie("cookie", "myCookie", 60, "/", "localhost", false, true)
		// 返回信息
		c.String(http.StatusOK, fmt.Sprintln("login success!"))
	})

	//r.Use(checkMiddleWare)
	// 访问查看信息的请求
	r.GET("/home", checkMiddleWare(), func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "home"})
		// 不可以连续返回两次JSON数据，跟上面一样，会导致数据冲突报错
		//c.JSON(200, gin.H{"data": "home"})
	})

	r.Run(":8888")
}