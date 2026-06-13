# simple-todo

A CLI todo app written in Go. Manage tasks from the command line with persistent JSON storage.

## Features

- Add tasks
- List all tasks
- Mark tasks as done
- Persist tasks to JSON file

## Requirements

- Go 1.16 or later
- Nix (optional, for development environment)

## Installation

```bash
git clone <repo-url>
cd simple-todo
go build -o simple-todo
```

## Usage

```bash
# Add a task
./simple-todo add buy milk
./simple-todo add wash car

# List all tasks
./simple-todo list

# Mark task as done (by index)
./simple-todo done 0

# List again to see updates
./simple-todo list
```

## Development

Set up development environment with Nix:

```bash
nix develop
```

Run tests:

```bash
go test ./...
```

Format and lint:

```bash
go fmt ./...
go vet ./...
pre-commit run --all-files
```

## Project Structure

```
simple-todo/
├── main.go          # CLI entry point
├── go.mod           # Module definition
├── go.sum           # Dependencies
├── todo/            # Todo package
│   ├── task.go      # Task struct and persistence
│   ├── commands.go  # Command logic (add, list, done)
│   ├── commands_test.go
│   └── task_test.go
├── flake.nix        # Nix dev environment
└── README.md        # This file
```

## License

MIT License - See LICENSE file for details
