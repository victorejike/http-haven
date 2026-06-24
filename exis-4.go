package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/calculate" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	op := r.URL.Query().Get("op")
	aStr := r.URL.Query().Get("a")
	bStr := r.URL.Query().Get("b")

	a, err := strconv.Atoi(aStr)
	if err != nil {
		http.Error(w, "Invalid value for a", http.StatusBadRequest)
		return
	}

	b, err := strconv.Atoi(bStr)
	if err != nil {
		http.Error(w, "Invalid value for b", http.StatusBadRequest)
		return
	}

	var result int

	switch op {
	case "add":
		result = a + b
	case "subtract":
		result = a - b
	case "multiply":
		result = a * b
	default:
		http.Error(w, "Unknown operation", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Result: %d", result)
}

func main() {
	http.HandleFunc("/", CalculateHandler)

	fmt.Println("Server running on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
