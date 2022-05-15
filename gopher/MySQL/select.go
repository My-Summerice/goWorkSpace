package main

import (
	"fmt"
)

func slt(name string) {
	var person []Person
	str := "select * from person where user_name = ?"

	err := db.Select(&person, str, name)
	if err != nil {
		fmt.Println("select failed, ", err)
		return
	}

	fmt.Println("select succ: ", person)
}