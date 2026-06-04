package main

import (
	"fmt"
	"sync"

)

// Simulating Go Routines
var wg sync.WaitGroup
var mut sync.Mutex

func main()  {
	
	// go greeter("Hello")
	// greeter("World")

	// websiteList := []string{
	// 	"https://github.com",
	// 	"https://meesho-clone-pink.vercel.app",
	// 	"https://google.com",
	// 	"https://airbnbapi-6s4b.onrender.com",
	// }
	// var weblist []string
	// for _ , url := range websiteList{
	// 	go getStatusCode(url)
	
	// 	wg.Add(1)
	// 	weblist = append(weblist, url)
	// }
	// wg.Wait()
	// fmt.Println("WebList: ", weblist)		

	fmt.Println("Race Conditions in GoRoutines")
	wg := &sync.WaitGroup{} // Using only this multiple race condition occur because all of the routines are trying to access the memory(same) at once raises race condition
	mut := &sync.Mutex{} //Using mutex locks the memory address untill the first go routine is completed , same reason why without mutex the wg executed not sequentially but mutex holds the door(memory)
	// go run --race main.go/. :- to see the race errors	

	var score = []int{0}
	
	wg.Add(3)	// Because 3 goroutines
	go func (wg *sync.WaitGroup , mut *sync.Mutex)  {
		mut.Lock()
		fmt.Println("1st routine")
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg,mut)
	go func (wg *sync.WaitGroup, mut *sync.Mutex)  {
		mut.Lock()
		fmt.Println("2nd routine ")
		score = append(score, 2)
		mut.Unlock()
		wg.Done()

		}(wg,mut)
	go func (wg *sync.WaitGroup, mut *sync.Mutex)  {
		mut.Lock()
		fmt.Println("3rd routine ")
		score = append(score, 3)
		mut.Unlock()
		wg.Done()

	}(wg,mut)

		wg.Wait()

	fmt.Println("Score:", score)

}


// func greeter(s string)  {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(30*time.Millisecond)
// 		fmt.Println(s)
// 	}
// }



// func getStatusCode(endpoint string)  {
// 	defer wg.Done()
// 	res ,err := http.Get(endpoint)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("%d status for website %v \n", res.StatusCode, endpoint)
// }