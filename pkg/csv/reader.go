package csv

import (
	"encoding/csv"
	"os"

	"github.com/pkg/errors"
)

func Read(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	csvReader := csv.NewReader(file)
	rows, err := csvReader.ReadAll()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return rows, nil
}
