package main

import "fmt"

func main(){
	i := 3
	N := 6
	var a []int
	for i := 0; i < 20; i++ {
	    a = append(a, i)
	    fmt.Println(len(a),cap(a))
	    fmt.Printf("%p\n", a)
	}
	//输出切片
	fmt.Println(a)
	fmt.Println(len(a),cap(a))
	fmt.Printf("%p\n", a)
	//移动数据指针
	a = a[1:] // 删除开头1个元素
	fmt.Println(a)
	fmt.Println(len(a),cap(a))
	fmt.Printf("%p\n", a)
	a = a[:len(a)-1] // 删除尾部1个元素
	fmt.Println(a)
	fmt.Println(len(a),cap(a))
	fmt.Printf("%p\n", a)
	//append()
	a = append(a[:i], a[i+1:]...) // 删除下标为i的元素
	fmt.Println(a)
	fmt.Println(len(a),cap(a))
	fmt.Printf("%p\n", a)
    a = append(a[:i], a[i+N:]...) // 删除下标i到i+N-1的元素
	fmt.Println(a)
	fmt.Println(len(a),cap(a))
	fmt.Printf("%p\n", a)
	//copy()
    a = a[:i+copy(a[i:], a[i+1:])] // 删除下标为i的元素
	fmt.Println(a)
	fmt.Println(len(a),cap(a))
	fmt.Printf("%p\n", a)
	a = a[:i+copy(a[i:], a[i+N:])] // 删除下标为i到i+N-1的元素
	fmt.Println(a)
	fmt.Println(len(a),cap(a))
	fmt.Printf("%p\n", a)
	//测试append()和copy()是否会影响地址
	a = append(a[:0],a[1:]...)
	fmt.Println(a)
    fmt.Println(len(a),cap(a))
    fmt.Printf("%p\n", a)
	a = a[:copy(a[0:], a[1:])]
	fmt.Println(a)
    fmt.Println(len(a),cap(a))
	fmt.Printf("%p\n", a)
	//测试使用指针删除切片开头元素是否会导致重新分配内存
	a = a[1:] // 删除开头1个元素
    fmt.Println(a)
    fmt.Println(len(a),cap(a))
    fmt.Printf("%p\n", a)

}
