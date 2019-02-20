package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)


type hotdog int
var tpl *template.Template


func init() {
	tpl = template.Must(template.ParseFiles("2_servers/2-http/2-forms/index.gohtml"))
}


func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method string
		URL *url.URL
		Submissions url.Values
		Header http.Header
		ContentLength int64
		Host string
	}{
		Method: r.Method,
		URL: r.URL,
		Submissions: r.Form,
		Header: r.Header,
		ContentLength: r.ContentLength,
		Host: r.Host,
	}

	err = tpl.ExecuteTemplate(w, "index.gohtml", data)

	if err != nil {
		log.Panic(err)
	}
}

func main() {
	http.ListenAndServe("0.0.0.0:8080", hotdog(0))
}
