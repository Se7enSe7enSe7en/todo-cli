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
	FILE_PATH      = "./todo.csv"
)

func OpenCsvFile() *os.File {
	var csvFile *os.File

	csvFile, errOpen := os.Open(FILE_PATH)
	if errOpen != nil {
		// if file exists return file
		if os.IsExist(errOpen) {
			return csvFile
		}

		// if file does not exist, create file
		if os.IsNotExist(errOpen) {
			csvFile, errCreate := os.Create(FILE_PATH)
			if errCreate != nil {
				log.Fatal("Cannot create file", errCreate)
			}
			return csvFile
		}

		// if it has an error not related to existence, terminate
		log.Fatal("Cannot open file: ", errOpen)
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
		createArg := strings.Join(args[index+1:], " ")
		fmt.Printf("createArg: %v\n", createArg)

		writer := csv.NewWriter(csvFile)
		defer writer.Flush()

		// TODO: read the file before appending a new record

		// headers := []string{"id", "description", "done", "created at", "priority"}
		// rows := [][]string{}
	}

}
