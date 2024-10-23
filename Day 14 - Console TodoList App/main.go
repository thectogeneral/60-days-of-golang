package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Task represents a to-do item.
type Task struct {
	ID        int
	Title     string
	Completed bool
}

// ToDoList manages a list of tasks and an ID tracker.
type ToDoList struct {
	Tasks  []Task
	NextID int
}

// AddTask adds a new task to the list.
func (t *ToDoList) AddTask(title string) {
	t.Tasks = append(t.Tasks, Task{ID: t.NextID, Title: title, Completed: false})
	t.NextID++
	fmt.Printf("Task '%s' added with ID %d\n", title, t.NextID-1)
}

// MarkTaskComplete marks a task as completed by ID.
func (t *ToDoList) MarkTaskComplete(id int) {
	for i, task := range t.Tasks {
		if task.ID == id {
			t.Tasks[i].Completed = true
			fmt.Printf("Task '%s' marked as complete\n", task.Title)
			return
		}
	}
	fmt.Println("Task not found.")
}

// ListTasks displays all tasks.
func (t *ToDoList) ListTasks() {
	if len(t.Tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}
	fmt.Println("Tasks:")
	for _, task := range t.Tasks {
		status := "Pending"
		if task.Completed {
			status = "Completed"
		}
		fmt.Printf("%d: %s [%s]\n", task.ID, task.Title, status)
	}
}

// DeleteCompletedTasks removes all completed tasks.
func (t *ToDoList) DeleteCompletedTasks() {
	var remainingTasks []Task
	for _, task := range t.Tasks {
		if !task.Completed {
			remainingTasks = append(remainingTasks, task)
		}
	}
	if len(t.Tasks) == len(remainingTasks) {
		fmt.Println("No completed tasks to delete.")
	} else {
		t.Tasks = remainingTasks
		fmt.Println("Completed tasks deleted.")
	}
}

// DisplayMenu shows available options.
func DisplayMenu() {
	fmt.Println("\nTo-Do List Application")
	fmt.Println("1. Add a task")
	fmt.Println("2. Mark task as complete")
	fmt.Println("3. List all tasks")
	fmt.Println("4. Delete completed tasks")
	fmt.Println("5. Exit")
	fmt.Print("Select an option: ")
}

// HandleInput processes user commands.
func HandleInput(todolist *ToDoList) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		DisplayMenu()

		// Capture user input
		scanner.Scan()
		option := strings.TrimSpace(scanner.Text())

		switch option {
		case "1":
			fmt.Print("Enter task description: ")
			scanner.Scan()
			title := strings.TrimSpace(scanner.Text())
			if title == "" {
				fmt.Println("Error: Task description cannot be empty.")
			} else {
				todolist.AddTask(title)
			}

		case "2":
			fmt.Print("Enter task ID to mark complete: ")
			scanner.Scan()
			idStr := strings.TrimSpace(scanner.Text())
			id, err := strconv.Atoi(idStr)
			if err != nil || id < 0 {
				fmt.Println("Error: Invalid task ID.")
			} else {
				todolist.MarkTaskComplete(id)
			}

		case "3":
			todolist.ListTasks()

		case "4":
			todolist.DeleteCompletedTasks()

		case "5":
			fmt.Println("Exiting the application.")
			return

		default:
			fmt.Println("Error: Invalid option.")
		}
	}
}

func main() {
	todolist := &ToDoList{NextID: 0}
	HandleInput(todolist)
}
