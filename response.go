package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Request status not allowed Request", http.StatusMethodNotAllowed)
		return
	}

	home, err := template.ParseFiles("Homepage/homepage.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	home.Execute(w, nil)

}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/AboutPage" {
		http.Error(w, "404 Page Not Found!", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not allowed", http.StatusMethodNotAllowed)
		return
	}

	code, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	code.Execute(w, nil)

}

func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/AboutPage", AboutPage)

	println("server is starting now !!!")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
