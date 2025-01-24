package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	t "github.com/kolkhis/terminal-todo/internal/tasks"
)

// TODO: Check out BubbleTea

func main() {

	var tl t.TaskList = t.NewTaskList()
	tl.LoadTaskList()
	if len(os.Args) > 1 {
		tl.ParseArgs()
	}

	fmt.Println("Terminal TODO")
	choice, err := GetMainMenuInput()

	if err != nil {
		log.Fatalf("Error while getting input: %v\n", err)
	}

	s := bufio.NewScanner(os.Stdin)
	switch choice {
	case 0:
		// User entered zero or a non-digit
		fmt.Println("Exiting.")
		os.Exit(0)

	case 1:
		newTask, err := tl.GetNewTaskInput()
		if err != nil {
			log.Fatalf("Failed to get input for new task.\nReason: %v\n", err)
		}
		tl.AddTaskToList(newTask)
		tl.SaveTaskList()
		fmt.Println("Current tasks:")
		tl.ViewTaskList()

	case 2:

		fmt.Println("--- Remove a task ---")
		fmt.Print("Enter task ID:\n> ")

		// s.Scan()
		// taskToDel := s.Text()
		// id, err := strconv.ParseInt(taskToDel, 10, 0)
		// if err != nil {
		// 	fmt.Printf("Error in ParseInt: %v\n", err)
		// } else {
		// 	tl.DeleteTask(int(id))
		// 	tl.SaveTaskList()
		// 	fmt.Println("Task successfully deleted.")
		// }
		var taskId int
		fmt.Scanln(&taskId)
		tl.DeleteTask(taskId)

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
		fmt.Println("Showing task list")
		tl.ViewTaskList()
	case 5:
		fmt.Println("Showing incomplete tasks.")
		tl.ViewIncompleteTasks()
	case 6:
		fmt.Println("Showing complete tasks.")
		tl.ViewCompleteTasks()
	case 7:
		tl.SaveTaskList()
	}

}

func GetMainMenuInput() (int, error) {
	fmt.Print(`
Select an option:
      1. Create a new task
      2. Remove a task
      3. Mark a task as complete
      4. View task list
      5. View incomplete tasks
      6. View complete tasks
      7. Save task list to file

    Blank, 'q', or '0' to quit. 

> `)
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	userInput := s.Text()

	if userInput == "q" || userInput == "0" {
		return 0, nil
	}

	parsedInput, err := strconv.Atoi(userInput)
	if err != nil || parsedInput < 1 || parsedInput > 7 {
		fmt.Println("Invalid input.")
		return 0, fmt.Errorf("Unable to parse input.\n")
	}

	return parsedInput, nil
}
