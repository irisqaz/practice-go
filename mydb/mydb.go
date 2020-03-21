package mydb

import (
	"database/sql"
	"fmt"
	"os"
)

// Done contains what i accomplished today
// var done = ""
// var filename = "mydb.txt"

// func init() {
// 	// f, err := os.Create(filename)
// 	// if err != nil {
// 	// 	log.Fatal("Could not open database file", filename)
// 	// }
// 	b, err := ioutil.ReadFile(filename)
// 	//b, err := ioutil.ReadAll(f)
// 	if err != nil {
// 		log.Fatal("Could not read database file", filename)
// 	}
// 	done = string(b)
// }

// Get returns the database value
// func Get() string {
// 	return done
// }

var db *sql.DB

// InitDB initialize the db instance
func InitDB(d *sql.DB) {
	db = d
}

// Close closes the db instance
func Close() {
	db.Close()
}

// Add sets the database value
func Add(value string) {
	// f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	log.Fatal("Could not read database file", filename)
	// }
	// defer f.Close()

	// _, err = f.WriteString(value + "\n")
	// if err != nil {
	// 	log.Fatal("Could not append database file", filename)
	// }
	_, err := db.Query("insert into done (task) values('" + value + "');")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

}

// List gets list of tasks done
func List() []string {
	// var l []string
	// f, err := os.Open(filename)
	// if err != nil {
	// 	log.Fatal("Could not read database file", filename)
	// }
	// defer f.Close()

	// scanner := bufio.NewScanner(f)
	// for scanner.Scan() {
	// 	l = append(l, "<li>"+scanner.Text()+"</li>")
	// }

	// return l
	var list []string

	rows, err := db.Query("select * from done;")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
		return list
	}

	var task string
	for rows.Next() {
		rows.Scan(&task)
		// fmt.Println(done)
		list = append(list, task)
	}

	return list
}
