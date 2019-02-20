package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hotdog", m)
}

func main() {
	d := hotdog(5)
	http.ListenAndServe("0.0.0.0:8080", d)
}
