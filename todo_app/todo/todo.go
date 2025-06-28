package todo

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Todo struct {
	Name        string `json:"name"`
	isCompleted bool   `json:"isCompleted"`
}

var todos []Todo

func LoadTodosFromTextFile(filename string) error {
	// Check if file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil // File doesn't exist, nothing to load
	}

	// Open file for reading
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		// Expected format: "1 Name: Buy groceries ...Completed: No"
		parts := strings.Split(line, "Name: ")
		if len(parts) < 2 {
			continue
		}
		nameCompleted := parts[1]
		nameParts := strings.Split(nameCompleted, " ...Completed: ")
		if len(nameParts) < 2 {
			continue
		}
		name := nameParts[0]
		completedStr := nameParts[1]
		isCompleted := strings.TrimSpace(strings.ToLower(completedStr)) == "yes"

		todos = append(todos, Todo{
			Name:        name,
			isCompleted: isCompleted,
		})
	}

	return scanner.Err()
}

func convertTodoInputToInt(input string) uint {
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("You have entered an invalid todo number")
		return 0
	}

	returnNum := uint(num - 1)
	return returnNum
}

func AcceptInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func AddTodo(todoItem string) {
	newTodo := Todo{Name: todoItem, isCompleted: false}
	todos = append(todos, newTodo)
}

func ListTodos() {
	if len(todos) == 0 {
		fmt.Println("\nYou have no todos at the moment")
		return
	}

	fmt.Println("\n These are your todos: ")
	for index, todo := range todos {
		var todoCompleted string
		if todo.isCompleted {
			todoCompleted = "Yes"
		} else {
			todoCompleted = "No"
		}
		fmt.Printf("%v %-25v ...%v \n", index+1, "Name: "+todo.Name, "Completed: "+fmt.Sprint(todoCompleted))
	}
}

func EditTodo() {
	todoNumberInput := AcceptInput("\nPlease enter the number of the todo you want to edit: ")
	todoNumber := convertTodoInputToInt(todoNumberInput)

	todos[todoNumber].Name = AcceptInput("\nPlease enter your todo item name: ")
}

func ToggleTodoCompleted() {
	todoNumberInput := AcceptInput("\nPlease enter the number of the todo you want to mark as completed: ")
	todoNumber := convertTodoInputToInt(todoNumberInput)

	todos[todoNumber].isCompleted = !todos[todoNumber].isCompleted
}

func DeleteTodo() {
	todoNumberInput := AcceptInput("\nPlease enter the number of the todo you want to delete: ")
	todoNumber := convertTodoInputToInt(todoNumberInput)

	todos = append(todos[:todoNumber], todos[todoNumber+1:]...)
}

func SaveToFile(name string) error {
	filename := name + ".txt"
	// Open or create the file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// If no todos
	if len(todos) == 0 {
		_, err = file.WriteString("You have no todos at the moment\n")
		return err
	}

	// Write each todo in your custom format
	for index, todo := range todos {
		todoCompleted := "No"
		if todo.isCompleted {
			todoCompleted = "Yes"
		}
		line := fmt.Sprintf("%v %-25v ...%v\n", index+1, "Name: "+todo.Name, "Completed: "+todoCompleted)
		_, err := file.WriteString(line)
		if err != nil {
			return err
		}
	}

	return nil

}
