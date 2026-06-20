package main

import (
	"HTTP-HAVEN/server"

	"net/http"
)

func main() {

	http.HandleFunc("/ping", server.Homeserver)

	println("server is now live")
	http.ListenAndServe(":8080", nil)
}
