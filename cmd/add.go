package cmd

import (
	"flag"
	"fmt"
	"os"
	"strings"
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
}
