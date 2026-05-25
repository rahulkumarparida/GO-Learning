package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

var args = os.Args

type Task struct {
	CreatedAt time.Time
	Todo      string
}
type UserData struct {
	name       string	
	age        uint
	profession string
	email      string
	tasks      []Task
}

var user UserData

func main() {
	// inputUserData()
	if len(args) < 2 {
		fmt.Println("\n1.todo add\n2.todo update\n3.todo delete\n4.todo list\n5.exit")
	}
	menuOption()

}

func menuOption() {
	switch args[1] {
	case "add":
		// Create Task
		createTask()
	case "update":
		// Update Task
		updateTask()
	case "delete":
		// Delete Task
		removeTask()
	case "list":
		// List Task
		listTasks()
	default:
		// Exit
		os.Exit(1)

	}
}

func inputUserData() UserData {
	var name string
	var age uint
	var profession string
	var email string

	fmt.Println("Please provide your Firstname:")
	fmt.Scan(&name)
	fmt.Println("Please provide your age:")
	fmt.Scan(&age)
	fmt.Println("Please provide your profession:")
	fmt.Scan(&profession)
	fmt.Println("Please provide your email:")
	fmt.Scan(&email)

	user = UserData{
		name:       name,
		age:        age,
		profession: profession,
		email:      email,
	}
	return user
}

func createTask() {

	var prevTask []Task
	if _, err := os.Stat("tasks.json"); err == nil {

		fileData, err := os.ReadFile("tasks.json")
		if err != nil {
			fmt.Println("Error reading the JSON:", err)

		}
		err = json.Unmarshal(fileData, &prevTask)

		if err != nil {
			fmt.Println("Error passing the JSON:", err)
			return
		}

	}

	scanner := bufio.NewScanner(os.Stdin)

	var response string
	fmt.Println("Add the task:")
	if scanner.Scan() {
		var res = scanner.Text()
		response = res
	}

	var task = Task{
		CreatedAt: time.Now(),
		Todo:      response,
	}

	prevTask = append(prevTask, task)
	// fmt.Println("Task:", prevTask)
	user.tasks = prevTask
	
	jsondata, err := json.MarshalIndent(user.tasks, "", "  ")
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	os.WriteFile("tasks.json", []byte(jsondata), 0644)
	listTasks()	

}

func listTasks() {
	var prevTask []Task

	if _, err := os.Stat("tasks.json"); err == nil {

		fileData, err := os.ReadFile("tasks.json")
		if err != nil {
			fmt.Println("Error reading the JSON:", err)

		}
		err = json.Unmarshal(fileData, &prevTask)
		if err != nil {
			fmt.Println("Error passing the JSON:", err)
			return
		}
		user.tasks = prevTask
	}

	var tasks = user.tasks
	for idx, val := range tasks {
		fmt.Println(idx, " : ", val.Todo)
	}
}

func updateTask() {
	var prevTask []Task

	if _, err := os.Stat("tasks.json"); err == nil {

		fileData, err := os.ReadFile("tasks.json")
		if err != nil {
			fmt.Println("Error reading the JSON:", err)

		}
		err = json.Unmarshal(fileData, &prevTask)
		if err != nil {
			fmt.Println("Error passing the JSON:", err)
			return
		}
		user.tasks = prevTask
	}

	scanner := bufio.NewScanner(os.Stdin)
	var indexNum int
	listTasks()
	fmt.Print("Choose the serial number:")
	fmt.Scan(&indexNum)
	var task = user.tasks[indexNum]
	fmt.Print("Updating the task:\n", task.Todo)
	fmt.Print("\nTo: ")
	if scanner.Scan() {
		var res = scanner.Text()
		user.tasks[indexNum].Todo = res
	}

	jsondata, err := json.MarshalIndent(user.tasks, "", "  ")
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// fmt.Println("JSON:",jsondata)

	os.WriteFile("tasks.json", []byte(jsondata), 0644)

}

func removeTask() {
	var prevTask []Task

	if _, err := os.Stat("tasks.json"); err == nil {

		fileData, err := os.ReadFile("tasks.json")
		if err != nil {
			fmt.Println("Error reading the JSON:", err)

		}
		err = json.Unmarshal(fileData, &prevTask)
		if err != nil {
			fmt.Println("Error passing the JSON:", err)
			return
		}
		user.tasks = prevTask
	}

	var indexNum int
	listTasks()
	fmt.Print("Choose the serial number:")
	fmt.Scan(&indexNum)
	var task = user.tasks
	var newTasks []Task
	for index, val := range task {
		if index == indexNum {
			continue
		}
		newTasks = append(newTasks, val)
	}
	user.tasks = newTasks

	jsondata, err := json.MarshalIndent(user.tasks, "", "  ")
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// fmt.Println("JSON:",jsondata)

	os.WriteFile("tasks.json", []byte(jsondata), 0644)
}
