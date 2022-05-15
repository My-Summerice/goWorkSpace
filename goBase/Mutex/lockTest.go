package main

import (
    "fmt"
    "sync"
)

/*
    开启3个协程交替打印1～100，目前程序仍有bug（deadlock）
*/

var ch1 = make(chan int)
var ch2 = make(chan int)
var ch3 = make(chan int)

func f1(wg *sync.WaitGroup) {
    for i := 1; i < 98; i += 3 {
        select {
        case <- ch1 :
                fmt.Println("f1 print: ", i)
                ch2 <- i
        }
    }
    wg.Done()
}

func f2(wg *sync.WaitGroup) {
    for i := 2; i < 99; i += 3 {
        select {
        case <- ch2 :
                fmt.Println("f2 print: ", i)
                ch3 <- i
        }
    }
    wg.Done()
}

func f3(wg *sync.WaitGroup) {
    for i := 3; i < 100; i += 3 {
        if i == 3 {
            ch1 <- i
        }
        select {
        case <- ch3 :
                fmt.Println("f3 print: ", i)
                ch1 <- i
        }
    }
    wg.Done()
}


func main() {
    var wg sync.WaitGroup
    wg.Add(3)
    go f1(&wg)
    go f2(&wg)
    go f3(&wg)
    wg.Wait()
}
