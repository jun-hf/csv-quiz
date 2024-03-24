package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func checkValidCSVFile(csvFileName string) (string, error) {
	if csvFileName == "" {
		return "", errors.New("Please provide file a csv file path usage:file=<csv file path>")
	}
	if !strings.HasSuffix(csvFileName, ".csv") {
		return "", errors.New("file is not a CSV file")
	}
	if _, err := os.Stat(csvFileName); os.IsNotExist(err) {
		return "", errors.New("file does not exist")
	}

	return csvFileName, nil
}

func main() {
	csvFilePath := flag.String("file", "", "A filePath")
	flag.Parse()
	csvFile, err := checkValidCSVFile(*csvFilePath)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	quizList, err := parseCsvFile(csvFile)
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
