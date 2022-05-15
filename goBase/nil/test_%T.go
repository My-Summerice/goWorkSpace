package main

import "fmt"

func main(){
    type Human struct{
        Name string
    }
    var people = Human{Name: "zhangsan"}
    fmt.Printf(" %T\n %v\n %+v\n %#v\n", people, people, people, people)
}
