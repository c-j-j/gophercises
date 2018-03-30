package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	problemCsvFilename := flag.String("problems", "problems.csv", "Problem CSV File")
	timeLimit := flag.Int("time-limit", 5, "Number of seconds before quiz ends")
	flag.Parse()

	problems := readProblemsFromFile(*problemCsvFilename)

	correctAnswers := playQuiz(problems, *timeLimit)
	printResults(correctAnswers, problems)
}

func playQuiz(problems []Problem, timeLimit int) int {
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
	answerChannel := make(chan string)

	correctAnswers := 0
	for i, problem := range problems {
		fmt.Printf("%d. What is %s?\n", i+1, problem.question)
		go readAnswerFromUserfuncName(answerChannel)

		select {
		case <-timer.C:
			fmt.Println("Time's up")
			return correctAnswers
		case answer := <-answerChannel:
			if problem.answer == answer {
				fmt.Printf("Correct\n")
				correctAnswers++
			} else {
				fmt.Printf("Wrong, correct answer is %s\n", problem.answer)
			}
		}
	}
	return correctAnswers
}

func readProblemsFromFile(problemCsvFilename string) []Problem {
	file, err := os.Open(problemCsvFilename)
	if err != nil {
		exit(problemCsvFilename)
	}
	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		exit("Unable to read CSV file")
	}
	problems := mapLinesToProblems(lines)
	return problems
}

func readAnswerFromUserfuncName(answerChannel chan string) {
	var answer string
	fmt.Scanf("%s", &answer)
	answerChannel <- answer
}

func printResults(correctAnswers int, problems []Problem) {
	fmt.Printf("Number of correct answers: %d out of %d", correctAnswers, len(problems))
}

type Problem struct {
	question string
	answer   string
}

func mapLinesToProblems(lines [][]string) []Problem {
	problems := make([]Problem, len(lines))

	for index, line := range lines {
		problems[index] = NewProblem(line[0], strings.TrimSpace(line[1]))
	}
	return problems
}

func NewProblem(question string, answer string) Problem {
	return Problem{question, answer}
}

func exit(problemCsvFilename string) {
	fmt.Printf("Unable to open CSV file: %s\n", problemCsvFilename)
	os.Exit(1)
}
