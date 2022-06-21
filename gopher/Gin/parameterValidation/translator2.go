package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

// 定义一个全局翻译器
var trans ut.Translator

// 初始化翻译器
func initTrans(locale string) (err error) {
	//  修改gin框架中的Validator引擎属性，实现自定义验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		
		/*
		// 注册一个获取 指定tag 的自定义方法
		v.RegisterTagNameFunc(func(stru interface{}) string {	//该方法传入的是结构体的指针
			t := reflect.TypeOf(stru).Elem()	// 使用反射包的方法获取指针指向的值对应的结构体内容
			for i = 0; i < t.NumField(); i++ {	// t.NumField()获取结构体的字段数量

			}
		})
		*/

		zhT := zh.New()	// 中文翻译器
		enT := en.New()	// 英文翻译器
		
		// 第一个参数是备用（fallback）的语言环境
		// 后面的参数是应该支持的语言环境（支持多个）
		// uni := ut.New(zhT, zhT) 也是可以的
		uni := ut.New(enT, zhT, enT)

		// locale 通常取决于 http 请求头的 “Accept-Language”
		// 通过断言获取需要的语种翻译器类型//个人理解
		// 也可使用 uni.FindTranslator(...) 传入多个locale进行查找
		var ok bool
		if trans, ok = uni.GetTranslator(locale); !ok {	//这一步的类型断言格式没看懂
			return fmt.Errorf("uni.GetTranslator(%s) faild", locale)
		}

		// 注册对应语种的翻译器
		switch locale {
		//case "en":
		//	err = enTranslations.RegisterDefaultTranslations(v, trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		}
		//return
	}
	return
}

type userInfo struct {
	Name		string	`form:"name" binding:"min=3,max=12"`// 貌似对值有限制就不用required
	Age			uint8	`form:"age" binding:"gte=8,lte=50"`
	Email 		string 	`form:"email" binding:"email,required"`
	Password	string	`form:"password" binding:"min=6,max=20,required"`
	RePassword	string	`form:"repassword" binding:"required,eqfield=Password"`	
}

func main() {
	if err := initTrans("zh"); err != nil {
		fmt.Printf("initTrans failed, err:%#v\n", err)
		return
	}

	r := gin.Default()

	r.GET("/userInfo", func(c *gin.Context) {
		var u userInfo
		// 解析前端数据并绑定到结构体u
		if err := c.ShouldBind(&u); err != nil {
			// 使用类型断言获取validator.ValidationErrors类型的errors
			errs, ok := err.(validator.ValidationErrors)
			if !ok {
				// 非此类型的错误直接返回err
				c.JSON(http.StatusOK, gin.H{
					"msg": err,
				})
				return
			}
			// 翻译错误
			c.JSON(http.StatusOK, gin.H{
				"msg": errs.Translate(trans),
			})
			return
		}
		// 写入数据库等具体业务逻辑代码...

		c.JSON(http.StatusOK, "success")
	})

	r.Run(":8989")
}