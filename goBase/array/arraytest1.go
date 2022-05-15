package main

import "fmt"

func main(){
	/*使用省略号定义
	q := [...]int{1,2,3}
	fmt.Printf("%T\n", q)
	fmt.Printf("%T\n", q[1])
	*/
	a := [...][3]int{{1,2,3},3:{3}}
	/*使用：初始化某个元素
	var array [4][2]int
	array = [4][2]int{0:{1},3:{}}
	*/
	fmt.Println(a)
	fmt.Println(a[:])
	fmt.Println(a[1:3])
	
}
