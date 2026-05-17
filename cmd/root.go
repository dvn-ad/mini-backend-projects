package cmd

import (
	"fmt"
	"os"
)

func Run() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: task-cli <command> [arguments]")
		return
	}

	command := args[1]
	cmdArgs := args[2:]

	switch command {
	case "add":
		if len(cmdArgs) < 1 {
			fmt.Println("Error: missing task description")
			return
		}
		fmt.Printf("'%s' added\n", cmdArgs[0])

	case "list":
		fmt.Println("Tasks:")

	default:
		fmt.Println("Unknown command:", command)
	}
}