package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/bilibili", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.bilibili.com")
	})
	r.GET("/baidu", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})

	r.Run()
}