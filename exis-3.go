package main

import (
	"fmt"
	"io"
	"net/http"
)

func CountHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/count" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method == http.MethodGet {
		fmt.Fprint(w, "Send a POST request with text to count words")
		return
	}

	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		length := len(body)

		fmt.Fprintf(w, "%d", length)
		return
	}

	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func main() {
	http.HandleFunc("/", CountHandler)

	fmt.Println("Server running on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
