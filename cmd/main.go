package main

import (
	"api-endpoints/src/login"
	"api-endpoints/src/templates"
	"fmt"
	"github.com/a-h/templ"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.Handle("/", templ.Handler(templates.Index()))
	http.Handle("/login", templ.Handler(templates.LoginPage()))

	http.HandleFunc("/api/login", login.Handler)
	fmt.Println("Server running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
