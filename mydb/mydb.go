package mydb

// Done contains what i accomplished today
var done = ""

// Get returns the database value
func Get() string {
	return done
}

// Set sets the database value
func Set(value string) {
	done = value
}
