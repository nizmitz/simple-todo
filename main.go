package main

import (
	"fmt"
	"github.com/nizmitz/simple-todo/todo"
	"os"
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
