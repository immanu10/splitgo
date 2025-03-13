package main

import (
	"html/template"
	"log"
	"net/http"
)

type Template struct{
	templates *template.Template
}

func (t *Template) Render(w http.ResponseWriter, name string, data interface{}) error{
	return t.templates.ExecuteTemplate(w,name,data)
}

func NewTemplate() *Template{
	return &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

type LoginResponse struct{
	Success bool
	Username string
}

func newLoginResponse(success bool, username string) LoginResponse{
	return LoginResponse{
		Success: success,
		Username: username,
	}
}

func main(){
	t := NewTemplate()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t.Render(w,"index",nil)
	})	
	http.HandleFunc("/login",func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet{
			t.Render(w,"login",newLoginResponse(false,""))	
			return
		}
		if r.Method == http.MethodPost{
			username := r.FormValue("username")
			
			t.Render(w,"login",newLoginResponse(true,username))	
		}
	})	
	http.HandleFunc("/signup",func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet{
			t.Render(w,"signup",nil)
			return
		}
		if r.Method == http.MethodPost{

		}
	})	

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/",http.StripPrefix("/static/",fs))
	
	log.Fatal(http.ListenAndServe(":8080", nil))
	
}