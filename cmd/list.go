package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/vmachadoss/gotasks/internal/storage"
)

/*
RunList lists all tasks when the "gotasks list" command is invoked.

Usage: gotasks list [--status <status>] [--priority <priority>] [--id <id>]
*/
func RunList(args []string) {
	listCmd := flag.NewFlagSet("list", flag.ContinueOnError)

	status := listCmd.String("status", "", "Filter by status (pending, in-progress, done)")
	priority := listCmd.String("priority", "", "Filter by priority (low, medium, high)")
	id := listCmd.Int("id", 0, "Filter by ID")

	if err := listCmd.Parse(args); err != nil {
		fmt.Println("Invalid arguments")
		return
	}

	if *status != "" {
		validStatus := map[string]bool{
			"pending":     true,
			"in-progress": true,
			"done":        true,
		}
		if !validStatus[*status] {
			fmt.Println("Invalid status. Use: pending, in-progress, done.")
			return
		}
	}
	if *priority != "" {
		validPriority := map[string]bool{
			"low":    true,
			"medium": true,
			"high":   true,
		}
		if !validPriority[*priority] {
			fmt.Println("Invalid priority. Use: low, medium, high.")
			return
		}
	}

	tasks, err := storage.FilterTasks(*status, *priority, *id)
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		os.Exit(1)
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	fmt.Printf("%-4s  %-50s  %-11s  %-10s\n", "ID", "DESCRIPTION", "STATUS", "PRIORITY")
	fmt.Println("----  --------------------------------------------------  -----------  ----------")

	for _, task := range tasks {
		fmt.Printf("%-4d  %-50.50s  %-11s  %-10s\n",
			task.ID,
			task.Description,
			task.Status,
			task.Priority,
		)
	}

	fmt.Printf("\n%d task(s) found.\n", len(tasks))
}
