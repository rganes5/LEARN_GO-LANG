package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world %s", r.URL.Path)
}

func main() {
	http.HandleFunc("/hello", handlerFunc)
	http.ListenAndServe(":9999", nil)
}
