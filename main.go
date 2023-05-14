package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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



	

	
}