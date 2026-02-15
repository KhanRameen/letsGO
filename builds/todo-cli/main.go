package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	//cli based todo app
	if len(os.Args) < 2 {
		fmt.Println("Usage: todo  [add|list|done|delete]")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		addTodo(os.Args[2:])
	case "list":
		listTodos()
	case "done":
		markDone(os.Args[2:])
	case "delete":
		deleteTodo(os.Args[2:])
	default:
		fmt.Println("Unknown command")
	}
}

func addTodo(args []string) {
	if len(args) < 1 {
		fmt.Println("Provide todo title")
		return
	}

	todos, _ := loadTodos()

	newTodo := Todo{
		ID:        len(todos) + 1,
		Title:     args[0],
		Completed: false,
	}

	todos = append(todos, newTodo)
	saveTodos(todos)

	fmt.Println("Todo Added!")
}

func listTodos() {
	todos, _ := loadTodos()

	for _, todo := range todos {
		status := " "
		if todo.Completed {
			status =  "✓"
		}
		fmt.Printf("[%s] %d: %s\n", status, todo.ID, todo.Title)

	}

}


func markDone(args []string) {
	if len(args)==0{
		fmt.Println("Provide todo ID")
		return
	}

	id, _ := strconv.Atoi(args[0])
	todos, _ := loadTodos()

	for i:= range todos {
		if todos[i].ID == id {
			todos[i].Completed = true
			break
		}
	}

	saveTodos(todos)
	fmt.Println("Todo Marked as Done!")
}


func deleteTodo(args []string) {
	if len(args)==0{
		fmt.Println("Provide todo ID")
		return
	}	

	id , _ := strconv.Atoi(args[0])
	todos, _ := loadTodos()

	var updatedTodos []Todo
	for _, todo := range todos {
		if todo.ID != id {
			updatedTodos = append(updatedTodos, todo)
		}
	}
	saveTodos(updatedTodos)
	fmt.Println("Todo Deleted!")
}