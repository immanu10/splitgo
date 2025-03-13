package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		tmpl, _ := template.ParseFiles("templates/auth/login.html")
		tmpl.Execute(w,nil)
	}
	if r.Method == "POST" {
		fmt.Fprintf(w, "login submitted %s", r.Form)
	}
}
func SignUpHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		tmpl, _ := template.ParseFiles("templates/auth/signup.html")
		tmpl.Execute(w,nil)
	}
	if r.Method == "POST" {

	}
}