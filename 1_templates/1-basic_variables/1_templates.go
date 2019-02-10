package main

import (
	"log"
	"os"
	"text/template"
)

type number struct {
	Word   string
	Number string
}

var tpl *template.Template

// init runs the first time a module is imported.
func init() {
	// load templates.
	// template.Must will raise a panic if the template pattern doesn't match.
	tpl = template.Must(template.ParseGlob("1_templates/1-basic_variables/*.gohtml"))
}

func main() {
	basic()
	slice()
	mapp()
	structt()
}

func basic() {
	// pick template to render and pass the data.
	err := tpl.ExecuteTemplate(os.Stdout, "basic.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}
}

func slice() {
	_ = tpl.ExecuteTemplate(os.Stdout, "slice.gohtml", []string{"One", "Two", "Three"})
}

func mapp() {
	_ = tpl.ExecuteTemplate(os.Stdout, "map.gohtml", map[string]string{
		"One":   "1",
		"Two":   "2",
		"Three": "3",
	})
}

func structt() {
	_ = tpl.ExecuteTemplate(os.Stdout, "struct.gohtml", []number{
		{Word: "One", Number: "1"},
		{Word: "Two", Number: "2"},
		{Word: "Three", Number: "3"},
	})
}
