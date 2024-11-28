package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

var CommandMap = map[string]string{
	"create": "create",
	"list":   "list",
}

const (
	NUM_OF_COLUMNS = 5
	FILE_NAME      = "todo.csv"
)

func OpenCsvFile() *os.File {
	var csvFile *os.File

	csvFile, err := os.Open(FILE_NAME)
	if err != nil {
		log.Fatal("Cannot open file")
	}

	if os.IsExist(err) {
		return csvFile
	}

	csvFile, err = os.Create(FILE_NAME)
	if err != nil {
		log.Fatal("Cannot create file")
	}

	return csvFile
}

func main() {
	args := os.Args[1:]

	var command string
	var index int
	for i, arg := range args {
		commandArg, ok := CommandMap[arg]
		if ok {
			fmt.Println("WE GOT A COMMAND: ", commandArg)
			command = commandArg
			index = i
		}
	}

	csvFile := OpenCsvFile()
	defer csvFile.Close()

	switch command {
	case "create":
		todo := strings.Join(args[index:], " ")
		fmt.Printf("todo: %v\n", todo)

		writer := csv.NewWriter(csvFile)
		defer writer.Flush()

		// TODO: read the file before appending a new record

		// headers := []string{"id", "description", "done", "created at", "priority"}
		// rows := [][]string{}
	}

}
