package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)


func main(){
	fmt.Println("This is a Calculator that calculates fromleft to right serially")
	fmt.Println("Required Format var1 op car op var3... e.g.: 2 * 3")
	fmt.Println("Make sure you have one space with each argument provided")
	fmt.Println("|-----------------------------------------------------------------|")
	fmt.Println("Enter the Calculation: ")
	takeAllArgs()
}

func takeAllArgs(){
	scanner := bufio.NewScanner(os.Stdin)
	var inp string;
	if scanner.Scan(){
		inp =  scanner.Text()
	}
	arrOfStrNum := strings.Split(inp, " ") 


	var previous string;
	previous = arrOfStrNum[0]
	var idxcalc int
	argsLength := len(arrOfStrNum)
	fmt.Println("Length:", argsLength)
	for i:=1;i<= argsLength;{
		val := previous
		prev , index := threeOpCode(val,arrOfStrNum[i],arrOfStrNum[i+1],i)
		previous = strconv.Itoa(prev)
		idxcalc = index
		// fmt.Println("Operations: ", previous)
		if idxcalc >= (argsLength) {
			break	
		}else{
			i = idxcalc	
		}

	}

	fmt.Println("Total calculated value from left to right is: ", previous)
	
}

func filterArgs(val string) any{
	numRegEx := regexp.MustCompile(`[0-9]`)
	oprRegEx := regexp.MustCompile(`[+*-/%]`)

	if numRegEx.MatchString(val){
		num,err:=strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		return num
	}else if oprRegEx.MatchString(val){
		return val
	}

	return  nil
}

func convToInt(val1 string) int{
	// fmt.Println("Failed at conversion")
	val,err := strconv.Atoi(val1)
	if err != nil {
		panic(err)
	}
	return  val
}

func threeOpCode(val1 string , op2 string , val3 string,i int) (int,int){
	// fmt.Println("Failed at 3 OP code")
	var operation int;
	var idx int;

	 int1 := convToInt(val1)
	 int2 := convToInt(val3)

	switch op2 {
	case "+":
		operation = int1 + int2
	case "-":
		operation = int1 - int2
	case "*":
		operation = int1 * int2
	case "/":
		operation = int1 / int2
	default:
		operation = 0
	}
	idx = i+2
	return operation , idx
}

