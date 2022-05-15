package main

import "fmt"


func main(){
	var a [3]int
	fmt.Println(a[0])
	fmt.Printf("%d\n\n", a[len(a)-1])
	//打印索引和元素
	for i, v := range a{
		fmt.Printf("%d %d\n", i, v)
		if i == len(a)-1 {
			fmt.Printf("\n")
		}
	}
	//只打印元素
	for _, v := range a{
		fmt.Printf("%d\n", v)
	}
}

