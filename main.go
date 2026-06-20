package main

import (
	"HTTP-HAVEN/server"

	"net/http"
)

func main() {

	http.HandleFunc("/ping", server.Homeserver)
	http.HandleFunc("/hello", server.Hello)

	println("server is now live")
	http.ListenAndServe(":8080", nil)
}
