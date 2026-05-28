package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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
	// gotoStatements()
	// fileHandling()
	exploringOs()
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
		if key == "JS"{
			defer fmt.Println("Key: ",key,"Val: ", val)
			continue
		}
		fmt.Println("Key: ",key,"Val: ", val)

	}
}

func structsBehave() {
rahulroxx := User{"Rahul",12,"rahul@email"}
rahulroxx.getEmail()
fmt.Printf("The values in details are: %+v\n", rahulroxx)

rahulroxx = rahulroxx.randomizeValues()
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

type User struct {
	Name string
	Age uint
	Email string

}

func (u User) getEmail() {
	fmt.Println("This is the required Email: ",u.Email)
}

func (u User) randomizeValues() User{
	u.Age += 10
	u.getEmail()
	u.Email = "New@gmail.com"
	u.Name = "Rahul"
	return u
}



func fileHandling() {
	// create a txt
	file,err := os.Create("./Hello.txt")
	if err != nil {
		panic(err)
	}
	content := "This is a content written to the file, Hello File."
	length , err := io.WriteString(file,content)

	if err != nil {
		panic(err)
	}

	fmt.Println("Length of the file: ",length)
	defer file.Close()
	readFile("./Hello.txt")
}

func readFile(file string){
	contentByte , err := os.ReadFile(file)
		if err != nil {
		panic(err)
	}
	fmt.Println("Conetnt: ",string(contentByte))
	
	// return  content 
}

func exploringOs()  {
	mkdir := os.Mkdir("test",0755)
	if mkdir != nil {
		fmt.Println("Erro:",mkdir)
	}
	chdir:= os.Chdir("./test")
	if chdir != nil {
		fmt.Println("Erro:",chdir)
	}
	file , err:= os.Create("./Testting.txt")
	if err != nil {
		fmt.Println("Erro:",err)
	}
	text := ("This is is a testing File i created by first creating a folder named test\nthen i changed the directory to test and created and writing in thisn file ")
	length , err := io.WriteString(file,text)
	if err != nil {
		fmt.Println("Erro:",err)
	}
	fmt.Println("Length: ",length)
	// Because we changed the directory that is why giving the filename directly works  
	defer readFile("./Testting.txt")
}

