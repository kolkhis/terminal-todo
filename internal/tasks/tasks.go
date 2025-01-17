package tasks

import (
	"encoding/json"
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
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	Id          int    `json:"id"`
}

type TaskList struct {
	Tasks       []Task `json:"tasks"`
	NextTaskId  int    `json:"nextTaskId"`
	storageFile string
	// TODO: Map task IDs to tasks for easier management/deletion
	// TaskIndex map[int]*Task
}

func NewTaskList() TaskList {
	tl := TaskList{
		NextTaskId: 0,
	}
	storageDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatalf("Couldn't set the storage file for task list: %v\n", err)
	}
	tl.storageFile = storageDir + "/terminal-todo/tasklist.json"
	return tl
}

// Constructor for Task
func NewTask(title, desc string) Task {
	t := Task{
		Title:       title,
		Description: desc,
		Completed:   false,
	}
	return t
}

func (tl *TaskList) AddTaskToList(t Task) {
	t.Id = tl.NextTaskId
	tl.Tasks = append(tl.Tasks, t)
	// tl.TaskIndex[t.Id] = &tl.Tasks[len(tl.Tasks)-1]
	tl.NextTaskId++
}

func (tl *TaskList) DeleteTask(id int) {
	log.Printf("START: Number of tasks: %v\n", len(tl.Tasks))
	for i := len(tl.Tasks) - 1; i >= 0; i-- {
		if id == tl.Tasks[i].Id {
			tl.Tasks = append(tl.Tasks[:i], tl.Tasks[i+1:]...) //TODO: Handle when it's the last element
			log.Printf("%vAppend operation complete%v\n", Colors["green"], Colors[""])
			tl.ViewTaskList()
			log.Printf("DONE: Number of tasks: %v\n", len(tl.Tasks))
			return
		}

	}
}

func (t *Task) formatTaskOutput() string {
	outputStr := fmt.Sprintf(`
    - Task
      %vTitle: %v
      %vDescription: %v
      %vCompleted: %v
      ID: %v
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
		t.Id,
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
func checkStorageFileDir() (bool, error) {
	// Create file if it's not there
	// Create directory if it's not there
	configDir, err := os.UserConfigDir()
	if err != nil {
		return false, fmt.Errorf("Error while trying to access the user's configuration dir: %v\n", err)
	}
	basepath := configDir + "/terminal-todo/"

	fmt.Printf("\033[31m --- Base path for storage file: %v ---\n\033[0m", basepath)
	fileInfo, err := os.Stat(basepath)
	if err != nil {
		fmt.Printf("Failed to stat directory at %v: %v\n", basepath, err)
		fmt.Printf("Attempting to create directory: %v\n", basepath)
		err = os.MkdirAll(basepath, os.FileMode(0755))
		if err != nil {
			fmt.Printf(Colors["red"]+"Error creating directory at %v: %v\n"+Colors[""], basepath, err)
			return false, fmt.Errorf("Failed to create directory at %v\n", basepath)
		} else {
			fmt.Printf("\033[32mSuccessfully created directory: %v\n", basepath)
		}
	} else if fileInfo.IsDir() {
		return true, nil
	}
	fmt.Printf("Couldn't determine the existence of the directory: %v\n", basepath)
	return false, fmt.Errorf("Failed to find or create directory for storage file.\n")
}

func (tl *TaskList) checkStorageFile() (bool, error) {

	info, err := os.Stat(tl.storageFile)
	if err != nil {
		fmt.Printf(Colors["red"]+"Failed to stat file at %v: %v\n"+Colors[""], tl.storageFile, err)
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf(Colors["red"]+"Error - File %v doesn't exist: %v\n"+Colors[""], tl.storageFile, err)
			fmt.Printf("Storage file does not exist at %s.\nCreating one...\n", tl.storageFile)
		} else {
			fmt.Printf("Ran into error while checking directory: %v\n", err)
		}
	} else {
		if info.Mode().IsRegular() {
			fmt.Printf("Regular file found at %v\n", info.Name())
			return true, nil
		}
		return false, fmt.Errorf("Storage file %v is not a regular file. FileMode: %v\n", info.Name(), info.Mode())
	}
	f, err := os.Create(tl.storageFile)

	if err != nil {
		fmt.Printf("Ran into error when creating file %v: %v\n", tl.storageFile, err)
		return false, err
	} else {
		f.Chmod(0777)
		return true, nil
	}
}

func (tl *TaskList) LoadTaskList() {
	if _, err := checkStorageFileDir(); err != nil {
		log.Fatalf("Unable to create the storage file directory: %v\n", err)
	}
	if _, err := tl.checkStorageFile(); err != nil {
		log.Fatalf("Unable to load or create storage file: %v\n", err)
	}

	storageFileDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatalf("Couldn't fetch user config dir: %v\n", err)
	}

	storageFile := storageFileDir + "/terminal-todo/tasklist.json"
	log.Printf("Opening storage file - %v", storageFile)

	file, err := os.Open(storageFile)
	if err != nil {
		log.Fatalf("Couldn't open file storage file at %v\n", storageFile)
	}
	defer file.Close()

	// Check if file is empty
	fileInfo, err := file.Stat()
	if err != nil {
		log.Printf("Error while trying to stat file %v: %v", file, err)
	}

	log.Printf("Size of the file: %v", fileInfo.Size())
	if fileInfo.Size() == 0 {
		log.Printf("Storage file %v is empty. Using new task list.\n", fileInfo.Name())
		return
	}
	decoder := json.NewDecoder(file)
	if err = decoder.Decode(tl); err != nil {
		log.Fatalf("Failed to read in contents of json file: %v\n", err)
	}

	fmt.Printf(Colors["green"]+"New Task list (from json):\n%v\n", tl)
	fmt.Printf(Colors["green"]+"nextTaskID: %v \n", tl.NextTaskId)

	tl.ViewTaskList()
	// TODO: If not empty, parse contents into a TaskList()
	//       Handle NextTaskId by len(tl.Tasks) + 1
	// TODO: Convert json contents into valid Go data structure
	//
	// Use encoding/json -
	//      Marshaling: Convert struct to JSON
	//      Unmarshaling: Convert JSON to struct

}

func (tl *TaskList) SaveTaskList() {
	file, err := os.OpenFile(tl.storageFile, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		log.Fatalf("Failed to open file for writing: %v\n", err)
	}
	defer file.Close()
	jsonData, err := json.MarshalIndent(*tl, "", "\t")
	if err != nil {
		log.Fatalf("Couldn't marshal data to json: %v\n", err)
	}
	fmt.Printf("Writing task list to file %v\n", tl.storageFile)
	fmt.Printf("Data being written:\n%v\n", tl)
	fmt.Printf("jsonData being written:\n%v\n", jsonData)
	bw, err := file.Write(jsonData)
	if err != nil {
		log.Fatalf("Couldn't write to file: %v\n", err)
	}
	fmt.Printf("%v bytes written.\n", bw)
}

// TODO: Add UI for these.
// MarkComplete() marks a task as completed, changing its Completed property to true
func (t *Task) SetComplete(complete bool)     { t.Completed = complete }
func (t *Task) SetDescription(newDesc string) { t.Description = newDesc }
func (t *Task) SetTitle(newTitle string)      { t.Title = newTitle }
