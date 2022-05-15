package main

import (
	"fmt"
)

func insert(name, sex, email string) {
    str := "insert into person(user_name, sex, email) values(?, ?, ?)"

    in, err := db.Exec(str, name, sex, email)
    if err != nil {
        fmt.Println("exec failed, ", err)
        return
    }

    id, err := in.LastInsertId()     //获得刚刚添加数据的自增ID
    if err != nil {
        fmt.Println("exec failed, ", err)
        return
    }

    fmt.Println("insert succ: ", id)
}