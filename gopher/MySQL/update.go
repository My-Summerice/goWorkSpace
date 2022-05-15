package main

import (
	"fmt"
)

func update(name string, id int) {
	str := "update person set user_name=? where user_id=?"

	up, err := db.Exec(str, name, id)
	if err != nil {
		fmt.Println("update failed, ", err)
		return
	}

	row, err := up.RowsAffected()	//被操作的行数
	if err != nil {
		fmt.Println("rows failed, ", err)
		return
	}

	fmt.Println("update succ: ", row)
}