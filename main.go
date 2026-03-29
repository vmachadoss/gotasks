package main

import (
	"fmt"
	"os"

	"github.com/vmachadoss/gotasks/cmd"
)

func main() {
	fmt.Println("Welcome to GoTasks!")
	subcommand := os.Args[1]
	args := os.Args[2:]

	switch subcommand {
	case "add":
		cmd.RunAdd(args)
	// case "list":
	// 	cmd.ListTasks()
	// case "delete":
	// 	cmd.DeleteTask(args)
	default:
		fmt.Println("Unknown subcommand")
	}
}
