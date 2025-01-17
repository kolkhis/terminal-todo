# Terminal TODO

An educational project.  

The idea is to make an interactive TODO list in the terminal using Go.  




## Features
* Add a task to the lsit
* Remove a task from the lsit
* Mark a task as complete on the list
* View tasks with their statuses.
* Save tasks to a list file so it persists between sessions.  
* Load tasks from the file when the program starts.  

Maybe:
- Ability to have open in a tmux pane and dynamically modify data
    - Run constantly and listen for user input
- Subtasks

## Concepts
The task list will be a slice.  
Each task will be a struct of type `Task`.  
Pass the task list itself (using pointers) to functions in order to manipulate it.  
File I/O for managing the list of tasks


Add list of commands to modify task list
e.g.,
```bash
todo add 
todo finish
todo del
todo modify
```


## Ideas
Choice between two different modes
- Interactive Mode
- One-off ad-hoc command mode
- Support for multiple task lists


