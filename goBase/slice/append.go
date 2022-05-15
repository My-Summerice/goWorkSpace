package main

import "fmt"

func main(){
	/*
	var a []int
	a = append(a, 1)
	fmt.Println(a)
	a = append(a, 2, 3)
	fmt.Println(a)
	a = append(a,/* []int{4,5,6},/ make([]int, 2)...)
	fmt.Println(a)
	a = append(a, a[0:3]...)
	*/
	var a []int
	for i := 0; i < 10; i++ {
		a = append(a, i)
		fmt.Printf("len:%d cap:%d pointer:%p\n", len(a), cap(a), a)
	}
	a = append(a,append(a,1)...)
	fmt.Printf("%d\nlen:%d cap:%d pointer:%p\n", a, len(a), cap(a), a)
}

