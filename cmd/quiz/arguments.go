package main

import (
	"flag"
	"path/filepath"
	"time"

	"github.com/pkg/errors"
)

type arguments struct {
	QuestionFilePath string
	TimerDuration    time.Duration
}

func parseArguments() (arguments, error) {
	var filePath string
	var timerSeconds int
	flag.StringVar(
		&filePath,
		"file",
		"./test/data/cmd/quiz/questions.csv",
		"[required] Path to the questions file (.csv)",
	)
	flag.IntVar(
		&timerSeconds,
		"time",
		0,
		"[optional] Quiz timer duration in seconds (no timer if zero or unspecified)",
	)
	flag.Parse()
	absoluteFilePath, err := filepath.Abs(filePath)
	if err != nil {
		return arguments{}, errors.WithStack(err)
	}
	timerDuration := time.Duration(timerSeconds) * time.Second
	arguments := arguments{
		QuestionFilePath: absoluteFilePath,
		TimerDuration:    timerDuration,
	}
	return arguments, nil
}
