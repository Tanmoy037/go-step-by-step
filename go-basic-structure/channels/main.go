package main

import (
	"fmt"
	"sync"
)

func main(){
	fmt.Println("Channels in golang")

	myCh := make(chan int, 5)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	// R ONLY
	go func (ch <-chan int, wg *sync.WaitGroup){
		val, isChanelOpen := <- myCh
		fmt.Println(isChanelOpen)
		fmt.Println(val)

		wg.Done()
	}(myCh, wg)
	// SEND ONLY
	go func (ch chan<- int, wg *sync.WaitGroup){
		myCh <- 5
		close(myCh)
		// myCh <- 6
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}