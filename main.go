package main

import (
	"fmt"
)

// get inputs from user : Title, Description, Priority
// create task : ID, Priority, Title, Description, Status, CreatedAt
// open/create file
// add tasks to file
// finish

func main() {
	var name string
	for name == "" {
		fmt.Scan(&name)
	}

	fmt.Println(name)
	// task := crud.TaskInput{
	// 	Title:       "task1",
	// 	Description: "task1",
	// 	Priority:    3,
	// 	// Status:      "updated",
	// 	// CreatedAt:   time.Now(),
	// }
	// fmt.Println(tools.UpdateTask(4, "In Progress"))
	// tasks := tools.AddTask(task)
	// tools.CreateFile(tasks)
}
