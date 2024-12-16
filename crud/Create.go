package crud

import (
	"time"
)

type Tasks struct{
	Tasks []Task `json:"tasks"`
}

type Task struct{
	ID 				  int `json:"id"`
	Priority 		int `json:"priority"`
	Title 			   string  `json:"title"`
	Description  string  `json:"description"`
	Status 			string  `json:"status"`
	CreatedAt 	 time.Time  `json:"created_at"`
}

type TaskInput struct{
	Title string
	Description string
	Priority int
}

func CreateTask(payload TaskInput) Task {
	return Task{
		Title: payload.Title,
		Description: payload.Description,
		Priority: payload.Priority,
		Status: "Pending",
		CreatedAt: time.Now(),
	}
}
