package main

import (
	"fmt"
	// "math/rand"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)


func main(){
	// go TwoMessageSimultaneoulsy("Hello")
	// TwoMessageSimultaneoulsy("World")
	
	// wg.Add(3)	
	// go ConcurretnCounter(10 , "func1")
	// go ConcurretnCounter(10 , "func2")
	// // defer wg.Wait()
	// ConcurretnCounter(10 , "func3")
	
	
	//  go SleepCount(1)	
	//  go SleepCount(2)	
	//  SleepCount(3)
	
	// wg := sync.WaitGroup{}		
	// students :=[]string{
	// 	"Rahul","Satya","Tushar","Siddharth",
	// }

	// wg.Add(len(students))	
	// for _, val := range students {
	// 	go AssignmentChecker(val ,&wg)
	// }

	// wg.Wait()
	// fmt.Println("Finished Chceking all the assigmenets")



	// resultChannel()

	bankAccount()

}

// Execution is random beacuse both the functions are executed at the same time , The faster one is logged first , When Go is used because invoking
// World
// Hello
// World
// Hello
// Hello
// World
// World
// Hello
// World
// Hello
// Hello
// World
func TwoMessageSimultaneoulsy(str string){

	data := string("Hello there this is runnign on backgroung")
	os.WriteFile("con.txt", []byte(data), 0644)
	for i := 0; i < 6; i++ {
		time.Sleep(20 * time.Millisecond)	
		fmt.Println(str)
	}
	// wg.Done()
}



func ConcurretnCounter(num int , funcName string) {

	for i := 0; i < num; i++ {
		time.Sleep(200* time.Millisecond)	
		fmt.Println(funcName , i)
	}
	// wg.Done()
}


func SleepCount(num time.Duration){

	fmt.Printf("%v work finished\n ", num)	
	time.Sleep(num * time.Second)						
	data := string("Hello there this is running on background"+strconv.Itoa(int(num))+"\n")
	os.WriteFile("test.txt", []byte(data), 0644)



}


func AssignmentChecker(name string, wg *sync.WaitGroup)  {
	fmt.Println("Checking ", name)
	wg.Done()
}

func DoSomeWork(num int) string{
	pwd ,_ := os.Getwd()	
	// val := time.Duration(rand.Intn(5))
	data := string("Hello there this is running on background"+strconv.Itoa(int(5))+"\n")	
	dirPath := filepath.Join(pwd, "tests")
	time.Sleep(5 * time.Second)
	os.Mkdir(dirPath,0755)
	filename := filepath.Join(dirPath,("test"+strconv.Itoa(num)+".txt"))
	fmt.Println(filename)
	err := os.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		panic(err)
	}
	return data
}
func resultChannel(){
	dataChan := make(chan string)
	wg := sync.WaitGroup{}
	go func ()  {
		numOfFile := 2000 // be carful choosing this number either way i tried 2000 files and it created those within 1 second as soon as the time.sleep ended after 5 second took a second and wrote all those through all the go threads worked at the same time.
		for i := 0; i < numOfFile; i++ {
			wg.Add(1)
			go func(){
				defer wg.Done()
				res:=  DoSomeWork(i)
				dataChan<-res

			}()	
		}
		wg.Wait()
		close(dataChan)
	}()

	for val := range dataChan {
		// r := <-dataChan
		fmt.Println("Val: ",val)
	}
}


func deposit(cash int) int{
	fmt.Println("Cash Deposit: " , cash)
	return  cash+100
}

var currAmount = 1000
func bankAccount(){

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(){
			defer wg.Done()
			mu.Lock()
			currAmount = deposit(currAmount)
			mu.Unlock()
			}()
	}
		
	wg.Wait()	
	fmt.Println("Total: ", currAmount)
	



}