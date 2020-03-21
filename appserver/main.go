package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/irisqaz/practice-go/mydb"

	_ "github.com/go-sql-driver/mysql"
)

const dburl = "localhost:3306"

func init() {
	var err error
	db, err := sql.Open("mysql", "root:changeme@tcp("+dburl+")/practice")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	mydb.InitDB(db)
}

var done []string

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
Done Today: <input type="text" name="done" value=""> 
<input type="submit">
</form>
<ul>
%s
</ul>
</body>
</html>
`

func main() {
	fmt.Println("Appserver listening on port 8080 ...")
	http.HandleFunc("/", handler1)
	http.ListenAndServe(":8080", nil)
	mydb.Close()
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
	l := mydb.List()
	reverse(&l)
	str := listView(l)
	fmt.Fprintf(rw, doc, str)
}

func post(rw http.ResponseWriter, req *http.Request) {
	newDone := req.FormValue("done")
	mydb.Add(newDone)
	l := mydb.List()
	reverse(&l)
	str := listView(l)
	fmt.Fprintf(rw, doc, str)
}

func listView(l []string) string {
	str := strings.Join(l, "</li><li>")
	str = "<li>" + str
	return str
}

func reverse(list *[]string) {
	l := *list
	last := len(l) - 1
	for i := 0; i <= last; i++ {
		l[i], l[last] = l[last], l[i]
		last = last - 1
	}
}
