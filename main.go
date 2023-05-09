package main

import "fmt"

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


	

	
}