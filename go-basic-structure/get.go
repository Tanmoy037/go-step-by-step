package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("This a how to create get request")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "https://jsonplaceholder.typicode.com/get/1" // add url here

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Headers: ", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

}
