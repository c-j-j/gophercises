package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	problemCsvFilename := flag.String("problems", "problems.csv", "Problem CSV File")
	flag.Parse()

	file, err := os.Open(*problemCsvFilename)

	if err != nil {
		exit(*problemCsvFilename)
	}

	lines, err := csv.NewReader(file).ReadAll()

	if err != nil {
		exit("Unable to read CSV file")
	}

	problems := mapLinesToProblems(lines)

	correctAnswers := 0

	for i, problem := range problems {
		fmt.Printf("%d. What is %s?\n", i+1, problem.question)
		var answer string
		fmt.Scanf("%s", &answer)

		if problem.answer == answer {
			fmt.Printf("Correct\n")
			correctAnswers++
		} else {
			fmt.Printf("Wrong, correct answer is %s\n", problem.answer)
		}
	}
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
