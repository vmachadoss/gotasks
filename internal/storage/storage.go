package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func tasksFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".gotasks", "tasks.json"), nil
}

func ensureStorageExists(filePath string) error {
	dir := filepath.Dir(filePath)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create the directory %s: %w", dir, err)
	}

	_, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		if err := os.WriteFile(filePath, []byte("[]"), 0644); err != nil {
			return fmt.Errorf("failed to create the file %s: %w", filePath, err)
		}
		return nil
	}

	if err != nil {
		return fmt.Errorf("erro ao verificar arquivo de tasks: %w", err)
	}

	return nil
}

func NewTask(description, priority string) *Task {
	now := time.Now().Format("2006-01-02T15:04:05Z07:00")

	return &Task{
		Description: description,
		Priority:    priority,
		Status:      "pending",
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func LoadTasks() ([]Task, error) {
	filePath, err := tasksFilePath()
	if err != nil {
		return nil, err
	}

	if err := ensureStorageExists(filePath); err != nil {
		return nil, err
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read the tasks file: %w", err)
	}

	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("failed to unmarshal the tasks file: %w", err)
	}

	return tasks, nil
}

func SaveTasks(tasks []Task) error {
	filePath, err := tasksFilePath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal the tasks: %w", err)
	}

	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write the tasks file: %w", err)
	}

	return nil
}

func NextID(tasks []Task) int {
	max := 0
	for _, task := range tasks {
		if task.ID > max {
			max = task.ID
		}
	}
	return max + 1
}
