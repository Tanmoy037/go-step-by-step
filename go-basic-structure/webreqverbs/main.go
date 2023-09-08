package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("This is Webrequest handeling code")
	PerformGetRequest()
	PerformPostRequest()
	PerformPostFormRequest()
}
func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostRequest() {
	const myurl = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
		{
			"coursename":"Let's go with golang",
			"price": 0,
			"platform" : "gobyexample"


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

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	data := url.Values{}

	data.Add("firstname", "Tanmoy")
	data.Add("lastname", "Santra")
	data.Add("email", "tanmoy@aegiscovenant.com")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
