package main

import (
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Get" {
		http.Error(w, "Page Not Found!", http.StatusNotFound)
		return
	}

	json := ` {"message": "User Created"}`

	w.Header().Set("content-type", "application/json")

	w.Write([]byte(json))

}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/AboutPage" {
		http.Error(w, "AboutPage Not found", http.StatusNotFound)
		return
	}

	w.Write([]byte("AboutPage"))
}

func main() {

	http.HandleFunc("/Get", HomeHandler)
	http.HandleFunc("/AboutPage", AboutPage)

	println("server is now live on port http://:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
