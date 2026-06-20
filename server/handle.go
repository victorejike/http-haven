package server

import (
	"fmt"
	"net/http"
)

func Homeserver(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

func Hello(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	name := r.URL.Query().Get("name")
	// Default to "Guest" if name parameter is missing
	if name == "" {
		name = "Guest"
	}

	// Respond with greeting
	fmt.Fprintf(w, "Hello, %s!", name)
}
