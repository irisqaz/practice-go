package mydb

import (
	"io/ioutil"
	"log"
)

// Done contains what i accomplished today
var done = ""
var filename = "mydb.txt"

func init() {
	// f, err := os.Create(filename)
	// if err != nil {
	// 	log.Fatal("Could not open database file", filename)
	// }
	b, err := ioutil.ReadFile(filename)
	//b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal("Could not read database file", filename)
	}
	done = string(b)
}

// Get returns the database value
func Get() string {
	return done
}

// Set sets the database value
func Set(value string) {
	done = value
	ioutil.WriteFile(filename, []byte(done), 666)
}
