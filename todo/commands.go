package todo

import (
	"fmt"
	"strconv"
	"strings"
)

func Add(args []string) {
	text := strings.Join(args, " ")
	Tasks = append(Tasks, Task{Text: text, Done: false})
	fmt.Println("added:", text)
	SaveTasks()
}

func List() {
	for i, t := range Tasks {
		mark := " "
		if t.Done {
			mark = "x"
		}
		fmt.Printf("%d [%s] %s\n", i, mark, t.Text)
	}
}

func Done(args []string) {
	idx, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("invalid index")
		return
	}
	Tasks[idx].Done = true
	fmt.Println("marked done:", Tasks[idx].Text)
	SaveTasks()
}
