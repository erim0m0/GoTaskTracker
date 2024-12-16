package tools

import (
	"crud/crud"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func readFile() *os.File {
	file, err := os.Open("notes.json")
	if err != nil {
		return nil
	}
	return file
}

func getTasks() []crud.Task {
	file := readFile()
	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil
	}

	var tasks crud.Tasks
	json.Unmarshal(byteValue, &tasks)
	//  Handle Error Unmarshal
	defer file.Close()
	return tasks.Tasks
}

func AddTask(task crud.TaskInput) []crud.Task {
	newTask := crud.CreateTask(task)
	tasks := getTasks()
	var id int
	for _, tk := range tasks {
		if id < tk.ID {
			id = tk.ID
		}
	}
	newTask.ID = id + 1
	tasks = append(tasks, newTask)
	return tasks
}

func CreateFile(tasks []crud.Task) {
	file, _ := os.Create("notes.json")
	jsonTasks := crud.Tasks{Tasks: tasks}
	jsonBytes, _ := json.Marshal(jsonTasks)
	_, err := file.Write(jsonBytes)
	if err != nil {
		return
	}
	defer file.Close()
}

func DeleteTask(id int) string {
	tasks := getTasks()
	deleted := false
	var newTasks []crud.Task

	for _, t := range tasks {
		if t.ID == id {
			deleted = true
			continue
		}
		newTasks = append(newTasks, t)
	}

	if deleted {
		CreateFile(newTasks)
		return fmt.Sprintf("Deleted a task %d!", id)
	}
	return "Not found :|"
}

func UpdateTask(id int, status string) string {
	tasks := getTasks()
	updated := false
	var newTasks []crud.Task

	for _, t := range tasks {
		if t.ID == id {
			updated = true
			t.Status = status
		}
		newTasks = append(newTasks, t)
	}
	if updated {
		CreateFile(newTasks)
		return fmt.Sprintf("Updated a task %d!", id)
	}
	return "Not found :|"
}
