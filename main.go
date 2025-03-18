package main

import (
	"html/template"
	"log"
	"net/http"
)

type Data struct {
	Url string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}
func main() {

	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "index.gohtml", nil)
	})

	var url string

	router.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) {
		url = r.PostFormValue("url")

		tpl.ExecuteTemplate(w, "shorten.gohtml", Data{
			Url: "localhost:8080/short",

		})
	})

	router.HandleFunc("/short", func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "short.gohtml", Data{
			Url: url,
		})
	})

	server := http.Server{
		Addr: ":8080",
		Handler: router,
	}
	log.Println("Starting server on port :8080")
	server.ListenAndServe()
}

