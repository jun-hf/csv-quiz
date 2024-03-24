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
	"time"
)

func checkValidCSVFile(csvFileName string) (string, error) {
	if !strings.HasSuffix(csvFileName, ".csv") {
		return "", errors.New("file is not a CSV file")
	}
	if _, err := os.Stat(csvFileName); os.IsNotExist(err) {
		return "", errors.New("file does not exist")
	}

	return csvFileName, nil
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

func main() {
	csvFilePath := flag.String("file", "problems.csv", "A csv filePath")
	quizDuration := flag.Int("seconds", 100, "Time limit on the quiz")
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
	timeOutQuiz := time.NewTimer(time.Duration(*quizDuration) * time.Second)

	fmt.Println("Starting the Game 🔥🔥🔥: ")
	correctAnswers := 0
	totalQuizAsked := 0
	var userInput string

	QuizLoop:
	for _, quiz := range(quizList) {
		fmt.Printf("%v :", quiz.question)
		answerCh := make(chan string)
		go func() {
			fmt.Scan(&userInput)
			answerCh <- strings.TrimSpace(userInput)
		}()
		select {
		case <-timeOutQuiz.C:
			break QuizLoop
		case solution := <- answerCh:
			totalQuizAsked++
			if solution == quiz.answer {
				correctAnswers++
				fmt.Println("Correct!!")
			} else {
				fmt.Printf("Wrong!! the answer is %v \n", quiz.answer)
			}
		}
	}
	fmt.Printf("\nCorrect answers: %v, Total quizs attempted %v\n", correctAnswers, totalQuizAsked)
}
