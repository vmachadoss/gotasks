package cmd

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/vmachadoss/gotasks/internal/storage"
)

// 	"github.com/vmachadoss/gotasks/internal/storage"

func RunAdd(args []string) {
	addCmd := flag.NewFlagSet("add", flag.ContinueOnError)
	priority := addCmd.String("priority", "medium", "Task priority (low, medium, high)")

	if err := addCmd.Parse(args); err != nil {
		fmt.Println("Invalid argument")
		return
	}

	descArgs := addCmd.Args()

	if len(descArgs) == 0 {
		fmt.Println("Uso: gotasks add <descrição> [--priority low|medium|high]")
		os.Exit(1)
	}

	description := strings.Join(descArgs, " ")

	validPriorities := map[string]bool{
		"low":    true,
		"medium": true,
		"high":   true,
	}

	if !validPriorities[*priority] {
		fmt.Println("Invalid priority. Use low, medium, or high instead.")
		os.Exit(1)
	}

	task := storage.NewTask(description, *priority)
	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		os.Exit(1)
	}

	tasks = append(tasks, task)
	if err := storage.SaveTasks(tasks); err != nil {
		fmt.Printf("Error saving task: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Task criada com sucesso!\n")
	fmt.Printf("  Desc:     %s\n", task.Description)
	fmt.Printf("  Status:   %s\n", task.Status)
	fmt.Printf("  Priority: %s\n", task.Priority)
}
