package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", HelloServer)
	http.HandleFunc("/bye", ByeServer)
	http.ListenAndServe(":8080", nil)
}
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}
func ByeServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bye!  Have a nice day!")
}
