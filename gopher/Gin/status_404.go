package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/user", notFound)
	r.NoRoute(printMessage) 
	r.Run()
}

func notFound(c *gin.Context) {
	message := c.DefaultQuery("message", "是的，你成功了，这里没有404！")
	c.String(http.StatusOK, fmt.Sprintf("运行结果怎么样呢？\n%s", message))
}

func printMessage(c *gin.Context) {
	c.String(http.StatusNotFound, "404 NOT FOUND by summerice")
}