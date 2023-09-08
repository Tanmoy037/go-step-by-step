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
	// EncodeJson()
	DecodeJson()
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

func DecodeJson(){
	jsonDataFromWeb := []byte(`
	{
		
		"coursename": "ReactJs Bootcamp",
		"Price": 399,
		"website": "udemy.com",
		"tags": ["web-dev","js"]
	}
	`)

	var udemyCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &udemyCourse)
		fmt.Printf("%v\n", udemyCourse)
	} else {
		fmt.Println ("JSON WAS NOT VALID")
	}

	// some cases where you just want to add data to key value 

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n",myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is: %T\n", k, v, v)
	}


}








