package game

import (
	"bufio"
	"fmt"

	"github.com/antonio-muniz/gophercises/pkg/csv"
	"github.com/pkg/errors"
)

func Play(settings Settings) error {
	questionRows, err := csv.Read(settings.QuestionReader)
	if err != nil {
		return err
	}
	questions, err := ParseQuestions(questionRows)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(settings.PlayerInput)
	answers := make([]string, 0, len(questions))
	for _, question := range questions {
		fmt.Fprintf(settings.PlayerOutput, "Question #%d: %s Answer: ", question.Number, question.Text)
		if !scanner.Scan() {
			return errors.WithStack(scanner.Err())
		}
		answer := scanner.Text()
		answers = append(answers, answer)
		fmt.Fprintln(settings.PlayerOutput)
	}

	score, err := CountScore(questions, answers)
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
