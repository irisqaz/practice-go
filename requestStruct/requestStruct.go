package main

// Request is a minimal HTTP request
type Request struct {
	method   string
	path     string // path to resource
	protocol string
}

// NewRequest returns a pointer to a Request
// Ex: "GET /hello.html HTTP/1.1" -> &Request{"GET", "/hello.html", "HTTP/1.1"}
func NewRequest(str string) *Request {

	return nil
}
