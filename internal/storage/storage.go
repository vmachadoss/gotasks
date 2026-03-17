package storage

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func NewTask(description, priority string) *Task {
	return &Task{
		Description: description,
		Priority:    priority,
		Status:      "pending",
		CreatedAt:   "",
		UpdatedAt:   "",
	}
}

func LoadTasks() ([]Task, error) {
	return nil, nil
}

func SaveTasks(tasks []Task) error {
	return nil
}
