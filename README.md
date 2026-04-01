# 🚀 GoTasks

> **A lightning-fast, minimal task management CLI for developers who live in the terminal**

[![Go Version](https://img.shields.io/badge/Go-1.19+-00ADD8?style=flat-square&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)](LICENSE)
[![Status](https://img.shields.io/badge/Status-Active-brightgreen?style=flat-square)](https://github.com/vmachadoss/gotasks)

---

## ✨ Features

- 🎯 **Add Tasks** - Create tasks with priority levels (low, medium, high)
- 📋 **List Tasks** - View all tasks with powerful filtering options
- 👁️ **View Details** - Get detailed information about a specific task
- ✏️ **Edit Tasks** - Update descriptions, priorities, and status in-flight
- 🔍 **Smart Filtering** - Filter by status, priority, or task ID
- 💾 **Persistent Storage** - Your tasks are saved locally
- ⚡ **Lightning Fast** - Built in Go for minimal resource usage
- 🎨 **Clean Interface** - Intuitive, minimal command structure

---

## 🚀 Quick Start

### Installation

```bash
go install github.com/vmachadoss/gotasks@latest
```

Or clone and build from source:

```bash
git clone https://github.com/vmachadoss/gotasks.git
cd gotasks
go build -o gotasks
./gotasks add "My first task"
```

---

## 📖 Complete Usage Guide

### 1️⃣ Add a Task

Create a new task with an optional priority level:

```bash
# Basic task (default priority: medium)
gotasks add "Complete project report"

# With specific priority
gotasks add "Fix critical bug" --priority high
gotasks add "Review pull request" --priority medium
gotasks add "Update dependencies" --priority low
```

**Options:**
- `--priority` - Task priority: `low`, `medium`, `high` (default: `medium`)

**Example Output:**
```
Task criada com sucesso!
  Desc:     Complete project report
  Status:   pending
  Priority: medium
```

---

### 2️⃣ List All Tasks

View all your tasks in a clean table format with optional filtering:

```bash
# List all tasks
gotasks list

# Filter by status
gotasks list --status pending
gotasks list --status in-progress
gotasks list --status done

# Filter by priority
gotasks list --priority high
gotasks list --priority medium
gotasks list --priority low

# Filter by ID
gotasks list --id 1

# Combine filters
gotasks list --status in-progress --priority high
```

**Example Output:**
```
ID    DESCRIPTION                                        STATUS        PRIORITY
----  --------------------------------------------------  -----------  ----------
1     Complete project report                           pending      medium
2     Fix critical bug                                  in-progress  high
3     Review pull request                               done         medium

3 task(s) found.
```

**Filter Options:**
- `--status` - Filter by status: `pending`, `in-progress`, `done`
- `--priority` - Filter by priority: `low`, `medium`, `high`
- `--id` - Filter by specific task ID

---

### 3️⃣ View Task Details

Get comprehensive information about a specific task, including timestamps:

```bash
# View task details
gotasks get 1
```

**Example Output:**
```
Task details:
  --------------------------------
  ID:          1
  Description: Complete project report
  Status:      pending
  Priority:    medium
  Created at:  2024-01-15T10:30:45Z
  Updated at:  2024-01-15T10:30:45Z
  --------------------------------
```

---

### 4️⃣ Edit a Task

Update an existing task by its ID:

```bash
# Update description
gotasks edit 1 --desc "Complete final project report"

# Change priority
gotasks edit 1 --priority high

# Update status
gotasks edit 1 --status in-progress
gotasks edit 1 --status done

# Update multiple fields at once
gotasks edit 1 --desc "Updated task" --priority high --status in-progress
```

**Options:**
- `--desc` - New task description
- `--priority` - Task priority: `low`, `medium`, `high`
- `--status` - Task status: `pending`, `in-progress`, `done`

**Example Output:**
```
Update successful!
  ID:       1
  Desc:     Complete final project report
  Status:   in-progress
  Priority: high
```

---

## 📊 Task Management Reference

### Priority Levels

| Level | Use Case |
|-------|----------|
| 🔴 **high** | Urgent/blocking tasks, critical bugs, tight deadlines |
| 🟡 **medium** | Regular tasks, standard work (default) |
| 🟢 **low** | Nice-to-have, backlog items, future improvements |

### Task Statuses

| Status | Meaning |
|--------|---------|
| **pending** | Not started yet |
| **in-progress** | Currently working on it |
| **done** | Completed ✓ |

---

## 💡 Real-World Workflow Example

```bash
# Start your day - create tasks
gotasks add "Morning standup" --priority high
gotasks add "Code review for PR #123" --priority medium
gotasks add "Update dependencies" --priority low
gotasks add "Fix bug #456" --priority high

# View all high priority tasks
gotasks list --priority high

# Start working - mark as in-progress
gotasks edit 1 --status in-progress

# View detailed info on a task
gotasks get 1

# Mark tasks as done
gotasks edit 1 --status done
gotasks edit 4 --status done

# See what's still pending
gotasks list --status pending

# Review all completed tasks today
gotasks list --status done
```

---

## 🏗️ Project Structure

```
gotasks/
├── main.go                 # CLI entry point & command router
├── cmd/                    # Command implementations
│   ├── add.go             # Add task handler
│   ├── list.go            # List tasks with filtering
│   ├── get.go             # Get task details
│   └── edit.go            # Edit task handler
├── internal/
│   └── storage/
│       └── storage.go     # Task persistence & filtering logic
└── README.md
```

---

## 🔮 Upcoming Features

- 🗑️ **Delete Tasks** - Remove unwanted tasks
- 🔍 **Advanced Search** - Search by keywords in descriptions
- 📊 **Statistics** - Get productivity insights and analytics
- 🏷️ **Tags** - Organize tasks with custom tags
- ⏰ **Due Dates** - Set deadlines for tasks
- 🌐 **Cloud Sync** - Optional cloud backup
- 🎨 **Custom Themes** - Personalize your output

---

## 🛠️ Development

### Build from Source

```bash
go build -o gotasks
```

### Run Tests

```bash
go test ./...
```

### Code Quality

```bash
# Format code
go fmt ./...

# Check for issues
go vet ./...
```

---

## 🤝 Contributing

Contributions are welcome! Here's how to get started:

1. **Fork** the repository
2. **Create** a feature branch (`git checkout -b feature/amazing-feature`)
3. **Make** your changes and test them
4. **Commit** with clear messages (`git commit -m 'Add amazing feature'`)
5. **Push** to your branch (`git push origin feature/amazing-feature`)
6. **Open** a Pull Request

### Found a Bug?

Please open an [issue](https://github.com/vmachadoss/gotasks/issues) with:
- A clear description
- Steps to reproduce
- Expected vs actual behavior
- Your environment (OS, Go version)

---

## 📝 License

MIT License © 2026 - See the [LICENSE](LICENSE) file for details.

This project is free to use, modify, and distribute.

---

## 👨‍💻 Author

**Vitor Machado** - [@vmachadoss](https://github.com/vmachadoss)

---

<div align="center">

### 🌟 Made with ❤️ for developers who value speed and simplicity

If you find GoTasks useful, please consider giving it a **star** ⭐

[Report Bug](https://github.com/vmachadoss/gotasks/issues) · [Request Feature](https://github.com/vmachadoss/gotasks/issues)

</div>
