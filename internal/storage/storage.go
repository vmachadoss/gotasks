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

// NewTask creates a new task struct with the given description and priority.
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

// LoadTasks loads the tasks from the tasks file.
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

// SaveTasks saves the tasks to the tasks file.
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

// NextID returns the next available task ID.
func NextID(tasks []Task) int {
	max := 0
	for _, task := range tasks {
		if task.ID > max {
			max = task.ID
		}
	}
	return max + 1
}

// UpdateTask updates the task with the given ID using the provided updates.
func UpdateTask(id int, updates map[string]string) (*Task, error) {
	tasks, err := LoadTasks()
	if err != nil {
		return nil, err
	}

	var found *Task
	for i, task := range tasks {
		if task.ID == id {
			for key, value := range updates {
				switch key {
				case "description":
					if value != "" {
						tasks[i].Description = value
					}
				case "priority":
					if value != "" {
						tasks[i].Priority = value
					}
				case "status":
					if value != "" {
						tasks[i].Status = value
					}
				}
			}
			tasks[i].UpdatedAt = time.Now().Format("2006-01-02T15:04:05Z07:00")
			found = &tasks[i]
			break
		}
	}

	if found == nil {
		return nil, fmt.Errorf("Task %d not found", id)
	}

	if err := SaveTasks(tasks); err != nil {
		return nil, err
	}

	return found, nil
}

// GetTaskByID returns the task with the given ID.
func GetTaskByID(id int) (*Task, error) {
	tasks, err := LoadTasks()
	if err != nil {
		return nil, err
	}

	for _, task := range tasks {
		if task.ID == id {
			t := task
			return &t, nil
		}
	}
	return nil, fmt.Errorf("task with ID %d not found", id)
}

// FilterTasks filters the tasks based on the given status, priority, and ID.
func FilterTasks(status, priority string, id int) ([]Task, error) {
	tasks, err := LoadTasks()
	if err != nil {
		return nil, err
	}

	var filtered []Task

	for _, task := range tasks {
		if status != "" && task.Status != status {
			continue
		}
		if priority != "" && task.Priority != priority {
			continue
		}
		if id != 0 && task.ID != id {
			continue
		}
		filtered = append(filtered, task)
	}

	return filtered, nil
}

func DeleteTask(id int) (*Task, error) {
	tasks, err := LoadTasks()
	if err != nil {
		return nil, err
	}

	var deleted *Task
	for i, task := range tasks {
		if task.ID == id {
			t := task
			deleted = &t

			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}

	if deleted == nil {
		return nil, fmt.Errorf("Task %d not found", id)
	}

	if err := SaveTasks(tasks); err != nil {
		return nil, err
	}

	return deleted, nil
}
