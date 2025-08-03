package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Se7enSe7enSe7en/todo-cli/internal/logger"
)

var CommandMap = map[string]string{
	"create": "create",
	"list":   "list",
}

const (
	NUM_OF_COLUMNS = 5
	FILE_PATH      = "./todo.csv"
)

type Priority uint8

const (
	Low Priority = iota
	Medium
	High
)

type Todo struct {
	id          string
	description string
	done        bool
	createdAt   time.Time
	priority    Priority
}

type TodoList interface {
	add()
}

func OpenCsvFile() *os.File {
	// Try to open existing file
	csvFile, err := os.OpenFile(FILE_PATH, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	return csvFile
}

func generateID() string {
	return time.Now().Format("20060102150405")
}

// Convert Todo struct to string slice for CSV storage
func toStringSlice(todo Todo) []string {
	// Convert bool to string
	doneStr := map[bool]string{true: "true", false: "false"}[todo.done]

	// Convert priority to string
	priorityStr := fmt.Sprintf("%d", todo.priority)

	// Convert to string slice
	return []string{
		todo.id,
		todo.description,
		doneStr,
		todo.createdAt.Format(time.RFC3339),
		priorityStr,
	}
}

func createTodo(todoList [][]string, description string) ([][]string, error) {
	// Generate a unique ID
	id := generateID()

	// Create new todo using the Todo struct
	newTodo := Todo{
		id:          id,
		description: description,
		done:        false,
		createdAt:   time.Now(),
		priority:    Low, // default priority
	}

	// Convert Todo struct to string slice and add to todoList
	todoList = append(todoList, toStringSlice(newTodo))

	return todoList, nil
}

func main() {
	logger.Initialize()
	logger.Info("VIBE CHECKKKKKKKKKKKKKKKKKKKKKK")

	args := os.Args[1:]

	var command string
	var index int
	for i, arg := range args {
		commandArg, ok := CommandMap[arg]
		if ok {
			fmt.Println("COMMAND: ", commandArg) // TEST
			command = commandArg
			index = i
			break
		}
	}

	csvFile := OpenCsvFile()
	defer csvFile.Close()

	switch command {
	case "create":
		// get the following arguments after the command
		createArg := strings.Join(args[index+1:], " ")
		fmt.Printf("createArg: %v\n", createArg)

		// init reader
		reader := csv.NewReader(csvFile)
		reader.FieldsPerRecord = -1

		// load todoList from the csv
		todoList, err := reader.ReadAll()
		if err != nil {
			log.Fatal("cannot read the csv: ", err)
		}

		// if we got empty, then initialize csv with the headers
		if len(todoList) == 0 {
			headers := []string{"id", "description", "done", "created at", "priority"}
			todoList = append(todoList, headers)
		}

		// init writer
		writer := csv.NewWriter(csvFile)
		defer writer.Flush()

		// create todo
		todoList, err = createTodo(todoList, createArg)
		if err != nil {
			log.Fatal("cannot create todo: ", err)
		}

		// Reset file position and clear existing content before writing
		csvFile.Seek(0, 0)
		csvFile.Truncate(0)

		// write back to csv
		if err := writer.WriteAll(todoList); err != nil {
			log.Fatal("cannot write to csv: ", err)
		}

		fmt.Println("Todo created successfully!")
	}

}
