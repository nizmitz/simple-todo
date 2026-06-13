package todo

import (
	"testing"
)

func TestLoadTasksFileNotFound(t *testing.T) {
	// file doesn't exist yet, should not crash
	LoadTasks()
	// just verify no panic
}

func TestTaskStruct(t *testing.T) {
	task := Task{Text: "test", Done: false}

	if task.Text != "test" {
		t.Errorf("expected 'test', got '%s'", task.Text)
	}
	if task.Done {
		t.Errorf("should not be done")
	}
}
