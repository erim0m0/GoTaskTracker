package main

import (
	"crud/crud"
	"crud/tools"
	"fmt"
)

// get inputs from user : Title, Description, Priority
// create task : ID, Priority, Title, Description, Status, CreatedAt
// open/create file
// add tasks to file
// finish

func main() {
	// title, description, priority := operations()
	operations()
	// task := crud.TaskInput{
	// 	Title:       title,
	// 	Description: description,
	// 	Priority:    priority,
	// }
	// fmt.Println(tools.UpdateTask(4, "In Progress"))
	// tasks := tools.AddTask(task)
	// tools.CreateFile(tasks)
} 

func operations() {
	var operation string
	fmt.Println("Welcome to Task Tracker :)")
	fmt.Println("\n1. Create a Task\n2. Update a Task\n3. List of Tasks\n4. Delete a Task")
	fmt.Print("Please Choose a Number:")
	fmt.Scan(&operation)
	switch operation {
	case "1"  : 
		title, description, priority := CreateTaskMenu()
		task := crud.TaskInput{
		Title:       title,
		Description: description,
		Priority:    priority,
		}
		create(task)
		return 
	// case "2":
	// 	return
	default: 
		operations()
	}
} 

func CreateTaskMenu() (string, string,  int) {
	var title, description string
	var priority int
	fmt.Println()
	for len(title) == 0 {
		fmt.Print("Title: ")
		fmt.Scan(&title)
	}
	fmt.Print("Description: ")
	fmt.Scan(&description)
	for priority != 1 && priority != 2 && priority != 3 {
		fmt.Println("Priority (1 = High, 2 = Medium, 3 = Low): ")
		fmt.Scan(&priority)
	}
	return title, description, priority
}

func create(task crud.TaskInput) {
	tasks := tools.AddTask(task)
	tools.CreateFile(tasks)
	fmt.Print("Created!")
}