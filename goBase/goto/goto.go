package main
import "fmt"

func main(){
    a := 1
    if a != 0 {
        goto test1
    }

    a = 3
    fmt.Println(a)
    return

    test1:
        fmt.Println(a)
}
