package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", myHandler)
	http.HandleFunc("/by", myHandler2)
	http.ListenAndServe(":8080", nil)
}

func myHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "<html><body><h1>Hello, world!</h1></body></html>")
}
func myHandler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><body><h1>Good bye, world!</h1></body></html>")
}
