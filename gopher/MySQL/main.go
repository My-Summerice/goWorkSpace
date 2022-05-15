package main

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"  //第三方开源的mysql驱动
    "github.com/jmoiron/sqlx"   //基于mysql驱动的封装
)

//定义变量用于接收数据库中对应变量的值
type Person struct {
    UserID      int     `db:"user_id"`
    UserNAME    string  `db:"user_name"`
    Sex         string  `db:"sex"`
    Email       string  `db:"email"`
}

type Place struct {
    Country     string  `db:"country"`
    City        string  `db:"city"`
    TelCode     int     `db:"telcode"`
}

//定义一个数据库类型的变量
var db *sqlx.DB

//初始化数据库
func initDB() {
    database, err := sqlx.Open("mysql", "root:summerice@tcp(127.0.0.1:3306)/test?charset=utf8")
    if err != nil {
        fmt.Println("open mysql failed,", err)
        return
    }
    db = database
    fmt.Println("mysql连接成功")
    //defer语句要写在可能出现的return的后面，否则在执行return后会报错：sql: database is closed
    //defer db.Close()  
}

func main() {
    initDB()
    defer db.Close()

    //insert("帝释天", "男", "654321")
    //delete(8)
    //update("张三",4)  
    slt("summerice")
}


