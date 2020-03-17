package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// URL is a url
const URL = "localhost:3306"

func main() {
	db, err := sql.Open("mysql", "root:changeme@tcp("+URL+")/practice")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	defer db.Close()

	rows, err := db.Query("select * from done;")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	var done string
	for rows.Next() {
		rows.Scan(&done)
		fmt.Println(done)
	}

}
