package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个gin模板引擎路由
	r := gin.Default()

	// 解析模板
    // Gin框架中使用LoadHTMLGlob()或者LoadHTMLFiles()方法进行HTML模板渲染
    // LoadHTMLGlob()方法是加载一组模板文件;
    // LoadHTMLFiles()方法是加载指定的一个模板文件；
    // 模板通常都放在templates文件夹下的
	r.LoadHTMLGlob("templates/index.html")
	r.GET("/index", func(c *gin.Context) {
		// HTML是*Context的一个方法,需要传入3个参数(code int, name string, obj interface{});无返回值
		// code是http响应的状态码，比如写200;name 是模板的名称(当模板没有通过define命名的时候，默认模板名字就是文件名),
		// 可以把H理解成一个json格式的
		// 返回的是渲染好的模板文件
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "我是测试"})
	})
	r.Run()
}