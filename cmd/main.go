package main

import (
	"net/http"

	"github.com/immanu10/splitgo/internal/handlers"
)


func main(){
	http.HandleFunc("/",handlers.RootHandler)	
	http.HandleFunc("/login",handlers.LoginHandler)	
	http.HandleFunc("/signup",handlers.SignUpHandler)	

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/",http.StripPrefix("/static/",fs))

	http.ListenAndServe(":8080", nil)
}