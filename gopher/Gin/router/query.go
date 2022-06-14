package main

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
)
/*
    1.URL的query参数可以通过DefaultQuery()或Query()方法获取
    2.DefaultQuery()若参数不存在，返回默认值，Query()若不存在，返回空串
    3.API?name=zs
	http://localhost:8080/user ?name=01	//报404
	http://localhost:8080/user? name=01	//仍显示程序中的默认值
	http://localhost:8080/user?name=01	//正确的修改参数形式
*/
func main() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		// 指定默认值
		// name := c.DefaultQuery("name", "summerice")
		name := c.Query("name")
		c.String(http.StatusOK, fmt.Sprintf("hello %s!", name))
	})

	r.Run()
}