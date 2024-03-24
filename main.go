package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
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
		}
	}
	// csvFilePath := "problems.csv"
	quizList, err := parseCsvFile(csvFilePath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Starting the Game 🔥🔥🔥: ")
	correctAnswers := 0
	totalQuizAsked := 0
	var userInput string

	for _, quiz := range(quizList) {
		fmt.Printf("%v :", quiz.question)
		totalQuizAsked++
		fmt.Scanln(&userInput)
		solution := strings.TrimSpace(userInput)
		if solution == quiz.answer {
			correctAnswers++
			fmt.Println("Correct!!")
		} else {
			fmt.Printf("Wrong!! the answer is %v \n", quiz.answer)
		}
	}
}

type Quiz struct {
	question string
	answer string
}

func parseCsvFile(fileName string) ([]Quiz, error){
	csvFile, err := os.Open(fileName)
	defer csvFile.Close()
	if err != nil {
		return make([]Quiz, 1), errors.New("Unable to open file")
	}
	csvReader := csv.NewReader(csvFile)
	result := make([]Quiz, 0, 50)

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err) // will not add it to the list
			continue
		}
		newQuiz := Quiz{question: record[0], answer: record[1]}
		result = append(result, newQuiz)
	}

	return result, nil
}
