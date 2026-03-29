package cmd

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/vmachadoss/gotasks/internal/storage"
)

func RunEdit(args []string) {
	editCmd := flag.NewFlagSet("edit", flag.ContinueOnError)

	desc := editCmd.String("desc", "", "Task description")
	priority := editCmd.String("priority", "", "Task priority (low, medium, high)")
	status := editCmd.String("status", "", "Task status (pending, in-progress, done)")

	if len(args) == 0 {
		fmt.Println("Use: gotasks edit <id> [--desc \"...\"] [--priority low|medium|high] [--status pending|in-progress|done]")
		os.Exit(1)
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Invalid ID: %q — must be an integer\n", args[0])
		os.Exit(1)
	}

	if err := editCmd.Parse(args[1:]); err != nil {
		fmt.Println("Invalid argument")
		return
	}

	if *priority != "" {
		validPriorities := map[string]bool{"low": true, "medium": true, "high": true}
		if !validPriorities[*priority] {
			fmt.Println("Prioridade inválida. Use: low, medium ou high")
			os.Exit(1)
		}
	}

	if *status != "" {
		validStatuses := map[string]bool{"pending": true, "done": true, "in-progress": true}
		if !validStatuses[*status] {
			fmt.Println("Invalid status. Use: pending, in-progress ou done")
			os.Exit(1)
		}
	}

	updates := map[string]string{
		"description": *desc,
		"priority":    *priority,
		"status":      *status,
	}

	if updates["description"] == "" && updates["priority"] == "" && updates["status"] == "" {
		fmt.Println("At least one field needs to be filled in: --desc, --priority or --status")
		os.Exit(1)
	}

	updated, err := storage.UpdateTask(id, updates)
	if err != nil {
		fmt.Printf("Failed to update task: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Update successful!\n")
	fmt.Printf("  ID:       %d\n", updated.ID)
	fmt.Printf("  Desc:     %s\n", updated.Description)
	fmt.Printf("  Status:   %s\n", updated.Status)
	fmt.Printf("  Priority: %s\n", updated.Priority)
}
