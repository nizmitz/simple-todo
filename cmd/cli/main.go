package main

import (
	"fmt"
	"os"
	"simple-todo/todo"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: todo [serve|add|list|done]")
		return
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	switch cmd {
	case "add":
		todo.LoadTasks()
		todo.Add(args)
	case "list":
		todo.LoadTasks()
		todo.List()
	case "done":
		todo.LoadTasks()
		todo.Done(args)
	default:
		fmt.Println("unknown command")
	}
}
