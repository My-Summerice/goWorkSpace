package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/user", printMessage)
	r.NoRoute(printNotFound) 
	r.Run()
}

func printMessage(c *gin.Context) {
	message := c.DefaultQuery("message", "是的，你成功了，这里没有404！")
	c.String(http.StatusOK, fmt.Sprintf("运行结果怎么样呢？\n%s", message))
}

func printNotFound(c *gin.Context) {
	c.String(http.StatusNotFound, "404 NOT FOUND by summerice")
}