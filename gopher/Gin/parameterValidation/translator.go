package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhs "github.com/go-playground/validator/v10/translations/zh"
)

var (
	validate = validator.New()			// 实例化验证器
	chinese	= zh.New()					// 获取中文翻译器
	uni		= ut.New(chinese, chinese)	// 设置成中文翻译器
	trans, _ = uni.GetTranslator("zh")	// 获取翻译字典
)

type User struct {
	Name	string	`form:"name" validate:"required,min=3,max=12"`
	Email	string 	`form:"email" validate:"required,email"`
	Age		uint8	`form:"age" validate:"required,gte=8,lte=30"`
}

func main() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		// 解析并绑定数据
		var user User
		err := c.ShouldBind(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err})
			return
		}
		// 注册翻译器
		_ = zhs.RegisterDefaultTranslations(validate, trans)
		// 使用验证器验证
		err = validate.Struct(user)
		if err != nil {
			if errors, ok := err.(validator.ValidationErrors); ok {
				// 翻译，并返回
				c.JSON(http.StatusInternalServerError, gin.H{
					"翻译前": errors.Error(),
					"翻译后": errors.Translate(trans),
				})
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{"msg": "success"})
	})
	r.Run()
}