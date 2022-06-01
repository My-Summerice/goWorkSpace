package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建路由，默认使用2个中间件Logger(),Recovery()
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get err %s"), err.Error())
			return
		}
		// 获取所有文件
		files := form.File["files"]
		// 遍历所有文件
		for _, file := range files {
			// 逐个存
			if err := c.SaveUploadedFile(file, "./upload/" + file.Filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s"), err.Error())
				return
			}
		}
		c.String(http.StatusOK, fmt.Sprintf("uploaded %d files over", len(files)))
	})
	r.Run()
}