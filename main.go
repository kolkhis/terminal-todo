package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	t "github.com/kolkhis/terminal-todo/internal/tasks"
)

// func HandleArgs() {
//     if len(os.Args) > 0 {
//         switch os.Args[0] {
//         case "add":
//         }
//     }
// }

func main() {
	var tl t.TaskList = t.NewTaskList()
	tl.LoadTaskList()
	fmt.Println("Terminal TODO")
	choice := GetMainMenuInput()

	s := bufio.NewScanner(os.Stdin)
	switch choice {
	case 0:
		// User entered zero or a non-digit
		log.Println("Case 0 hit.")
		fmt.Println("Exiting.")
		os.Exit(0)
	case 1:
		fmt.Println("--- Add a task ---")
		fmt.Print("New task name: ")
		s.Scan()
		title := s.Text()

		fmt.Print("New Task description: ")
		s.Scan()
		desc := s.Text()

		newTask := t.NewTask(title, desc)
		fmt.Println("New Task Created:")
		newTask.ViewTask()

		var confirmation string
		for confirmation != "y" && confirmation != "n" && confirmation != "q" {
			fmt.Printf("Create this task? [y/N (q to quit)] ")
			s.Scan()
			confirmation = s.Text()
			switch confirmation {
			case "y":
				// newTask := t.NewTask(title, desc)
				tl.AddTaskToList(newTask)
				tl.SaveTaskList()
				break
			case "n":
				fmt.Println("OK - Discarding task.")
				break
			case "q":
				fmt.Println("Exiting.")
				os.Exit(0)
				break
			default:
				fmt.Println("Invalid selection. ")
			}
		}

		fmt.Println("Current tasks:")
		tl.ViewTaskList()
	case 2:
		fmt.Println("--- Remove a task ---")
		fmt.Print("Enter task ID:\n> ")

		// var taskId int
		// fmt.Scanln(&taskId)
		// tl.DeleteTask(taskId)

		s.Scan()
		taskToDel := s.Text()
		id, err := strconv.ParseInt(taskToDel, 10, 0)
		if err != nil {
			fmt.Printf("Error in ParseInt: %v\n", err)
		} else {
			tl.DeleteTask(int(id))
			tl.SaveTaskList()
			fmt.Println("Task successfully deleted.")
		}

	case 3:
		fmt.Println("Mark a task as complete.")
		fmt.Print("Enter Task ID:\n> ")
		s.Scan()
		taskId := s.Text()
		parsedId, err := strconv.Atoi(taskId)
		if err != nil {
			fmt.Printf("Couldn't parse input: %v\n", err)
		}
		tl.SetComplete(parsedId, true)
		tl.SaveTaskList()

	case 4:
		fmt.Println("View task list")
		tl.ViewTaskList()
	case 5:
		tl.SaveTaskList()
	}

}

func GetMainMenuInput() int {
	fmt.Print(`
Select an option:
      1. Create a new task
      2. Remove a task
      3. Mark a task as complete
      4. View task list
      5. Save task list to file

    'q' or '0' to quit. 

> `)
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	userInput := s.Text()

	if userInput == "q" || userInput == "0" {
		return 0
	}

	parsedInput, err := strconv.Atoi(userInput)
	if err != nil || parsedInput < 1 || parsedInput > 5 {
		log.Println("Invalid input.")
		return 0
	}

	return parsedInput
}

func GetNewTaskInput() t.Task {
	// s := bufio.NewScanner(os.Stdin)
	return t.Task{}
}
