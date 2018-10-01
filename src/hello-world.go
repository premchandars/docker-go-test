package main

import (
	"fmt"
	"net/http"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello world ! . This is a website server by a Go HTTP server.")
  })
  http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello world ! ")
  })
  http.ListenAndServe(":3001", nil)
}
