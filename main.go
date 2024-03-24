package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func boundCheck(i int, arr []string) bool {
	if i >= len(arr) {
		return false
	}
	return true
}

func checkValidCSVFile(csvFileName string) (string, error) {
	if !strings.HasSuffix(csvFileName, ".csv") {
		return "", errors.New("file is not a CSV file")
	}
	if _, err := os.Stat(csvFileName); os.IsNotExist(err) {
		return "", errors.New("file does not exist")
	}

	return csvFileName, nil
}

func main() {
	if len(os.Args) < 3 {
		log.Fatal(errors.New("Please provide csv file name. usage: -f <fileName.csv>"))
	}

	var csvFilePath string
	progArgs := os.Args[1:]
	for i := 0; i < len(progArgs); i++ {
		switch progArgs[i] {
		case "-f":
			if !boundCheck(i+1, progArgs) {
				log.Fatal(errors.New("Please provide csv file name. usage: -f <fileName.csv>"))
			}
			fileName, err := checkValidCSVFile(progArgs[i+1])
			if err != nil {
				log.Fatal(err)
			}
			csvFilePath = fileName
		default:
			fmt.Println("Let's play the game")
		}
	}
	fmt.Println(csvFilePath)
}