package cmd

import (
	"fmt"
	"internal/storage"
	"internal/task"
)

func AddTask(title string, description string) {
	tasks := storage.Load()
	newTask := task.Task{
		ID:          generateID(),
		Title:       title,
		Description: description,
		Status:      false,
	}

	tasks = append(tasks, newTask)
	storage.Save(tasks)

	fmt.Println("Task added successfully!")
}