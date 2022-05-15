package main

import (
	"log"
	"net/http"
)

func main() {
	//start Server
	srv := &http.Server{Addr: ":8082", Handler: http.HandlerFunc(handle)}
	//use TLS start server
	log.Printf("Serving on hattps://0.0.0.0:8082")
	log.Fatal(srv.ListenAndServeTLS("server.crt", "server.key"))
}

//handler
func handle(w http.ResponseWriter, r *http.Request) {
	//write request security
	log.Printf("Got connection: %s", r.Proto)
	//send a message to Client
	w.Write([]byte("Hello this is a http2 message!"))
}