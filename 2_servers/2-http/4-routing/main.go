package main

import (
	"io"
	"log"
	"net/http"
)

func d(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "dog")
}

func c(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "cat")
}

func main() {
	// matches any variant of /dog/ due to trailing slash (/dog/lotsof/)
	http.HandleFunc("/dog/", d)
	// Only matches /cat
	http.HandleFunc("/cat", c)

	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Panic(err)
	}
}
