package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./2_servers/2-http/5-serve-files/static"))))
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(writer, `<img src="/static/yougotit.gif"/>`)
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}

}
