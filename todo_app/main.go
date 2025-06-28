package main

import (
	"fmt"
	"todoapp/todo"
)

func userPromptOptions(todoName string) {
	userOption := todo.AcceptInput("\nPress 1 to add todo \nPress 2 to list todo \nPress 3 to edit todo \nPress 4 to mark/unmark todo as completed \nPress 5 to delete todo \nPress 6 to save to file and exit \n\nSelect an option: ")
	switch userOption {
	case "1":
		todo.AddTodo(todo.AcceptInput("\nPlease enter your todo item name: "))
	case "2":
		todo.ListTodos()
	case "3":
		todo.EditTodo()
	case "4":
		todo.ToggleTodoCompleted()
	case "5":
		todo.DeleteTodo()
	case "6":
		err := todo.SaveToFile(todoName)
		if err != nil {
			fmt.Println("Failed to save to file", err)
		} else {
			fmt.Println("Successfully saved to file")
		}
		return
	default:
		fmt.Println("You have selected an invalid option")

	}

	userPromptOptions(todoName)
}

func main() {
	fmt.Println("Welcome to todo app")

	todoName := todo.AcceptInput("Please enter your todo list name: ")
	fmt.Println("Your todo name is: ", todoName)

	todo.LoadTodosFromTextFile(todoName + ".txt")
	userPromptOptions(todoName)

}
