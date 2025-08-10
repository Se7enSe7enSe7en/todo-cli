package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Se7enSe7enSe7en/todo-cli/internal/logger"
	"github.com/Se7enSe7enSe7en/todo-cli/internal/stringUtil"
	"github.com/Se7enSe7enSe7en/todo-cli/pkg/set"
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
		logger.Error("Cannot create file", err)
	}
	return csvFile
}

func generateID(todoListStr [][]string) string {
	/*
		TODO: Id generation should be simpler, since we will use the id for accessing the todo

		algorithm # 1: find the lowest available number and use that for the id
			- [x] put the ids in a set for (O(1) look up time)
			- [x] create a loop starting from 1, iterate by 1, lookup the number
			in your set
			- [x] if it exists, then move to the next, if it doesn't, return the number
	*/

	const idIndex = 0
	idSet := set.NewSet()

	var maxId int = 0

	for index, todo := range todoListStr {
		if index == 0 {
			// skip header row
			continue
		}

		idStr := todo[idIndex]
		logger.Debug("idStr: %v", idStr)

		// add id to the set
		idSet.Add(idStr)

		// find the highest id number
		idInt, err := strconv.Atoi(idStr)
		if err != nil {
			logger.Error("generateId() ParseInt error: ", err)
			return ""
		}
		if idInt > maxId {
			maxId = idInt
		}
	}

	logger.Debug("idSet: %v", idSet.List())

	// check all numbers from 1 to maxId (highest id number)
	for num := 1; num < maxId; num++ {
		numStr := strconv.Itoa(num)

		// if the number is available, return the number as the id
		if !idSet.Contains(numStr) {
			return numStr
		}
	}

	// if no available low number id is found, return the maxId + 1 as the id
	return strconv.Itoa(maxId + 1)
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
	id := generateID(todoList)

	logger.Debug("generated id: %v", id)

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

	logger.Debug("todoList: %v", todoList)

	return todoList, nil
}

func listTodo(todoListStr [][]string) {
	/*
		TODO: make table in terminal
		- [x] make sure the it looks nice
		- [x] the width of the column automatically adjusts depending on the longest cell
		- [] there should be a limit for the width column, and wrap the text when it gets
		too long, when it is, add additional row (adjust vertically)
	*/
	// Get the dimensions of the todoList
	numOfRows := len(todoListStr)
	numOfColumns := len(todoListStr[0]) // it is guranteed that all rows have the same amount of columns

	// Init a 2d list for the formatted todo list
	todoListStrWithPadding := make([][]string, numOfRows)
	for i := range todoListStrWithPadding {
		todoListStrWithPadding[i] = make([]string, numOfColumns)
	}

	// column traversal (down to right)
	for col := 0; col < numOfColumns; col++ {
		maxColumnWidth := 0

		strColumn := []string{}

		for row := 0; row < numOfRows; row++ {

			cell := todoListStr[row][col]

			// logger.Debug("todoListStr[%v][%v]: %v", y, x, cell)

			strColumn = append(strColumn, cell)

			if len(cell) > maxColumnWidth {
				maxColumnWidth = len(cell)
			}
		}

		// logger.Debug("maxColumnWidth for column %v: %v", x, maxColumnWidth)

		for row, str := range strColumn {
			todoListStrWithPadding[row][col] = stringUtil.AddBothSidesPadding(stringUtil.AddRightSidePadding(str, maxColumnWidth))
		}

	}

	// logger.Debug("todoListStrWithPadding: %v", todoListStrWithPadding)

	// Print each row with the pipes and padding
	for i, row := range todoListStrWithPadding {
		strRow := stringUtil.AddPipes(row)
		fmt.Println(strRow)

		// Add header separator after the first row
		if i == 0 {
			fmt.Println(stringUtil.HeaderSeparator(len(strRow)))
		}
	}

}

func main() {
	args := os.Args[1:]

	logger.Debug("args: %v", args)

	var command string
	var index int
	for i, arg := range args {
		commandArg, ok := CommandMap[arg]
		if ok {
			logger.Debug("command: %v", commandArg)
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
		logger.Debug("createArg: %v", createArg)

		// init reader
		reader := csv.NewReader(csvFile)
		reader.FieldsPerRecord = -1 // (?)

		// load todoList from the csv
		todoList, err := reader.ReadAll()
		if err != nil {
			logger.Error("cannot read the csv: ", err)
		}

		// if we got empty, then append headers to the todoList
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
			logger.Error("cannot create todo: ", err)
		}

		// Reset file position and clear existing content before writing
		csvFile.Seek(0, 0)
		csvFile.Truncate(0)

		// write back to csv
		if err := writer.WriteAll(todoList); err != nil {
			logger.Error("cannot write to csv: ", err)
		}

		fmt.Println("todo created")

	case "list":
		// get the follow arguments after the command
		listArg := strings.Join(args[index+1:], " ")
		logger.Debug("listArg: %v", listArg)

		// init reader
		reader := csv.NewReader(csvFile)
		reader.FieldsPerRecord = -1
		logger.Debug("reader: %v", reader)

		// load todoList from the csv
		todoList, err := reader.ReadAll()
		if err != nil {
			logger.Error("cannot read the csv: ", err)
		}
		logger.Debug("todoList: %v", todoList)

		// if we got empty, then append headers to the todoList
		if len(todoList) == 0 {
			headers := []string{"id", "description", "done", "created at", "priority"}
			todoList = append(todoList, headers)
		}

		listTodo(todoList)
	}
}
