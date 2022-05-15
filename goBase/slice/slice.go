package main

import "fmt"

func main(){
	/*
	a := [...][3]int{{1,2,3},3:{3}}
	fmt.Println(a)
	fmt.Println(a[1:3])
	*/
	// 声明字符串切片
    	var strList []string
    	// 声明整型切片
    	var numList []int
    	// 声明一个空切片
    	var numListEmpty = []int{}
    	// 输出3个切片
    	fmt.Println(strList, numList, numListEmpty)
    	// 输出3个切片大小
    	fmt.Println(len(strList), len(numList), len(numListEmpty))
    	// 切片判定空的结果
    	fmt.Println(strList == nil)
    	fmt.Println(numList == nil)
    	fmt.Println(numListEmpty == nil)


}
