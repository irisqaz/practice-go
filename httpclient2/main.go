package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// url := mypkg.Prompt("URL")
	url := "http://localhost:8080/"
	// create http request
	//req, _ := http.NewRequest("GET", url, nil)

	// if req.Host == "" {
	// 	fmt.Println("Incorrect URL provided")
	// 	return
	// }

	// fmt.Println()
	// fmt.Println("   Client will send this Request: ")
	// myhttp.PrintReq(req)

	// fmt.Println("Connecting to", req.Host, "and sending request for", req.RequestURI)
	// fmt.Println()

	//client := http.Client{}
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer resp.Body.Close()
	// fmt.Println("   Server sent this response: ")

	// myhttp.PrintResp(resp)
	io.Copy(os.Stdout, resp.Body)

	// contentType := resp.Header.Get("Content-Type")
	// if strings.Contains(contentType, "html") {
	// 	handle(resp)
	// }
}

func handle(resp *http.Response) {
	fmt.Println("will handle the response")
	document, _ := goquery.NewDocumentFromReader(resp.Body)
	document.Find("a").Each(handleElement)
}

func handleElement(index int, element *goquery.Selection) {
	str, exists := element.Attr("href")
	if exists {
		fmt.Println(str)
	}
}
