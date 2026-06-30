package main

import (
	"ascii-art-web/handlers"
	"log"
	"net/http"
)

func main() {
	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("./web"))))
	http.HandleFunc("/", handlers.HandlerSwitch)

	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
