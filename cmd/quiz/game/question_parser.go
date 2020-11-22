package game

import "fmt"

func ParseQuestions(questionRows [][]string) ([]Question, error) {
	questions := make([]Question, 0, len(questionRows))
	for index, row := range questionRows {
		if len(row) < 2 {
			return nil, fmt.Errorf("too few columns at question row %v", index)
		}
		questions = append(questions, Question{
			Number:        index + 1,
			Text:          row[0],
			CorrectAnswer: row[1],
		})
	}
	return questions, nil
}
