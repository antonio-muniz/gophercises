package main

import (
	"flag"
	"path/filepath"

	"github.com/pkg/errors"
)

type arguments struct {
	QuestionFilePath string
}

func parseArguments() (arguments, error) {
	var filePath string
	flag.StringVar(
		&filePath,
		"f",
		"./test/data/cmd/quiz/questions.csv",
		"[required] Path to the questions file (.csv)",
	)
	flag.Parse()
	absoluteFilePath, err := filepath.Abs(filePath)
	if err != nil {
		return arguments{}, errors.WithStack(err)
	}
	arguments := arguments{
		QuestionFilePath: absoluteFilePath,
	}
	return arguments, nil
}
