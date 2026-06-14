package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"simple-todo/todo"
)

func handleListTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo.Tasks)
	w.WriteHeader(http.StatusOK)
}

func handleAddTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req struct {
		Text string `json:"text"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	todo.Add([]string{req.Text})
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo.Tasks[len(todo.Tasks)-1])
	w.WriteHeader(http.StatusOK)
}

func handleDoneTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	parts := strings.Split(r.URL.Path, "/")
	idx := parts[len(parts)-1]

	todo.Done([]string{idx})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func main() {
	todo.LoadTasks()

	http.HandleFunc("/tasks", handleListTasks)
	http.HandleFunc("/tasks/add", handleAddTask)
	http.HandleFunc("/tasks/done/", handleDoneTask)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
