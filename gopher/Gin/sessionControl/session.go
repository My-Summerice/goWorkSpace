package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

/*
Sessions

gorilla/sessions为自定义session后端提供cookie和文件系统session以及基础结构。

主要功能是：

    简单的API：将其用作设置签名（以及可选的加密）cookie的简便方法。
    内置的后端可将session存储在cookie或文件系统中。
    Flash消息：一直持续读取的session值。
    切换session持久性（又称“记住我”）和设置其他属性的便捷方法。
    旋转身份验证和加密密钥的机制。
    每个请求有多个session，即使使用不同的后端也是如此。
    自定义session后端的接口和基础结构：可以使用通用API检索并批量保存来自不同商店的session。
*/

// 初始化一个cookie存储对象
// something-very-secret应该是一个你自己的密钥，只要不被别人知道就行
var store = sessions.NewCookieStore([]byte("something-very-secret"))

func main() {
	http.HandleFunc("/save", SaveSession)
	http.HandleFunc("/get", GetSession)
	http.HandleFunc("/delete", DelSession)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("HTTP server failed, err:", err)
		return
	}
}

func SaveSession(w http.ResponseWriter, r *http.Request) {
	// 获取一个session对象，session-name是session的名字
	session, err := store.Get(r, "session-name")
	if err != nil {
		// StatusInternalServerError状态内部服务器错误
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 在session中存储值
	session.Values["name"] = "summerice"
	session.Values[10] = 10
	// 保存更改
	session.Save(r, w)
}

func GetSession(w http.ResponseWriter, r *http.Request) {
	// 同上
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// 获取session中存储的值
	name := session.Values["name"]
	number := session.Values[10]
	fmt.Printf("name is %s, number is %d\n", name, number)
}

func DelSession(w http.ResponseWriter, r *http.Request) {
	// 同上
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// 删除
	// 将session的最大存储时间设置为小于0即为删除
	session.Options.MaxAge = -1
	session.Save(r, w)
}