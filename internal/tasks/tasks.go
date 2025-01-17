package tasks

import (
	"errors"
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

func (tl *TaskList) DeleteTask(t Task) {
	for i := range len(tl.Tasks) {
		if t.Id == tl.Tasks[i].Id {
			// remove from list
			// tl.Tasks[i]
		}
	}
}

func (tl *TaskList) AddTaskToList(t Task) {
	t.Id = tl.nextTaskId
	tl.Tasks = append(tl.Tasks, t)
	tl.nextTaskId++
}

func (t *Task) formatTaskOutput() string {
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

// ViewTaskList Outputs the current task list to the terminal
// TODO: Add method (or option) to show only incomplete tasks, and only complete tasks
func (tl *TaskList) ViewTaskList() {
	fmt.Println("\033[32m--- All tasks ---\033[0m")
	for idx := range tl.Tasks {
		output := tl.Tasks[idx].formatTaskOutput()
		fmt.Println(output)
	}
}

func (tl *TaskList) ViewCompletedTasks() {
	fmt.Println("--- Completed Tasks ---")
	completedTasks := tl.GetCompletedTasks()
	for idx := range completedTasks {
		output := completedTasks[idx].formatTaskOutput()
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

// Check the storage file path, create it if it doesn't exist.
func checkStorageFileDir(basepath string) bool {
	fmt.Printf("\033[31m --- Base path for storage file: %v ---\n\033[0m", basepath)
	// Create file if it's not there
	// Create directory if it's not there
	fileInfo, err := os.Stat(basepath)
	if err != nil {
		fmt.Printf("Failed to stat directory at %v: %v\n", basepath, err)
		fmt.Printf("Attempting to create directory: %v\n", basepath)
		err = os.MkdirAll(basepath, os.FileMode(0755))
		if err != nil {
			fmt.Printf(Colors["red"]+"Error creating directory at %v: %v\n"+Colors[""], basepath, err)
			return false
		} else {
			fmt.Printf("\033[32mSuccessfully created directory: %v\n", basepath)
		}
	} else if fileInfo.IsDir() {
		return true
	}
	fmt.Printf("Couldn't determine the existence of the directory: %v\n", basepath)
	return false
}

func checkStorageFile(basepath string) bool {
	storageFile := basepath + "tasklist.json"
	_, err := os.Stat(storageFile)
	if err != nil {
		fmt.Printf(Colors["red"]+"Failed to stat file at %v: %v\n"+Colors[""], storageFile, err)
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf(Colors["red"]+"Error - File %v doesn't exist: %v\n"+Colors[""], storageFile, err)
			fmt.Printf("Storage file does not exist at %s.\nCreating one...\n", storageFile)
			return false
		} else {
			fmt.Printf("Ran into error while checking directory: %v\n", err)
		}
	}

	_, err = os.Create(storageFile)

	if err != nil {
		fmt.Printf("Ran into error when creating file %v: %v\n", storageFile, err)
		return false
	} else {
		return true
	}
}

func (*TaskList) LoadTaskList() {
	basepath := os.Getenv(`HOME`) + "/.config/terminal-todo/" // Don't use tilde, no pathname expansion
	if !checkStorageFileDir(basepath) {
		log.Fatalf("Unable to create the storage file directory in %v\n", basepath)
	}
	if !checkStorageFile(basepath) {
		log.Fatalf("Unable to load storage file at %v\n", basepath)
	}

	storageFile := basepath + "tasklist.json"
	log.Printf("Opening file at: %v\n", storageFile)

	file, err := os.Open(storageFile)
	if err != nil {
		log.Fatalf("Couldn't open file storage file at %v\n", storageFile)
	}
	defer file.Close()
	// TODO: Add logic to check if file is empty -
	// TODO: If not empty, parse contents into a TaskList()

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
// TODO: Add UI for these.
// MarkComplete() marks a task as completed, changing its Completed property to true
func (t *Task) MarkComplete()                    { t.Completed = true }
func (t *Task) ChangeDescription(newDesc string) { t.Description = newDesc }
func (t *Task) ChangeTitle(newTitle string)      { t.Title = newTitle }
