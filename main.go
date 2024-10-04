package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	mainPage := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))

		tmpl.Execute(w, nil)
	}

	http.HandleFunc("/", mainPage)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
