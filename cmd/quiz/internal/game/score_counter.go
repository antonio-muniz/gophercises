package game

import "errors"

func CountScore(questions []Question, answers []string) (int, error) {
	if len(answers) > len(questions) {
		return -1, errors.New("there are more answers than questions")
	}
	if len(answers) < len(questions) {
		return -1, errors.New("there are less answers than questions")
	}
	var score int
	for index, question := range questions {
		answer := answers[index]
		if answer == question.CorrectAnswer {
			score++
		}
	}
	return score, nil
}
