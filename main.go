package main

import (
	"bufio"
	"fmt"
	"os"

	t "github.com/kolkhis/terminal-todo/internal/tasks"
)

func GetInputForTask() {}

func main() {
	var tl t.TaskList = t.NewTaskList()
	tl.LoadTaskList()
	// tl := t.NewTaskList()

	fmt.Println("Terminal TODO")

	newT := t.NewTask("Finish project", "Finish the cli task list project", 0)
	tl.AddTaskToList(newT)
	fmt.Printf("Current task list:\n")
	tl.ViewTaskList()

	fmt.Println(`
	Select an option:
	1. Create a new task
	2. Remove a task
	3. Mark a task as complete
	4. View task list
	`)

	var choice int
	fmt.Scanln(&choice)

	s := bufio.NewScanner(os.Stdin)
	switch choice {
	case 1:
		fmt.Println("--- Add a task ---")
		fmt.Print("New task name: ")
		s.Scan()
		title := s.Text()

		fmt.Print("New Task description: ")
		s.Scan()
		desc := s.Text()
		id := len(tl.Tasks) + 1

		fmt.Printf(`
        New Task Created:
        Title: %v
        Description: %v
        
        `, title, desc)

		var confirmation string
		for confirmation != "y" && confirmation != "n" && confirmation != "q" {
			fmt.Printf("Create this task? [y/N (q to quit)] ")
			s.Scan()
			confirmation = s.Text()
			switch confirmation {
			case "y":
				newTask := t.NewTask(title, desc, id)
				tl.AddTaskToList(newTask)
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
		fmt.Println("Remove a task")
	case 3:
		fmt.Println("Mark a task as complete.")
	case 4:
		fmt.Println("View task list")
		for t := 0; t < len(tl.Tasks); t++ {
			fmt.Printf("Title: %v\n", tl.Tasks[t].Title)
			fmt.Printf("Description:%v\n", tl.Tasks[t].Description)
			fmt.Printf("Completed: %v\n\n", tl.Tasks[t].Completed)
		}

	}

}
