package main

import (
	"api-endpoints/templates"
	"fmt"
	"github.com/a-h/templ"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", templ.Handler(templates.Index()))
	fmt.Println("Server running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
