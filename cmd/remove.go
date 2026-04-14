package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/vmachadoss/gotasks/internal/storage"
)

/*
RunRm deletes a task by its ID.

Usage: gotasks rm <id>
*/
func RunRm(args []string) {
	if len(args) == 0 {
		fmt.Println("Use: gotasks rm <id>")
		os.Exit(1)
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Invlaid ID: %q - must be an number\n", args[0])
		os.Exit(1)
	}

	task, err := storage.GetTaskByID(id)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Task to be deleted:")
	fmt.Printf("  ID:          %d\n", task.ID)
	fmt.Printf("  Description: %s\n", task.Description)
	fmt.Printf("  Status:      %s\n", task.Status)
	fmt.Printf("  Priority:    %s\n", task.Priority)
	fmt.Println()

	fmt.Printf("Are you sure you want to delete this task? (y/n): ")
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		os.Exit(1)
	}

	confirmed := strings.ToLower(strings.TrimSpace(input))

	if confirmed != "y" {
		fmt.Println("Deletion cancelled.")
		return
	}

	if _, err := storage.DeleteTask(id); err != nil {
		fmt.Printf("Error deleting task: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Task %d deleted successfully.\n", id)
}
