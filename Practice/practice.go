package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	// sumOfArr()
	// findLargest()
	// totalEvenNum()
	// fizzBuzz()	
	// fmt.Print("\nEnter a Sring: ")
	// scanner := bufio.NewScanner(os.Stdin)
	// var inp string
	// if scanner.Scan() {
		// 	inp = scanner.Text()
		// }
		
		// reverseString(inp)
		// checkPalindrome(inp)
		// linearSearch(inp)
		// frequncyCounter(inp)
		// removeDuplicates(inp)
		twoSum()	
}

func sumOfArr() {
	fmt.Println("Add Number as Inputs to add the element!!")

	fmt.Println("Enter Input Num: ")

	reader := bufio.NewScanner(os.Stdin)
	var inp string
	if reader.Scan() {
		inp = reader.Text()

	}

	fmt.Println("Input: ", inp)
	arrOfStrNum := strings.Split(inp, " ")

	var sum int
	for _, val := range arrOfStrNum {
		value, err := strconv.Atoi(val)

		if err != nil {
			panic(err)
		}

		sum += value

	}

	fmt.Println("Sum: ", sum)

}

func findLargest() {
	fmt.Println("To find the largest number among the provided arr")
	scanner := bufio.NewScanner(os.Stdin)
	var inp string
	if scanner.Scan() {
		inp = scanner.Text()
		fmt.Println("Input: ", inp)
	}
	arrOfStrNum := strings.Split(inp, " ")
	var max int = 0

	for _, val := range arrOfStrNum {

		value, err := strconv.Atoi(val)

		if err != nil {
			panic(err)
		}

		if value > max {
			max = value
		}
	}
	fmt.Println("Max: ", max)

}

func totalEvenNum() {
	fmt.Println("To find the total even number among the provided arr")
	scanner := bufio.NewScanner(os.Stdin)
	var inp string
	if scanner.Scan() {
		inp = scanner.Text()
		fmt.Println("Input: ", inp)
	}
	arrOfStrNum := strings.Split(inp, " ")
	var num int = 0

	for _, val := range arrOfStrNum {

		value, err := strconv.Atoi(val)

		if err != nil {
			// panic(err)
			fmt.Println("Re-enter again with one spacing after each number e.g: 12 12 3 324 ")

		}

		if value%2 == 0 {
			num += 1
		}
	}
	fmt.Println("Total even: ", num)

}

func reverseString(inp string) string {
	// fmt.Print("\nEnter a Sring: ")
	// scanner := bufio.NewScanner(os.Stdin)

	// if scanner.Scan() {
	// 	inp = scanner.Text()
	// }
	fmt.Println("Given String: ", inp)
	arrOfStr := strings.Split(inp, "")
	slices.Reverse(arrOfStr)
	var str string = strings.Join(arrOfStr, "")

	fmt.Println("Revese String: ", str)
	return str
}

func checkPalindrome(inp string) {

	revValue := reverseString(inp)

	if revValue == inp {
		fmt.Println("Yes it is a palindrome.")
	} else {
		fmt.Println("No it is not a palindrome")
	}

}

func linearSearch(inp string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the Number you want to search: ")
	arrOfStr := strings.Split(inp, " ")
	fmt.Println("Among these: ", arrOfStr)
	var searchInput string
	if scanner.Scan() {
		searchInput = scanner.Text()
	}

	for idx, val := range arrOfStr {
		if val == searchInput {
			fmt.Println("The element is in index : ", idx)
			break
		}
	}

}

func frequncyCounter(inp string) {
	counter := make(map[string]int)

	for _, val := range inp {
		counter[string(val)]++

	}

}

func removeDuplicates(inp string) {
	var newString string

	for _, val := range inp {

		if !strings.Contains(newString, string(val)) {
			newString += string(val)

		} else {
			continue
		}

	}
	fmt.Println("New String: ", newString)

}

func fizzBuzz() {

	for i := range 100{
		if i == 0{
			continue
		}else if i%3 == 0 && i%5== 0{
			fmt.Println(i,":","FizzBuzz")
		}else if i%5 == 0{
			fmt.Println(i,":","Buzz")
		}else if i%3 == 0{
			fmt.Println(i,":","Fizz")
		}else{
			continue
		}
	}
}


func twoSum() {
	scanner := bufio.NewScanner(os.Stdin)
	var inpString string
	var total int
	fmt.Println("Enter the number array: ")	
	
	if scanner.Scan(){
		inpString = scanner.Text()
	}

	fmt.Println("Enter the target value: ")
	
	if scanner.Scan(){
		target  := scanner.Text()
		tar,err := strconv.Atoi(target)
		if err != nil {
			panic(err)
		}
		total = tar
	}

	arrOfStrInt := strings.Split(inpString, " ")
	var arrOfNumInt []int
	for _,val := range arrOfStrInt {
		val,err:= strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		arrOfNumInt = append(arrOfNumInt, val )
	}
	arrOfSum := []int{}
	for idx := range arrOfNumInt {
		for jdx := range arrOfNumInt{
			if idx == jdx {
				continue
			}
			if arrOfNumInt[idx] + arrOfNumInt[jdx] == total {
				arrOfSum = append(arrOfSum, arrOfNumInt[idx])
				arrOfSum = append(arrOfSum, arrOfNumInt[jdx])
				break
			}	
		}
		if len(arrOfSum) > 0 {
			break
		}

	}
	fmt.Println("Answer: ",arrOfSum)
}


