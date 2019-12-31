package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloServer)
	http.HandleFunc("/bye", byeServer)
	http.ListenAndServe(":8080", nil)
}
func helloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}
func byeServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bye!  Have a nice day!")
}
