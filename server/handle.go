package server

import (
	
	"fmt"
	"net/http"
)

func Homeserver(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "pong")
}