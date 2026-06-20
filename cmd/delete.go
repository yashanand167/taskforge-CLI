package cmd

import (
	"fmt"
	"internal/storage"
	"internal/task"
)

func DeleteTask(id string) {
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

	tasks = append(tasks[:index], tasks[index+1:]...)
	storage.Save(tasks)

	fmt.Println("Task deleted successfully!")
}