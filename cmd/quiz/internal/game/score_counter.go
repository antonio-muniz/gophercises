package game

import "errors"

func ComputeScore(questions []Question, answers []string) (int, error) {
	if len(answers) > len(questions) {
		return -1, errors.New("there are more answers than questions")
	}
	var score int
	for index, answer := range answers {
		question := questions[index]
		if answer == question.CorrectAnswer {
			score++
		}
	}
	return score, nil
}
