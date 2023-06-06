package main

import (
  "fmt"
  "log"
  "net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, this is Docker T41K41")
}

func about(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, this is Docker T41K41")
}

func main() {
  http.HandleFunc("/", index)
  http.HandleFunc("/about", about)
  log.Println("application started at port :8000")
  log.Fatal(http.ListenAndServe(":8000", nil))
}