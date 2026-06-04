package main

import (
	"fmt"
	"sync"
)

func main()  {
	wg := &sync.WaitGroup{}

	fmt.Println("Channels")

	myChannel := make(chan int , 2)

	// myChannel <- 4
	// fmt.Println("Channel: ", <-myChannel)

	wg.Add(2)
		// <-chan Show the data is going outside the channel (While reading)
	go func(ch <-chan int, wg *sync.WaitGroup){
		// Number of reader should match the number of writer else we use buffer in while declaring the channle mentioning the expected number of writer it will ignore those values but wont throw an error at atleast
		// fmt.Println("Channel: ", <-ch)	
		val , isChannelOpen := <-myChannel
		fmt.Println("Channel: ", isChannelOpen , " val: ", val)	
		wg.Done()	
	}(myChannel,wg)
			// chan<- Show the data is going inside the channel (While writing)
	go func(ch chan<- int, wg *sync.WaitGroup ){
		// close(myChannel)
		ch <- 4
		// ch <- 2
		// close(myChannel)
		wg.Done()	
	}(myChannel,wg)
	
	wg.Wait()
	
}