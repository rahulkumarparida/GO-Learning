package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	// "strings"
	"time"
)

func main() {

	fmt.Println("Welcome to pizza guy!!!")
	// takeInput()
	// formatDate()
	// pointerBehaviour()
	// arrayBehaviour()
	// slicesBehaviour()
	// mapsBehaviour()
	// structsBehave()
	gotoStatements()
}


func takeInput()  {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the rating of the pizza: ")
	inp, _ := reader.ReadString('\n')
	fmt.Println("inp:", inp)
	
}

func formatDate(){
	// Check the time package
	timeNow := time.Now()
	fmt.Println("Time: ", timeNow)
	formatedDate := timeNow.Format("01-02-2004 15:04:05 Tuesday")
	fmt.Println("Standard Format: ", formatedDate ) // MM-DD-YYYY
	// arr := strings.Split(formatedDate," ")
	// fmt.Println("Arr: ", arr[2])

	createDate := time.Date(2020,time.September , 29 , 17 , 35 , 23 , 12 , time.UTC)
	fmt.Println("Created: ", createDate)


}


// Build command
//  rahulroxx@omarchian ~/..../Basic2  main  GOOS="darwin" go build 
//  rahulroxx@omarchian ~/..../Basic2  main  go build        


func pointerBehaviour(){
	// var ptr *int
	// fmt.Println(ptr)

	myNum := 69

	secNum := &myNum

	fmt.Println("2nd: ", *secNum) //69

	myNum = 23

	fmt.Println("2nd: ", *secNum) //23

	*secNum = *secNum * 2

	fmt.Println("Num: ", myNum) // 46

}

func arrayBehaviour() {
	var food[3]string

	food[0] = "Dahibara"
	food[1] = "Biryani"
	fmt.Println("Food: ", food)
	fmt.Println("Length: ", len(food))

}


func slicesBehaviour(){
	var fruitsList = []string{"Watermelon","Mango","Orange"}

	fruitsList = append(fruitsList, "Grapes" , "Licchi")

	fmt.Println("List: ", fruitsList)
	// var favFruits []string
	// favFruits = fruitsList[1:]
	// fmt.Println("Fav list: ", favFruits)

	var myNums = make([]int , 3)

	myNums[0] = 12
	myNums[1] = 156
	myNums[2] = 38
	// myNums[3] = 14 // Throws error because index out of range total length given 
	// but if we do append
	myNums = append(myNums, 123,2344,456) // Does the Job

	fmt.Println("My Nums: ", myNums)

	// sorting it possible unlike array in here
	sort.Ints(myNums)
	fmt.Println("soreted: ", myNums)

	// remove orangr from fruitlist
	index :=  2
	fruitsList = append(fruitsList[:index], fruitsList[index+1:]... )
	fmt.Println("Fruites: ", fruitsList)

}


func mapsBehaviour() {
	var maps = make(map[string]string)

	maps["JS"] = "Javascript"
	maps["JV"] = "Java"
	maps["PY"] = "Python"
	maps["GO"] = "GOlang"

	for key , val := range maps {
		fmt.Println("Key: ",key,"Val: ", val)

	}
}


type User struct {
	Name string
	Age uint
	Email string

}
func structsBehave() {
rahulroxx := User{"Rahul",12,"rahul@email"}

fmt.Printf("The values in details are: %+v\n", rahulroxx)
	
}


func gotoStatements() {
	rougueValue := 1
	
	for  rougueValue<10 {
		fmt.Println("Val: ",rougueValue)
		rougueValue++

		if rougueValue == 5 {
			goto roxx
		}

	}

roxx:
	fmt.Println("This is label works as a we learnt as intermediate code representation")
	mapsBehaviour()

}

// continue from the functions
