# Go Tutorial Projects

This directory contains two beginner-friendly Go projects designed to teach fundamental programming concepts and Go language features.

## Projects Overview

### 1. 🧠 Quizza - Interactive Quiz Game

A command-line quiz application that tests knowledge across various topics including math, history, science, and geography.

### 2. ✅ Todo App - Command-Line Task Manager

A feature-rich command-line todo application with file persistence and task management capabilities.

---

## 🧠 Quizza

### Features

- Interactive command-line quiz interface
- 10 diverse questions covering multiple topics
- Real-time scoring system
- Percentage calculation
- Case-insensitive answer matching
- Ability to quit game at any time
- Immediate feedback on answers

### Topics Covered

- **Mathematics**: Basic arithmetic
- **History**: World events and figures
- **Geography**: World capitals and countries
- **Science**: Chemistry, astronomy, and biology
- **Literature**: Famous authors and works

### How to Run

```bash
cd tutorial/quizza
go run main.go
```

### Sample Questions

1. What is 15 \* 12?
2. Who is the first president of the United States?
3. What is the capital of France?
4. What is the largest planet in our solar system?
5. What is the chemical symbol for gold?

### Learning Objectives

- Basic input/output operations
- String manipulation and formatting
- Loops and conditional statements
- Slice iteration
- Case-insensitive string comparison
- User interaction handling

---

## ✅ Todo App

### Features

- **Add Tasks**: Create new todo items
- **List Tasks**: View all todos with completion status
- **Edit Tasks**: Modify existing todo items
- **Toggle Completion**: Mark/unmark tasks as completed
- **Delete Tasks**: Remove unwanted todos
- **File Persistence**: Save and load todos from text files
- **Custom List Names**: Create multiple todo lists

### How to Run

```bash
cd tutorial/todo_app
go run main.go
```

### Usage Flow

1. **Start**: Enter a name for your todo list
2. **Load**: Automatically loads existing todos from file (if exists)
3. **Menu**: Choose from 6 available operations
4. **Save**: Save and exit when done

### Menu Options

```
Press 1 to add todo
Press 2 to list todo
Press 3 to edit todo
Press 4 to mark/unmark todo as completed
Press 5 to delete todo
Press 6 to save to file and exit
```

### File Format

Todos are saved in a human-readable format:

```
1 Name: Buy groceries ...Completed: No
2 Name: Finish project ...Completed: Yes
3 Name: Call dentist ...Completed: No
```

### Learning Objectives

- Struct definition and usage
- Method implementation
- File I/O operations
- String parsing and manipulation
- Menu-driven programming
- Error handling
- Input validation
- Array/slice operations

---

## Project Structure

```
tutorial/
├── quizza/
│   ├── data/
│   │   └── quiz_data.go    # Question data and structure
│   ├── go.mod              # Module definition
│   └── main.go             # Main application logic
├── todo_app/
│   ├── todo/
│   │   └── todo.go         # Todo logic and file operations
│   ├── go.mod              # Module definition
│   └── main.go             # Main application entry point
└── README.md               # This file
```

## Prerequisites

- Go 1.24.4 or higher
- Terminal/Command prompt
- Text editor (for viewing saved todo files)

## Getting Started

### Clone and Navigate

```bash
git clone <repository-url>
cd tutorial
```

### Run Quizza

```bash
cd quizza
go run main.go
```

### Run Todo App

```bash
cd todo_app
go run main.go
```

## Learning Path

### For Beginners

1. **Start with Quizza** - Simpler structure, focuses on basic I/O and control flow
2. **Progress to Todo App** - More complex with file operations and data persistence

### Key Go Concepts Demonstrated

#### In Quizza:

- Package organization
- Variable declarations
- Function definitions
- String formatting with `fmt`
- User input with `bufio`
- Slice iteration with `range`
- Conditional statements
- Type conversion

#### In Todo App:

- Struct definitions
- Method receivers
- File operations (`os.Open`, `os.Create`)
- Error handling patterns
- String manipulation
- Scanner usage
- Menu-driven architecture
- Data persistence

## Code Examples

### Quizza - User Input Function

```go
func getUserInput(prompt string) string {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print(prompt)
    input, _ := reader.ReadString('\n')
    return strings.TrimSpace(input)
}
```

### Todo App - Todo Structure

```go
type Todo struct {
    Name        string `json:"name"`
    isCompleted bool   `json:"isCompleted"`
}
```

## Extensions and Improvements

### Possible Enhancements for Quizza:

- Add difficulty levels
- Implement timer functionality
- Create different quiz categories
- Add high score tracking
- Implement multiplayer mode

### Possible Enhancements for Todo App:

- Add due dates
- Implement priority levels
- Add todo categories/tags
- Create JSON export/import
- Add search functionality
- Implement recurring tasks

## Best Practices Demonstrated

1. **Code Organization**: Logical separation of concerns
2. **Error Handling**: Graceful handling of file operations
3. **User Experience**: Clear prompts and feedback
4. **Data Validation**: Input sanitization and validation
5. **File Operations**: Safe file reading and writing
6. **Memory Management**: Efficient slice operations

## Troubleshooting

### Common Issues:

**File Permission Errors**

```bash
# Ensure write permissions in the directory
chmod 755 .
```

**Module Not Found**

```bash
# Initialize module if go.mod is missing
go mod init projectname
```

**Build Errors**

```bash
# Check Go version
go version

# Verify module integrity
go mod tidy
```

## Contributing

These tutorial projects welcome improvements and additional features! Please:

1. Fork the repository
2. Create a feature branch
3. Add clear comments to your code
4. Test thoroughly
5. Submit a pull request

## Learning Resources

- [Go Official Documentation](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go Tour](https://tour.golang.org/)

## License

These tutorial projects are open source and available under the [MIT License](LICENSE).
