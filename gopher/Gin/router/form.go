package main

import(
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	表单参数

    1.表单传输为post请求，http常见的传输格式为四种：
        application/json
        application/x-www-form-urlencoded
        application/xml
        multipart/form-data 
    2.表单参数可以通过PostForm()方法获取，该方法
	默认解析的是x-www-form-urlencoded或from-data格式的参数
*/

func main() {
	r := gin.Default()
	r.POST("/form", func(c *gin.Context) {
		types := c.DefaultPostForm("type", "post")
		username := c.PostForm("username")
		password := c.PostForm("userpassword")
		c.String(http.StatusOK, fmt.Sprintf("username:%s\npassword:%s\ntype:%s", username, password, types))
	})
	r.Run()
}