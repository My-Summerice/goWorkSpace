package main

import (
	"fmt"
    "net/http"
    _"time"

    "github.com/gin-gonic/gin"
    _"github.com/gin-gonic/gin/binding"
    "github.com/go-playground/validator/v10"
)

/*
介绍

Validator 是基于 tag（标记）实现结构体和单个字段的值验证库，它包含以下功能：

    使用验证 tag（标记）或自定义验证器进行跨字段和跨结构体验证。
    关于 slice、数组和 map，允许验证多维字段的任何或所有级别。
    能够深入 map 键和值进行验证。
    通过在验证之前确定接口的基础类型来处理类型接口。
    处理自定义字段类型（如 sql 驱动程序 Valuer）。
    别名验证标记，它允许将多个验证映射到单个标记，以便更轻松地定义结构体上的验证。
    提取自定义的字段名称，例如，可以指定在验证时提取 JSON 名称，并在生成的 FieldError 中使用该名称。
    可自定义 i18n 错误消息。
    Web 框架 gin 的默认验证器。
*/

// 定义一个添加用户参数结构体
type AddUser struct {
    Name    string  `form:"name" validate:"required"`
    Email   string  `form:"email" validate:"required"`
    Age     uint8   `form:"age" validate:"required,gte=15,lte=30"`
}

func main() {
    r := gin.Default()
    r.GET("/addUser", func(c *gin.Context) {
        var addUser AddUser
        // 解析并绑定数据
        err := c.ShouldBind(&addUser)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
            return
        }
        /*  v     值的默认格式
            %+v   添加字段名(如结构体就先输出结构体的字段类型，再输出该字段的值)
            %#v　 相应值的Go语法表示   */
        fmt.Printf("addUser: %#v\n", addUser)
        // 使用Validator验证
        validate := validator.New()
        err = validate.Struct(addUser)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
            return
        }
        c.JSON(http.StatusOK, gin.H{"msg": "add success!"})
    })

    r.Run(":8888")
}

/*
当涉及到一些复杂的校验规则，这些已有的校验规则就不能满足我们的需求了。
例如现在有一个需求，存在db的用户信息中创建时间与更新时间都要大于某一时间，假设是从前端传来的（当然不可能，哈哈）。
现在我们来写一个简单示例，学习一下怎么对这个参数进行校验。
*/

/* 有bug
type Info struct {
    createTime  time.Time   `form:"createTime" binding:"required,timing" time_format:"2006-01-02"`
    updateTime  time.Time   `form:"updateTime" binding:"required,timing" time_format:"2006-01-02"`
}

// 自定义验证规则断言
func timing(fl validator.FieldLevel) bool {
    if date, ok := fl.Field().Interface().(time.Time); ok {
        today := time.Now()
        // 判断当前时间是否在指定时间之后
        if today.After(date) {
            return false
        }
    }
    return true
}

func main() {
    r := gin.Default()
    // 注册验证
    if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
        // 注册自定义标签
        err := v.RegisterValidation("timing", timing)
        if err == nil {
            fmt.Println("success")
        }
    }

    r.GET("/time", getTime)
    r.Run()
}

func getTime(c *gin.Context) {
    var now Info
    // 数据模型绑定查询字符串验证
    if err := c.ShouldBind(&now); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, "数据正确")
}
*/