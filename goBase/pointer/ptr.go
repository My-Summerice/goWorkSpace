package main

import "fmt"

func main(){
	var(
		cat int = 1
		str string = "zou"
	)
	fmt.Printf("%p %p", &cat, &str)

}
