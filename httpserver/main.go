package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", handler1)
	http.ListenAndServe(":8080", nil)
}

func handler1(resp http.ResponseWriter, req *http.Request) {

}
