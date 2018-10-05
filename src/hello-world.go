package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting http server")
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello world ! . This is a website server by a Go HTTP server.")
  })
  http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello world ! ")
  })
	fmt.Println("Server started in port 3001")
  http.ListenAndServe(":3001", nil)
}
