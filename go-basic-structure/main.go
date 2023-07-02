package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"time"
)
func add(x,y int) int {
	return x+y
}

func main(){
	//variable
	var name = "Tanmoy"
	fmt.Println("Hello from Go")
	fmt.Println("Hello",name)
	//for loop
	for i := 0; i<5; i++ {
		fmt.Println(i)
	}
	// if else
	x := 10
	if x < 5 {
		fmt.Println("less than 5")
	}else{
		fmt.Println("greater than 5")
	}
	// function



		result := add(4,6)
		fmt.Println(result)
 

	// taking input from user 
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Enter your name: ")

	// input, _ := reader.ReadString('\n')
	// fmt.Println("Hello",input)

	// Conversions
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter you rating for our app 1 to 5: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating ",input)
	fmt.Printf("Type of this rating is %T",input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input),64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("\nAdded 1 to your rating: ", numRating+1)
	}

	// Time Handling
	presentTime := time.Now()

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	// Go build (Need to run in terminal)
	// GOOS="windows" go build
	// GOOS="linux" go build	


	// pointer

	myNumber := 23
	
	var ptr = &myNumber

	fmt.Println("Address of myNumber is: ", ptr)
	fmt.Println("Value of myNumber is: ", *ptr)

	*ptr = *ptr + 2

	fmt.Println("New value of myNumber is: ", myNumber)

	//array
	// var fruitList[4]string	

	// fruitList[0] = "Apple"
	// fruitList[1] = "Banana"
	// fruitList[2] = "Orange"
	
	// fmt.Println(fruitList)



	// slices
	var fruitList = []string{"Apple","Banana","Orange"}
	fmt.Printf("Type of fruitlist is %T\n",fruitList)

	fruitList = append(fruitList,"Mango","Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	////using sort function we can sort a slice

	//how to remove a value from a slice based on index

	var courses = []string{"Go","Python","Java","C++","ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index],courses[index+1:]...)
	fmt.Println(courses)



	

	
}