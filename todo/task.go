package todo

import (
	"encoding/json"
	"os"
)

const filename = "tasks.json"

type Task struct {
	Text string
	Done bool
}

var Tasks []Task

func LoadTasks() {
	data, err := os.ReadFile(filename)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &Tasks)
	if err != nil {
		return
	}
}

func SaveTasks() {
	data, _ := json.MarshalIndent(Tasks, "", " ")
	err := os.WriteFile(filename, data, 0644)
	if err != nil {
		return
	}
}
