package handler

import (
	"html/template"
	"log"
	"net/http"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w)
}

var templates = template.Must(template.ParseFiles("profile.html"))

func renderTemplate(w http.ResponseWriter) {
	err := templates.ExecuteTemplate(w, "profile.html", "Hello World!")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Handler() {
	http.HandleFunc("/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
