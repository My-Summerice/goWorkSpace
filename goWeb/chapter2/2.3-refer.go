package main

import (
	"net/http"
)

type Refer struct {
	handler http.Handler
	refer	string
}

func (this *Refer) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	if r.Referer() == this.refer {
		this.handler.ServeHTTP(w, r)
	}else {
		w.WriteHeader(403)
	}
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is myHandler"))
}

func myHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello summerice !"))
}

func main() {
	referer := &Refer{
		handler: http.HandlerFunc(myHandler),
		refer:	"www.shirdon.com",
	}
	http.HandleFunc("/hello", myHello)
	http.ListenAndServe(":8082", referer)
}
