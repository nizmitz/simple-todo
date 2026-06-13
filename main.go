package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const filename = "tasks.json"

type Task struct {
	Text string
	Done bool
}

var tasks []Task

func loadTasks() {
	data, err := os.ReadFile(filename)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return
	}
}

func saveTasks() {
	data, _ := json.MarshalIndent(tasks, "", " ")
	err := os.WriteFile(filename, data, 0644)
	if err != nil {
		return
	}
}

func main() {
	loadTasks()
	if len(os.Args) < 2 {
		fmt.Println("Usage: todo <command>")
		os.Exit(1)
	}
	command := os.Args[1]

	switch command {
	case "add":
		text := strings.Join(os.Args[2:], " ")
		tasks = append(tasks, Task{Text: text, Done: false})
		fmt.Println("added:", text)
		saveTasks()

	case "list":
		for i, t := range tasks {
			mark := " "
			if t.Done {
				mark = "x"
			}
			fmt.Printf("%d [%s] %s\n", i, mark, t.Text)
		}

	case "done":
		idx, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("invalid index")
			return
		}
		tasks[idx].Done = true
		fmt.Println("marked done:", tasks[idx].Text)
		saveTasks()

	default:
		fmt.Println("unknown command")
	}
}
