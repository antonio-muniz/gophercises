package game

import (
	"bufio"
	"fmt"

	"github.com/antonio-muniz/gophercises/pkg/csv"
	"github.com/antonio-muniz/gophercises/pkg/timing"
	"github.com/pkg/errors"
)

func Play(settings Settings) error {
	questions, err := loadQuestions(settings)
	if err != nil {
		return err
	}
	answers, err := collectAnswers(questions, settings)
	if err != nil {
		return err
	}
	return displayScore(questions, answers, settings)
}

func loadQuestions(settings Settings) ([]Question, error) {
	questionRows, err := csv.Read(settings.QuestionReader)
	if err != nil {
		return nil, err
	}
	questions, err := ParseQuestions(questionRows)
	if err != nil {
		return nil, err
	}
	return questions, nil
}

func collectAnswers(questions []Question, settings Settings) ([]string, error) {
	scanner := bufio.NewScanner(settings.PlayerInput)
	answers := make([]string, 0, len(questions))

	err := timing.WithTimeout(settings.TimerDuration, func(finish chan error) {
		for _, question := range questions {
			fmt.Fprintf(settings.PlayerOutput, "Question #%d: %s Answer: ", question.Number, question.Text)
			if !scanner.Scan() {
				finish <- errors.WithStack(scanner.Err())
				return
			}
			answer := scanner.Text()
			answers = append(answers, answer)
			fmt.Fprintln(settings.PlayerOutput)
		}
		finish <- nil
	})

	switch err {
	case nil:
		return answers, nil
	case timing.ErrTimeout:
		fmt.Fprintln(settings.PlayerOutput)
		fmt.Fprintln(settings.PlayerOutput, "Time is up!")
		return answers, nil
	default:
		return nil, err
	}
}

func displayScore(questions []Question, answers []string, settings Settings) error {
	score, err := ComputeScore(questions, answers)
	if err != nil {
		return err
	}
	if score == len(questions) {
		fmt.Fprintf(settings.PlayerOutput, "You scored %v/%v! Congratulations!\n", score, score)
	} else {
		fmt.Fprintf(settings.PlayerOutput, "You scored %v/%v!\n", score, len(questions))
	}
	return nil
}
