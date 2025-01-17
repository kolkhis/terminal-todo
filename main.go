package main

import (
	"bufio"
	"fmt"
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
	tl.LoadTaskList() // if it exists - TODO: Add file support

	fmt.Println("Terminal TODO")

	fmt.Print(`
Select an option:
      1. Create a new task
      2. Remove a task
      3. Mark a task as complete
      4. View task list
      5. Save task list to file

    'q' or '0' to quit. 

> `)

	var choice int

	fmt.Scanln(&choice)

	s := bufio.NewScanner(os.Stdin)
	switch choice {
	case 0:
		// User entered zero or a non-digit
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
				newTask := t.NewTask(title, desc)
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
		s.Scan()
		taskToDel := s.Text()

		id, err := strconv.ParseInt(taskToDel, 10, 0)
		if err != nil {
			fmt.Printf("Error in ParseInt: %v\n", err)
		} else {
			tl.DeleteTask(int(id))
			tl.SaveTaskList()
		}

	case 3:
		fmt.Println("Mark a task as complete.")
	case 4:
		fmt.Println("View task list")
		tl.ViewTaskList()
	case 5:
		tl.SaveTaskList()
	}

}
