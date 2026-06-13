package main

import (
	"fmt"
	"os"
	"simple-todo/todo"
)

func main() {
	todo.LoadTasks()

	if len(os.Args) < 2 {
		fmt.Println("usage: todo add|list|done")
		return
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	switch cmd {
	case "add":
		todo.Add(args)
	case "list":
		todo.List()
	case "done":
		todo.Done(args)
	default:
		fmt.Println("unknown command")
	}
}
