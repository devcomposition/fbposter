package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	// define an handler to display our pages
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, "home.page.gohtml")
	})

	// start the web server
	fmt.Println("Starting front end service on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panic(err)
	}
}

//go:embed templates
var templateFS embed.FS

func render(w http.ResponseWriter, s string) {
	partials := []string{
		"templates/base.layout.gohtml",
		"templates/header.partial.gohtml",
		"templates/footer.partial.gohtml",
	}

	// append the template we received as a parameter
	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("templates/%s", s))

	for _, x := range partials {
		templateSlice = append(templateSlice, x)
	}

	// parse the templates
	tmpl, err := template.ParseFS(templateFS, templateSlice...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// execute the template
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
