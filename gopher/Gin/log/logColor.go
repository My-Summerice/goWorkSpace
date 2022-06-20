package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// 禁止日志颜色
	gin.DisableConsoleColor()
	// 强制日志着色
	//gin.ForceConsoleColor()

	// logging to a file.
	f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	// 如果需要同时将日至写入文件和控制台，使用以下代码
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()
	r.GET("/logging", func(c *gin.Context) {
		c.JSON(200, "success")
	})

	r.Run()
}