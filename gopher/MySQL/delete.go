package main

import (
	"fmt"
)

func delete(id int) {
	str := "delete from person where user_id=?"

	de, err := db.Exec(str, id)
	if err != nil {
		fmt.Println("delete failed, ", err)
        return
	}

    row, err := de.RowsAffected()
    if err != nil {
        fmt.Println("rows failed, ", err)
        return
    }

	fmt.Println("delete succ: ", row)
}