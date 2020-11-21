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
	scanner := bufio.NewScanner(settings.PlayerInput)
	for index, row := range questionRows {
		fmt.Fprintf(settings.PlayerOutput, "Question #%d: %s Answer: ", index+1, row[0])
		if !scanner.Scan() {
			return errors.WithStack(scanner.Err())
		}
		fmt.Fprintln(settings.PlayerOutput)
	}
	return nil
}
