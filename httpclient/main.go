package main

import (
	"fmt"
	"net/http"

	"github.com/irisqaz/practice-go/myhttp"
	"github.com/irisqaz/practice-go/mypkg"
)

func main() {
	url := mypkg.Prompt("URL")

	// create http request
	req, _ := http.NewRequest("GET", url, nil)

	if req.Host == "" {
		fmt.Println("Incorrect URL provided")
		return
	}

	fmt.Println()
	fmt.Println("   Client will send this Request: ")
	myhttp.PrintReq(req)

	fmt.Println("Connecting to", req.Host, "and sending request for", req.RequestURI)
	fmt.Println()

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer resp.Body.Close()
	fmt.Println("   Server sent this response: ")

	myhttp.PrintResp(resp)

}
