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

// ViewTaskList Outputs the current task list to the terminal
// TODO: Add method (or option) to show only incomplete tasks, and only complete tasks
func (tl *TaskList) ViewTaskList() {
	fmt.Println("--- All tasks ---")
	for idx := range tl.Tasks {
		t := tl.Tasks[idx]
		// fmt.Printf("\nTitle: %v\n", t.Title)
		// fmt.Printf("Description: %v\n", t.Description)
		// fmt.Printf("Completed: %v\n", t.Completed)
		// fmt.Printf("ID: %v\n", t.Id)
		fmt.Printf(`
Task:

    Title: %v
    Description: %v
    Completed: %v
    ID: %v
            `, t.Title, t.Description, t.Completed, t.Id)
	}
}

func (tl *TaskList) ViewCompletedTasks() {
	fmt.Println("--- Completed Tasks ---")
	completedTasks := tl.GetCompletedTasks()
	for idx := range completedTasks {
		t := completedTasks[idx]
		fmt.Printf(`
Completed task:

    Title: %v 
    Description: %v
    Completed: %v
    ID: %v
            `, t.Title, t.Description, t.Completed, t.Id)
	}
}

// Return a slice containing pointers/references to the actual tasks in the main
// TaskList that are marked as complete.
func (tl *TaskList) GetCompletedTasks() []*Task {
	var ct []*Task
	for idx := range tl.Tasks {
		if tl.Tasks[idx].Completed {
			ct = append(ct, &tl.Tasks[idx])
		}
	}
	return ct
}

// TODO: Add UI for these.
// MarkComplete() marks a task as completed, changing its Completed property to true
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
