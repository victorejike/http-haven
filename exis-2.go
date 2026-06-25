package main

import (
	"fmt"
	"net/http"
)

func HelloHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	name := r.URL.Query().Get("name")

	if name == "" {
		name = "Gust"
	}

	fmt.Fprintf(w, "Hello, %s!\n", name)

}

func main() {
	http.HandleFunc("/hello", HelloHandle)

	println("server is starting up now ")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

//
