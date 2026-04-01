package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/vmachadoss/gotasks/internal/storage"
)

func RunGet(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: get <task_id>")
		os.Exit(1)
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Invalid task ID: %q - must be a number\n", args[0])
		os.Exit(1)
	}

	task, err := storage.GetTaskByID(id)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Task details:")
	fmt.Println("  --------------------------------")
	fmt.Printf("  ID:         %d\n", task.ID)
	fmt.Printf("  Description: %s\n", task.Description)
	fmt.Printf("  Status:      %s\n", task.Status)
	fmt.Printf("  Priority:    %s\n", task.Priority)
	fmt.Printf("  Created at:  %s\n", task.CreatedAt)
	fmt.Printf("  Updated at:  %s\n", task.UpdatedAt)
	fmt.Println("  --------------------------------")
}
