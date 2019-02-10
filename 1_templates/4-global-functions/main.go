package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var tplPipe *template.Template
// FuncMap is used to specify functions that can be used in the template.
var funcMapPipe = template.FuncMap{
	"fdbl":  double,
	"fsq":   square,
	"fsqrt": sqRoot,
}

// init runs the first time a module is imported.
func init() {
	// Use template.New so we can pass a funcMap.
	tplPipe = template.Must(template.New("").Funcs(funcMapPipe).ParseFiles("1_templates/3-pipelines/tplPipe.gohtml"))
}

func double(num float64) float64 {
	return num * 2
}

func square(num float64) float64 {
	return math.Pow(num, 2)
}

func sqRoot(num float64) float64 {
	return math.Sqrt(num)
}

func main() {
	err := tplPipe.ExecuteTemplate(os.Stdout, "tplPipe.gohtml", 3.0)
	if err != nil {
		log.Fatalln(err)
	}
}
