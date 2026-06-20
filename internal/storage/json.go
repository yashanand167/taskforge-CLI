package storage

import (
	"encoding/json"
	"os"
	"internal/task"
)

func Load() []task.Task {
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		return []task.Task{}
	}
	
	var tasks []task.Task
	err = json.Unmarshal(data, &tasks)

	if(err != nil) {
		return []task.Task{}
	}

	return tasks
}

func Save(tasks []task.Task) {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return
	}

	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		return
	}
}