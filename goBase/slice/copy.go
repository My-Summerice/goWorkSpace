package main

import "fmt"

func main(){
	slice1 := []int{1,2,3,4,5,6}
	slice2 := []int{7,8,9}
	var slice3 []int
	a := copy(slice1, slice2)
	b := copy(slice2, slice1[3:])
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(copy(slice3, slice1))
	fmt.Println(slice3)
}
