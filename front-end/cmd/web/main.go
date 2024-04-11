package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const (
    PORT = ":3000"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, "test.page.tmpl")
	})

	fmt.Printf("Starting front end service on port %s",PORT)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Panic(err)
	}
}

func render(w http.ResponseWriter, t string) {

	partials := []string{
		"./cmd/web/templates/base.layout.tmpl",
		"./cmd/web/templates/header.partial.tmpl",
		"./cmd/web/templates/footer.partial.tmpl",
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("./cmd/web/templates/%s", t))

	for _, x := range partials {
		templateSlice = append(templateSlice, x)
	}

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
