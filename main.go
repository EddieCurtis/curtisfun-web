package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*"))

type Page struct {
	Title string
}

func createPage(title string) (*Page, error) {
	return &Page{Title: title}, nil
}

func render(w http.ResponseWriter, p *Page) {
	// you access the cached templates with the defined name, not the filename
	err := templates.ExecuteTemplate(w, p.Title, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	p, _ := createPage("index")
	render(w, p)
}

func pgpHandler(w http.ResponseWriter, r *http.Request) {
	p, _ := createPage("pgp")
	render(w, p)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/pgp", pgpHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}