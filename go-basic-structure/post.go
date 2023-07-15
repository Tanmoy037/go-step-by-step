package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("This a how to create Post request")
	PerformPostRequest()
}

func PerformPostRequest() {
	const myurl = "https://jsonplaceholder.typicode.com/post" // add url here
	requestBody := strings.NewReader(`
		{
			"coursename": "crud with golang",
			"teacher": "Tanmoy",
			"price": 0
		}
	`)
	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

}
