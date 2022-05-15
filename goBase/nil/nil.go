package main

import "fmt"

func main(){
    var m map[string]int
    var p *int
    fmt.Printf("%T\n", nil) 
    fmt.Printf("m.pointer: %p\np.pointer: %p\n", m, p)
    //fmt.Println(nil == nil)
    //fmt.Println(m == p)
}
