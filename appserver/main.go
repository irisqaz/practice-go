package main

import (
	"fmt"
	"net/http"

	"github.com/irisqaz/practice-go/mydb"
)

var done string

func main() {
	http.HandleFunc("/", handler1)
	http.ListenAndServe(":8080", nil)
}

func handler1(body http.ResponseWriter, req *http.Request) {
	doc := `<!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <meta http-equiv="X-UA-Compatible" content="ie=edge">
        <title>Document</title>
    </head>
    <body>
    <p>%s</p>
    </body>
    </html>
    `
	done = mydb.Done
	fmt.Fprintf(body, doc, done)
}
