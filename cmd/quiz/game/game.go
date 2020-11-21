package game

import (
	"bufio"
	"fmt"
	"io"

	"github.com/antonio-muniz/gophercises/pkg/csv"
	"github.com/pkg/errors"
)

func Play(writer io.Writer, reader io.Reader) error {
	rows, err := csv.Read("C:\\Users\\ahnm_\\go\\src\\github.com\\antonio-muniz\\gophercises\\test\\data\\cmd\\quiz\\questions.csv")
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(reader)
	for index, row := range rows {
		fmt.Fprintf(writer, "Question #%d: \"%s\" Answer: ", index+1, row[0])
		if !scanner.Scan() {
			return errors.WithStack(scanner.Err())
		}
		fmt.Fprintln(writer)
	}
	return nil
}
