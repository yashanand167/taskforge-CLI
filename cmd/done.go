package cmd

import (
	"fmt"
	"internal/storage"
	"internal/task"
)

func DoneTask(id string) {
	tasks := storage.Load()
	index := -1

	for i, t := range tasks {
		if t.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Task not found!")
		return
	}

	tasks[index].Status = true
	storage.Save(tasks)

	fmt.Println("Task marked as done successfully!")
}