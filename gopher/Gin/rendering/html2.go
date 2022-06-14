package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	define定义模板名字LoadHTMLGlob渲染多个HTML模板


*/

func main() {
	r := gin.Default()
	// **代表所有目录,*代表所有文件
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/posts", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.html", gin.H{"title": "posts and posts"})
	})
	r.GET("/users", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.html", gin.H{"title": "users and users"})
	})
	r.GET("/nothing", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "nothing~~~~~"})
	})

	r.Run()
}