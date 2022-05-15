package main
import (
        "container/list"
        "fmt"
)
func main(){
    l := list.New()
    // 尾部添加
    l.PushBack("zou")
    // 头部添加
    l.PushFront(325)
    // 尾部添加后保存元素句柄
    element := l.PushBack("ice")
    // 在"ice"后添加666
    l.InsertAfter(666, element)
    // 在"ice"前添加”summer"
    l.InsertBefore("summer", element)
    for i := l.Front(); i != nil; i = i.Next() {
        fmt.Println(i.Value)
    }
    // 删除
    l.Remove(element)
    for i := l.Front(); i != nil; i = i.Next() {
        fmt.Println(i.Value)
    }

}
