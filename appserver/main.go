package main

import (
	"fmt"
	"net/http"

	"github.com/irisqaz/practice-go/mydb"
)

var done string

var doc = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
<form action="/" method="post">
Done Today: <input type="text" name="done" value="%s"><br>
<input type="submit">
</form>
</body>
</html>
`

func main() {
	http.HandleFunc("/", handler1)
	http.ListenAndServe(":8080", nil)
}

func handler1(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		get(rw)
	case "POST":
		post(rw, req)
	}

}

func get(rw http.ResponseWriter) {
	done = mydb.Get()
	fmt.Fprintf(rw, doc, done)
}

func post(rw http.ResponseWriter, req *http.Request) {
	done = req.FormValue("done")
	mydb.Set(done)
	fmt.Fprintf(rw, doc, done)
}
