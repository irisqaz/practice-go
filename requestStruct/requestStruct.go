package main

// Request is a minimal HTTP request
type Request struct {
	method   string
	path     string // path to resource
	protocol string
}

// RequestStruct returns a pointer to a Request
// Ex: "GET /hello.html HTTP/1.1" -> &Request{"GET", "/hello.html", "HTTP/1.1"}
func RequestStruct(str string) *Request {

	return nil
}
