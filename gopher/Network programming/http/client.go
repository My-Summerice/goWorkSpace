package main

import (
    "fmt"
    "io"
    "net/http"
)

func main() {
     //resp, _ := http.Get("http://www.baidu.com")
     //fmt.Println(resp)
    resp, _ := http.Get("http://127.0.0.1:8000/go")
    fmt.Println(resp)
    defer resp.Body.Close()
    // 200 ok
    fmt.Println(resp.Status)
    fmt.Println(resp.Header)

    buf := make([]byte, 1024)
    for {
        // 接收服务端信息
        n, err := resp.Body.Read(buf)
    //    fmt.Println(n)
        if err != nil && err != io.EOF {
            fmt.Println(err)
            return
        } else {
            fmt.Println("读取完毕")
            res := string(buf[:n])
            fmt.Println(res)
            break
        }
    }
}
