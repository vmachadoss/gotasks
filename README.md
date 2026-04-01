# 🎯 GoTasks

> A blazing-fast, minimal task management CLI built with Go for developers who want to stay organized without the bloat.

[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-blue?style=for-the-badge)](LICENSE)
[![Status](https://img.shields.io/badge/Status-Active-brightgreen?style=for-the-badge)](https://github.com/vmachadoss/gotasks)

## ✨ Features

- 🚀 **Lightning Fast** - Built with Go for maximum performance
- 📋 **Task Management** - Create, edit, and manage your daily tasks
- 🎯 **Priority Levels** - Mark tasks as low, medium, or high priority
- ✅ **Status Tracking** - Keep track of pending, in-progress, and completed tasks
- 💾 **Persistent Storage** - Your tasks are automatically saved locally
- 🎨 **Clean CLI** - Minimal, intuitive command interface
- 🔧 **Zero Configuration** - Works out of the box

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
./gotasks
```

## 📖 Usage

### Add a Task

Create a new task with an optional priority level:

```bash
# Basic task
gotasks add "Complete project report"

# With priority
gotasks add "Fix critical bug" --priority high
gotasks add "Review PR" --priority medium
gotasks add "Clean up code" --priority low
```

**Options:**
- `--priority` - Task priority: `low`, `medium`, `high` (default: `medium`)

### Edit a Task

Update an existing task by its ID:

```bash
# Update description
gotasks edit 1 --desc "New task description"

# Change priority
gotasks edit 1 --priority high

# Update status
gotasks edit 1 --status in-progress
gotasks edit 1 --status done

# Multiple updates at once
gotasks edit 1 --desc "Updated task" --priority high --status in-progress
```

**Options:**
- `--desc` - New task description
- `--priority` - Task priority: `low`, `medium`, `high`
- `--status` - Task status: `pending`, `in-progress`, `done`

## 🏗️ Project Structure

```
gotasks/
├── main.go                 # CLI entry point
├── cmd/
│   ├── add.go             # Add task handler
│   └── edit.go            # Edit task handler
├── internal/
│   └── storage/
│       └── storage.go     # Task persistence layer
└── README.md
```

## 💡 Workflow Example

```bash
# Start your day
gotasks add "Morning standup" --priority high
gotasks add "Code review" --priority medium
gotasks add "Update documentation" --priority low

# Update as you progress
gotasks edit 1 --status in-progress
gotasks edit 1 --status done
gotasks edit 2 --status in-progress

# Keep your list clean
gotasks edit 3 --status done
```

## 🔮 Upcoming Features

- 📱 Task listing with filters
- 🗑️ Task deletion
- 🔍 Search functionality

## 🤝 Contributing

Contributions are welcome! Feel free to:

1. Clone the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 👨‍💻 Author

**Vitor Machado** - [@vmachadoss](https://github.com/vmachadoss)

---

<div align="center">

**Made with ❤️ for developers who value simplicity and speed**

⭐ If you find this useful, please consider giving it a star!

</div>
