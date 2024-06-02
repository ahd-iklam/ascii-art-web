package main

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("templates/index.html", "templates/result.html"))

func main() {
	http.HandleFunc("/", excuteInexPage)
	http.HandleFunc("/ascii-art", excuteAsciiArtResult)
	http.ListenAndServe(":8080", nil)
}
