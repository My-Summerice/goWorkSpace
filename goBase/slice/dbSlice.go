package main

import "fmt"

func main(){
    var a [][]int
    slice := [][]int{{10},{1,2}}
    a = slice
    a = append(a, slice...)
    fmt.Println(a,len(a[0]))
}
