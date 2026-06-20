package cmd

import (
	"fmt"
	"internal/storage"
	"internal/task"
)

func ListTasks() {
	tasks := storage.Load()
	if len(tasks) == 0 {
		fmt.Println("No tasks found!")
		return
	}

	fmt.Println("Tasks:")
	for _, t := range tasks {
		status := "Incomplete"
		if t.Status {
			status = "Complete"
		}
		fmt.Printf("ID: %s\nTitle: %s\nDescription: %s\nStatus: %s\n\n", t.ID, t.Title, t.Description, status)
	}
}