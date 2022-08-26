package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, "main.page.html")
	})

	fmt.Println("Web started on 8081 port!")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Panic(err)
	}
}

//go:embed template
var templateFS embed.FS

func render(w http.ResponseWriter, s string) {
	partials := []string{
		"template/base.layout.html",
		"template/header.html",
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("template/%s", s))

	for _, x := range partials {
		templateSlice = append(templateSlice, x)
	}

	tmpl, err := template.ParseFS(templateFS, templateSlice...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data struct {
		BrokerURL string
	}

	data.BrokerURL = os.Getenv("BROKER_URL")

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
