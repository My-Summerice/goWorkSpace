package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic-gin"
)

func main() {
	r := gin.Default()

	r.POST("/loginFORM", func(c *gin.Context) {
		var form Login
		// Bind()默认解析并绑定form格式
		// 根据请求头中的content-type自动推断
		if err := c.Bind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error", err.Error()})
			return
		}
		// 判断用户名、密码是否正确
		if form.User != "root" || form.password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})
	r.Run(":8000")
}