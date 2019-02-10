package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template
// FuncMap is used to specify functions that can be used in the template.
var funcMap = template.FuncMap{
	"fdateMDY": monthDayYear,
}

// init runs the first time a module is imported.
func init() {
	// Use template.New so we can pass a funcMap.
	tpl = template.Must(template.New("").Funcs(funcMap).ParseFiles("1_templates/3-pipelines/tplDate.gohtml"))
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tplDate.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
