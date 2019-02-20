package main

import (
	"html/template"
	"log"
	"net/http"
)

type hotdog int

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("2_servers/2-http/3-headers/index.gohtml"))
}

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Made-up_header", "A made up header")

	data := struct {
		Header http.Header
	}{
		Header: w.Header(),
	}

	err := tpl.ExecuteTemplate(w, "index.gohtml", data)

	if err != nil {
		log.Panic(err)
	}
}

func main() {
	err := http.ListenAndServe("0.0.0.0:8080", hotdog(0))
	if err != nil {
		log.Panic(err)
	}
}
