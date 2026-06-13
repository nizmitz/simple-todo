package todo

import (
	"testing"
)

func TestAddSuccess(t *testing.T) {
	Tasks = []Task{}

	Add([]string{"buy", "milk"})

	if len(Tasks) != 1 {
		t.Errorf("expected 1 task, got %d", len(Tasks))
	}
	if Tasks[0].Text != "buy milk" {
		t.Errorf("expected 'buy milk', got '%s'", Tasks[0].Text)
	}
	if Tasks[0].Done {
		t.Errorf("new task should not be done")
	}
}

func TestAddEmpty(t *testing.T) {
	Tasks = []Task{}

	Add([]string{})

	if len(Tasks) != 1 {
		t.Errorf("should add empty task")
	}
	if Tasks[0].Text != "" {
		t.Errorf("expected empty text")
	}
}

func TestDoneSuccess(t *testing.T) {
	Tasks = []Task{{Text: "task1", Done: false}}

	Done([]string{"0"})

	if !Tasks[0].Done {
		t.Errorf("expected task to be marked done")
	}
}

func TestDoneInvalidIndex(t *testing.T) {
	Tasks = []Task{{Text: "task1", Done: false}}

	// should not crash
	Done([]string{"abc"})

	if Tasks[0].Done {
		t.Errorf("should not modify on invalid index")
	}
}
