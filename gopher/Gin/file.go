package main

import (
	"net/http"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"	
)

/*
	上传单个文件

    1.multipart/form-data格式用于文件上传

    2.gin文件上传与原生的net/http方法类似，不同在于gin把原生的request封装到c.Request中

*/

/*
func main() {
	r := gin.Default()
	// 限制上传最大尺寸（有大佬验证说是为了限制处理该文件所占用的最大内存https://www.jianshu.com/p/9ec31f232f5f）
	r.MaxMultipartMemory = 8 << 20	// 8bit左移20位相当于8bit*2^10=8byte, 8byte*2^10=8M
	r.POST("/upload", func(c *gin.Context) {
		// 从请求中拿到文件流
		file, err := c.FormFile("file")
		if err != nil {
			c.String(500, "上传图片出错")
		}

		c.SaveUploadedFile(file, "summerice-" + file.Filename)
		c.JSON(http.StatusOK, gin.H{"文件保存成功": file.Filename})
		//c.String(http.StatusOK, fmt.Sprintf("%v ", r))
	})
	
	r.Run()
}
*/


func main() {
	r := gin.Default()
	r.POST("/upload", func(c *gin.Context) {
		_, headers, err := c.Request.FormFile("file")
		if err != nil {
			log.Printf("Error when try to get file: %v", err)
		}
		// headers.Size 获取文件大小
		if headers.Size > 1024*1024*2 {
			fmt.Println("文件太大了！")
			return
		}
		//headers.Header.Get("Content-Type")获取上传文件的类型
		if headers.Header.Get("Content-Type") != "image/png" {
			fmt.Println("只允许上传png图片")
			return
		}
		c.SaveUploadedFile(headers, "./upload/" + headers.Filename)
		c.String(http.StatusOK, headers.Filename)
	})
	r.Run()
}
