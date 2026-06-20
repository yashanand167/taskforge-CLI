package cmd

import (
	"fmt"
	"os"
)

func Execute() {
	if len(os.Args) < 2 {
		fmt.Println("No command provided.")
		return
	}

	command := os.Args[1]
	fmt.Printf("Executing command: %s\n", command)

	switch command {
	case "add":
		if len(os.Args) < 4 {
			fmt.Println("Usage: add <title> <description>")
			return
		}
		title := os.Args[2]
		description := os.Args[3]
		AddTask(title, description)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: delete <id>")
			return
		}
		id := os.Args[2]
		DeleteTask(id)
		
	default:
		fmt.Printf("Unknown command: %s\n", command)
	}
}