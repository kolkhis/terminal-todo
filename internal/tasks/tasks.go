package tasks

import (
	"fmt"
	"log"
	"os"
)

var Colors = map[string]string{
	"black":   "\033[30m",
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"magenta": "\033[35m",
	"cyan":    "\033[36m",
	"white":   "\033[37m",
	"":        "\033[0m",
}

var CompletedColor = map[bool]string{
	true:  Colors["green"],
	false: Colors["red"],
}

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

func (t *Task) formatOutput() string {
	outputStr := fmt.Sprintf(`
    - Task
      %vTitle: %v
      %vDescription: %v
      %vCompleted: %v
      %v
    `,
		Colors["cyan"],
		Colors[""]+
			t.Title,
		Colors["yellow"],
		Colors[""]+
			t.Description,
		CompletedColor[t.Completed],
		t.Completed,
		Colors[""],
	)
	return outputStr
}

func (tl *TaskList) DeleteTask(t Task) {}

// ViewTaskList Outputs the current task list to the terminal
// TODO: Add method (or option) to show only incomplete tasks, and only complete tasks
func (tl *TaskList) ViewTaskList() {
	fmt.Println("\033[32m--- All tasks ---\033[0m")
	for idx := range tl.Tasks {
		output := tl.Tasks[idx].formatOutput()
		fmt.Println(output)
	}
}

func (tl *TaskList) ViewCompletedTasks() {
	fmt.Println("--- Completed Tasks ---")
	completedTasks := tl.GetCompletedTasks()
	for idx := range completedTasks {
		output := completedTasks[idx].formatOutput()
		fmt.Println(output)
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
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Printf("Storage file does not exist at %s. Creating one.\n", path)
		os.Create(path)
	} else {
		fmt.Printf("fileInfo: %s\n", fileInfo)
		fmt.Printf("fileinfo.IsDir: %v\n", fileInfo.IsDir())
	}
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
