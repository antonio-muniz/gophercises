package quiz

import (
	"github.com/antonio-muniz/gophercises/quiz_game/csv"
)

func LoadProblems(problemFile string) ([]Problem, error) {
	problemRecords, err := csv.Load(problemFile)
	if err != nil {
		return nil, err
	}

	problems := createProblems(problemRecords)

	return problems, nil
}

func createProblems(problemRecords [][]string) []Problem {
	problems := make([]Problem, len(problemRecords))
	for index, problemRecord := range problemRecords {
		question := problemRecord[0]
		correctAnswer := problemRecord[1]
		problems[index] = NewProblem(question, correctAnswer)
	}
	return problems
}
