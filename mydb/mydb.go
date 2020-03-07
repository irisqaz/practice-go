package mydb

import (
	"bufio"
	"log"
	"os"
)

// Done contains what i accomplished today
var done = ""
var filename = "mydb.txt"

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
func Get() string {
	return done
}

// Add sets the database value
func Add(value string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Could not read database file", filename)
	}
	defer f.Close()

	_, err = f.WriteString(value + "\n")
	if err != nil {
		log.Fatal("Could not append database file", filename)
	}

}

// List gets list of tasks done
func List() []string {
	var l []string
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal("Could not read database file", filename)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		l = append(l, "<li>"+scanner.Text()+"</li>")
	}

	return l
}
