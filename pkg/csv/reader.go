package csv

import (
	"encoding/csv"
	"io"

	"github.com/pkg/errors"
)

func Read(reader io.Reader) ([][]string, error) {
	csvReader := csv.NewReader(reader)
	rows, err := csvReader.ReadAll()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return rows, nil
}
