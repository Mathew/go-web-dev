package main

import (
	"os"
	"text/template"
)

var tpl *template.Template
// FuncMap is used to specify functions that can be used in the template.

// init runs the first time a module is imported.
func init() {
	// Use template.New so we can pass a funcMap.
	tpl = template.Must(template.ParseGlob("1_templates/4-templates/*.gohtml"))
}

func main() {
	_ = tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", nil)
}
