package main

import (
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template
// FuncMap is used to specify functions that can be used in the template.
var funcMap = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

// init runs the first time a module is imported.
func init() {
	// Use template.New so we can pass a funcMap.
	tpl = template.Must(template.New("").Funcs(funcMap).ParseGlob("1_templates/2-functions/*.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	return s[:3]
}

func main() {
	_ = tpl.ExecuteTemplate(os.Stdout, "tplDate.gohtml", "Hello World!")
}
