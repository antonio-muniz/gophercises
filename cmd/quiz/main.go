package main

import (
	"fmt"
	"os"

	"github.com/antonio-muniz/gophercises/cmd/quiz/internal/game"
)

func main() {
	arguments, err := parseArguments()
	if err != nil {
		os.Exit(1)
	}
	questionFile, err := os.Open(arguments.QuestionFilePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer questionFile.Close()
	settings := game.Settings{
		QuestionReader: questionFile,
		TimerDuration:  arguments.TimerDuration,
		PlayerInput:    os.Stdin,
		PlayerOutput:   os.Stdout,
	}
	err = game.Play(settings)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
