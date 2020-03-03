package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/irisqaz/practice-go/myhttp"

	"github.com/PuerkitoBio/goquery"
)

// URL is a url
const URL = "http://localhost:8080/"

func main() {
	reader := csv.NewReader(os.Stdin)
	reader.Comma = ' '
	var command []string
	for {
		fmt.Println("Enter command to send (or exit)")
		//fmt.Scanln(&command)
		command, _ = reader.Read()
		if command[0] == "exit" {
			break
		}
		switch command[0] {
		case "GET":
			get()
		case "POST":
			post(command[1])
		default:
		}
	}
}

func get() {
	resp, err := http.Get(URL)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}

func post(value string) {
	data := url.Values{}
	data.Set("done", value)
	reader := strings.NewReader(data.Encode())

	req, _ := http.NewRequest("POST", URL, reader)
	if req.Host == "" {
		fmt.Println("Incorrect URL provided")
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	fmt.Println("*** client will send this request:")

	myhttp.PrintReq(req)
	fmt.Println()

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println()
	//handle(resp)
	io.Copy(os.Stdout, resp.Body)

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
