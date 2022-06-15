package main

import (
	"fmt"
	"time"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	// 限定大于10
	Age			int			`form:"age" binding:"required,gt=10"`
	Name		string		`form:"name" binding:"required"`
	Birthday	time.Time	`form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func main() {
	r := gin.Default()
	r.GET("/man", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBind(&person); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintln(err))
			return
		}
		c.String(http.StatusOK, fmt.Sprintf("%#v", person))
	})
	r.Run(":8888")
}