package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type Mapping struct {
	Content string
	Layout  string
}

var mapping = []Mapping{
	Mapping{"index", "base"},
	Mapping{"about", "base"},
	Mapping{"contact", "simple"},
}

func main() {
	var templMap = make(map[string]*template.Template)

	for _, m := range mapping {
		contentFile := fmt.Sprintf("%s.page.tmpl", m.Content)
		contentPath := filepath.Join(".", "template", contentFile)
		layoutFile := fmt.Sprintf("%s.layout.tmpl", m.Layout)
		layoutPath := filepath.Join(".", "template", layoutFile)
		t, err := template.ParseFiles(contentPath, layoutPath)
		if err != nil {
			log.Fatalf("Error when parsing files %s and %s: %v", m.Content, m.Layout, err)
		}
		templMap[m.Content] = t

	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templMap["index"].Execute(w, nil)
	})
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		templMap["about"].Execute(w, nil)
	})
	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		templMap["contact"].Execute(w, nil)
	})
	http.ListenAndServe(":8080", nil)
}
