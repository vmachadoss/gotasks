package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func setupTestStorage(t *testing.T) func() {
	tmpDir := t.TempDir()

	TestStoragePath = filepath.Join(tmpDir, "tasks.json")

	return func() {
		TestStoragePath = ""
	}
}

func TestNewTask_Fields(t *testing.T) {
	task := NewTask("Create a new task", "high")

	if task.Description != "Create a new task" {
		t.Errorf("expected description %q, got %q", "Create a new task", task.Description)
	}
	if task.Priority != "high" {
		t.Errorf("expected priority %q, got %q", "high", task.Priority)
	}
	if task.Status != "pending" {
		t.Errorf("expected status %q, got %q", "pending", task.Status)
	}
	if task.ID != 0 {
		t.Errorf("expected ID %d, got %d", 0, task.ID)
	}
	if task.CreatedAt != task.UpdatedAt {
		t.Errorf("expected CreatedAt == UpdatedAt on creation, got %q and %q", task.CreatedAt, task.UpdatedAt)
	}
	if task.CreatedAt == "" {
		t.Errorf("expected CreatedAt to be set, got empty string")
	}
}

func TestNextID_EmptySlice(t *testing.T) {
	id := NextID([]Task{})
	if id != 1 {
		t.Errorf("expected 1, got %d", id)
	}
}

func TestNextID_NonSequential(t *testing.T) {
	tasks := []Task{
		{ID: 1},
		{ID: 5},
		{ID: 3},
	}
	id := NextID(tasks)
	if id != 6 {
		t.Errorf("expected 6, got %d", id)
	}
}

func TestSaveAndLoadTasks(t *testing.T) {
	cleanup := setupTestStorage(t)
	defer cleanup()

	original := []Task{
		{ID: 1, Description: "Task A", Priority: "high", Status: "pending"},
		{ID: 2, Description: "Task B", Priority: "low", Status: "done"},
	}

	if err := SaveTasks(original); err != nil {
		t.Fatalf("SaveTasks failed: %v", err)
	}

	loaded, err := LoadTasks()
	if err != nil {
		t.Fatalf("LoadTasks failed: %v", err)
	}

	if len(loaded) != len(original) {
		t.Fatalf("expected %d tasks, got %d", len(original), len(loaded))
	}

	for i, task := range loaded {
		if task.ID != original[i].ID {
			t.Errorf("task[%d]: expected ID %d, got %d", i, original[i].ID, task.ID)
		}
		if task.Description != original[i].Description {
			t.Errorf("task[%d]: expected description %q, got %q", i, original[i].Description, task.Description)
		}
		if task.Priority != original[i].Priority {
			t.Errorf("task[%d]: expected priority %q, got %q", i, original[i].Priority, task.Priority)
		}
		if task.Status != original[i].Status {
			t.Errorf("task[%d]: expected status %q, got %q", i, original[i].Status, task.Status)
		}
	}
}

func TestLoadTasks_EmptyFile(t *testing.T) {
	cleanup := setupTestStorage(t)
	defer cleanup()

	tasks, err := LoadTasks()
	if err != nil {
		t.Fatalf("expected no error on empty storage, got: %v", err)
	}
	if len(tasks) != 0 {
		t.Errorf("expected 0 tasks, got %d", len(tasks))
	}
}

func TestGetTaskByID_Found(t *testing.T) {
	cleanup := setupTestStorage(t)
	defer cleanup()

	var tasks []Task
	for i := 0; i <= 10; i++ {
		tasks = append(tasks, Task{ID: i + 1, Description: fmt.Sprintf("Task %d", i+1)})
	}

	if err := SaveTasks(tasks); err != nil {
		t.Fatalf("SaveTasks failed: %v", err)
	}

	task, err := GetTaskByID(2)
	if err != nil {
		t.Fatalf("expected task, got error: %v", err)
	}
	if task.Description != "Task 2" {
		t.Errorf("expected %q, got %q", "Task 2", task.Description)
	}
}

func TestGetTaskByID_NotFound(t *testing.T) {
	cleanup := setupTestStorage(t)
	defer cleanup()

	if err := SaveTasks([]Task{{ID: 1}}); err != nil {
		t.Fatalf("SaveTasks failed: %v", err)
	}

	_, err := GetTaskByID(99)
	if err == nil {
		t.Error("expected error for non-existent ID, got nil")
	}
}

func TestUpdateTask_SingleField(t *testing.T) {
	cleanup := setupTestStorage(t)
	defer cleanup()

	original := []Task{
		{ID: 1, Description: "Original", Priority: "low", Status: "pending"},
	}
	if err := SaveTasks(original); err != nil {
		t.Fatalf("SaveTasks failed: %v", err)
	}

	updated, err := UpdateTask(1, map[string]string{"priority": "high"})
	if err != nil {
		t.Fatalf("UpdateTask failed: %v", err)
	}

	if updated.Priority != "high" {
		t.Errorf("expected priority %q, got %q", "high", updated.Priority)
	}
	if updated.Description != "Original" {
		t.Errorf("description should not have changed, got %q", updated.Description)
	}
	if updated.Status != "pending" {
		t.Errorf("status should not have changed, got %q", updated.Status)
	}
}

func TestUpdateTask_MultipleFields(t *testing.T) {
	cleanup := setupTestStorage(t)
	defer cleanup()

	if err := SaveTasks([]Task{
		{ID: 1, Description: "Old desc", Priority: "low", Status: "pending"},
	}); err != nil {
		t.Fatalf("SaveTasks failed: %v", err)
	}

	updated, err := UpdateTask(1, map[string]string{
		"description": "New desc",
		"status":      "done",
	})
	if err != nil {
		t.Fatalf("UpdateTask failed: %v", err)
	}

	if updated.Description != "New desc" {
		t.Errorf("expected %q, got %q", "New desc", updated.Description)
	}
	if updated.Status != "done" {
		t.Errorf("expected %q, got %q", "done", updated.Status)
	}
	if updated.Priority != "low" {
		t.Errorf("priority should not have changed, got %q", updated.Priority)
	}
}

func TestUpdateTask_NotFound(t *testing.T) {
	cleanup := setupTestStorage(t)
	defer cleanup()

	if err := SaveTasks([]Task{{ID: 1}}); err != nil {
		t.Fatalf("SaveTasks failed: %v", err)
	}

	_, err := UpdateTask(99, map[string]string{"status": "done"})
	if err == nil {
		t.Error("expected error for non-existent ID, got nil")
	}
}

func TestDeleteTask_Success(t *testing.T) {
	cleanup := setupTestStorage(t)
	defer cleanup()

	if err := SaveTasks([]Task{
		{ID: 1, Description: "Keep me"},
		{ID: 2, Description: "Delete me"},
	}); err != nil {
		t.Fatalf("SaveTasks failed: %v", err)
	}

	deleted, err := DeleteTask(2)
	if err != nil {
		t.Fatalf("DeleteTask failed: %v", err)
	}
	if deleted.ID != 2 {
		t.Errorf("expected deleted task ID 2, got %d", deleted.ID)
	}

	remaining, err := LoadTasks()
	if err != nil {
		t.Fatalf("LoadTasks after delete failed: %v", err)
	}
	if len(remaining) != 1 {
		t.Errorf("expected 1 task remaining, got %d", len(remaining))
	}
	if remaining[0].ID != 1 {
		t.Errorf("expected remaining task ID 1, got %d", remaining[0].ID)
	}
}

func TestDeleteTask_NotFound(t *testing.T) {
	cleanup := setupTestStorage(t)
	defer cleanup()

	if err := SaveTasks([]Task{{ID: 1}}); err != nil {
		t.Fatalf("SaveTasks failed: %v", err)
	}

	_, err := DeleteTask(99)
	if err == nil {
		t.Error("expected error for non-existent ID, got nil")
	}
}

func TestFilterTasks(t *testing.T) {
	cleanup := setupTestStorage(t)
	defer cleanup()

	if err := SaveTasks([]Task{
		{ID: 1, Status: "pending", Priority: "high"},
		{ID: 2, Status: "done", Priority: "low"},
		{ID: 3, Status: "pending", Priority: "low"},
	}); err != nil {
		t.Fatalf("SaveTasks failed: %v", err)
	}

	t.Run("by status", func(t *testing.T) {
		results, err := FilterTasks("pending", "", 0)
		if err != nil {
			t.Fatalf("FilterTasks failed: %v", err)
		}
		if len(results) != 2 {
			t.Errorf("expected 2 pending tasks, got %d", len(results))
		}
	})

	t.Run("by priority", func(t *testing.T) {
		results, err := FilterTasks("", "low", 0)
		if err != nil {
			t.Fatalf("FilterTasks failed: %v", err)
		}
		if len(results) != 2 {
			t.Errorf("expected 2 low priority tasks, got %d", len(results))
		}
	})

	t.Run("by id", func(t *testing.T) {
		results, err := FilterTasks("", "", 3)
		if err != nil {
			t.Fatalf("FilterTasks failed: %v", err)
		}
		if len(results) != 1 {
			t.Errorf("expected 1 task with ID 3, got %d", len(results))
		}
	})

	t.Run("combined status and priority", func(t *testing.T) {
		results, err := FilterTasks("pending", "low", 0)
		if err != nil {
			t.Fatalf("FilterTasks failed: %v", err)
		}
		if len(results) != 1 {
			t.Errorf("expected 1 task, got %d", len(results))
		}
		if results[0].ID != 3 {
			t.Errorf("expected task ID 3, got %d", results[0].ID)
		}
	})

	t.Run("no match", func(t *testing.T) {
		results, err := FilterTasks("in-progress", "", 0)
		if err != nil {
			t.Fatalf("FilterTasks failed: %v", err)
		}
		if len(results) != 0 {
			t.Errorf("expected 0 tasks, got %d", len(results))
		}
	})
}

func TestEnsureStorageExists_CreatesFile(t *testing.T) {
	cleanup := setupTestStorage(t)
	defer cleanup()

	if _, err := os.Stat(TestStoragePath); !os.IsNotExist(err) {
		t.Fatal("expected file to not exist before LoadTasks")
	}

	if _, err := LoadTasks(); err != nil {
		t.Fatalf("LoadTasks should create file automatically, got error: %v", err)
	}

	if _, err := os.Stat(TestStoragePath); err != nil {
		t.Errorf("expected file to exist after LoadTasks, got: %v", err)
	}
}
