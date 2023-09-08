package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int	
	Platform string `json:"website"`
	Password string	`json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Json tutorial")
	EncodeJson()
}

func EncodeJson(){
	udemyCourses := []course{
		{"ReactJs Bootcamp",399, "udemy.com","abc123",[]string{"web-dev", "js"}},
		{"NodeJs Bootcamp",499, "udemy.com","qwe123",[]string{"web-dev", "js"}},
		{"Typescript",399, "udemy.com","pqr123", nil },
	}

	finalJson, err := json.MarshalIndent(udemyCourses, "","\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

}