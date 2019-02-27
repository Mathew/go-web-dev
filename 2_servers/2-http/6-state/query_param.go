package main

import (
	"io"
	"net/http"
)

func foo(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("q")
	io.WriteString(w, "My param: " + v)
}

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}
