package main

import (
	"fmt"
	"strings"
)

func main() {
	// no sep in s
	s := strings.SplitN("a,b,c,d", "!", 4)
	fmt.Println(s)

	// sep == nil
	s = strings.SplitN("a,b,c,d", "", 4)
	fmt.Println(s)

	// sep == nil, s == nil
	s = strings.SplitN("", "", 4)
	fmt.Println(s)

	// n == 0
	s = strings.SplitN("a,b,c,d", ",", 0)
	fmt.Println(s)

	// n > 0 
	s = strings.SplitN("a,b,c,d", ",", 3)
	fmt.Println(s)

	// n < 0 
	s = strings.SplitN("a,b,c,d", ",", -1)
	fmt.Println(s)
}