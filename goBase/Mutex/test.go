package main

import (
    "fmt"
    "sync"
)

func f1(ch1, ch2 chan int, wg *sync.WaitGroup, m *sync.Mutex) {
    for i := 1; i < 101; i += 3 {
        <- ch1
        m.Lock()
        fmt.Printf("f1 print: %d\n", i)
        ch2 <- i
        m.Unlock()
    }
    wg.Done()
}

func f2(ch3, ch2 chan int, wg *sync.WaitGroup, m *sync.Mutex) {
    for i := 2; i < 99; i += 3 {
        <- ch2
        m.Lock()
        fmt.Printf("f2 print: %d\n", i)
        ch3 <- i
        m.Unlock()
    }
    wg.Done()
}

func f3(ch1, ch3 chan int, wg *sync.WaitGroup, m *sync.Mutex) {
    for i := 3; i < 100; i += 3 {
        if i == 3 {
            ch1 <- i
        }
        <- ch3
        m.Lock()
        fmt.Printf("f3 print: %d\n", i)
        ch1 <- i
        m.Unlock()
    }
    wg.Done()
}
/*
func myPrint(ch1, ch2, ch3 chan int) {
    m.Lock()
    select {
    case a := <- ch1:
        fmt.Println("f1 print: ", a)
    }
    select {
    case a := <- ch2:
        fmt.Println("f2 print: ", a)
    }
    select {
    case a := <- ch3:
        fmt.Println("f1 print: ", a)
    }
    m.Unlock()
}
*/
func main() {
    var wg sync.WaitGroup
    var m sync.Mutex
    ch1 := make(chan int, 1)
    ch2 := make(chan int, 1)
    ch3 := make(chan int, 1)
    wg.Add(3)
    go f1(ch1, ch2, &wg, &m)
    go f2(ch3, ch2, &wg, &m)
    go f3(ch2, ch3, &wg, &m)
    wg.Wait()
//    go myPrint(ch1, ch2, ch3)
    /*
    for i := 1; i < 101; i++{
        fmt.Printf("f%d print: %d\n", i % 3, <- ch1)
    }
    
    for {
        select {
        case a := <- ch1:
            fmt.Println("f1 print: ", a)
        }
        select {
        case a := <- ch1:
            fmt.Println("f2 print: ", a)
        }
        select {
        case a := <- ch1:
            fmt.Println("f1 print: ", a)
        }
        if 0 == <- ch3 {
            break
        }
    }*/
}
