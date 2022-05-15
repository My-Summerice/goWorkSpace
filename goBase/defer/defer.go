package main

import (
    "fmt"
)

type person struct {
    name string
    age int
}

func (p person) man() {
    defer hello()
    fmt.Println(p.name, "\n", p.age)
}

func hello() {
    fmt.Println("hello")
}

func main() {
    p := person {
        name: "ice",
        age: 20,
    }
    defer p.man()
    p.age = 21
    p.man()
}
