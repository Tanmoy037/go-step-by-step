package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)


func main(){
	fmt.Println("Generating random number using crypto: ")
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)
}