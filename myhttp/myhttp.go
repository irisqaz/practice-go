package myhttp

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

// PrintReq dumps the request as a string
func PrintReq(req *http.Request) {
	reqBytes, _ := httputil.DumpRequest(req, true)

	fmt.Println()
	fmt.Printf("%v \n", string(reqBytes))
}

// PrintResp dumps the response as a string
func PrintResp(resp *http.Response) {
	reqBytes, _ := httputil.DumpResponse(resp, true)

	fmt.Println()
	fmt.Printf("%v \n", string(reqBytes))
}
