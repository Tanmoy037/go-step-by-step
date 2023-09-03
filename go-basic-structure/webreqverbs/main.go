package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("This is Webrequest handeling code")
	PerformGetRequest()
}
func PerformGetRequest(){
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	fmt.Println("Status code: ",response.StatusCode)
	fmt.Println("Content length is: ",response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
