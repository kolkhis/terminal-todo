package tasks

import (
	"fmt"
	"log"
	"os"
)

type Task struct {
	Title       string
	Description string
	Completed   bool
	Id          int
}

type TaskList struct {
	Tasks      []Task
	nextTaskId int
}

func NewTaskList() TaskList {
	tl := TaskList{
		nextTaskId: 0,
	}
	return tl
}

// Constructor for Task
func NewTask(title, desc string, id int) Task {
	t := Task{
		Title:       title,
		Description: desc,
		Id:          id,
		Completed:   false,
	}
	return t
}

func (tl *TaskList) AddTaskToList(t Task) {
	t.Id = tl.nextTaskId
	tl.Tasks = append(tl.Tasks, t)
	tl.nextTaskId++
}

func (tl *TaskList) DeleteTask(t Task) {}

// TODO: Add method to show only incomplete tasks, and only complete tasks
func (tl *TaskList) ViewTaskList() {
	for idx := range tl.Tasks {
		t := tl.Tasks[idx]
		// fmt.Printf("\nTitle: %v\n", t.Title)
		// fmt.Printf("Description: %v\n", t.Description)
		// fmt.Printf("Completed: %v\n", t.Completed)
		// fmt.Printf("ID: %v\n", t.Id)
		fmt.Printf(`
    Title: %v
    Description: %v
    Completed: %v
    ID: %v
            `, t.Title, t.Description, t.Completed, t.Id)
	}
}

// TODO: Add UI for these.
// Mark a task as completed
func (t *Task) MarkComplete() { t.Completed = true }

func (t *Task) ChangeDescription(newDesc string) { t.Description = newDesc }
func (t *Task) ChangeTitle(newTitle string)      { t.Title = newTitle }

func (*TaskList) LoadTaskList() {
	path := "/etc/terminal-todo/tasklist.json"
	fmt.Printf("Attempting to open file at; %v\n", path)
	file, err := os.Open(path)
	if err != nil {
		log.Printf("Couldn't open file /etc/terminal-todo/tasklist.json")
	}
	defer file.Close()
}

func (*TaskList) SaveTaskList() {}

func MakeGreeting(first string, last string) string {
	greeting := fmt.Sprintf("Hello, %v %v", first, last)
	return greeting
}

func (tl *TaskList) ShowType() {
	for i := range tl.Tasks {
		fmt.Println("Type of task:", tl.Tasks[i])
	}
}

// git commit -m "go: Update notes on interfaces, move function notes to different file"
