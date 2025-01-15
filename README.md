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


## Concepts
The task list will be a slice.  
Each task will be a struct of type `Task`.  
Pass the task list itself (using pointers) to functions in order to manipulate it.  
File I/O for managing the list of tasks







