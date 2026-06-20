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
		
	}
}