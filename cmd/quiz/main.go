package main

import (
	"os"

	"github.com/antonio-muniz/gophercises/cmd/quiz/internal/game"
	"github.com/pkg/errors"
)

func main() {
	questionFile, err := os.Open("C:\\Users\\ahnm_\\go\\src\\github.com\\antonio-muniz\\gophercises\\test\\data\\cmd\\quiz\\questions.csv")
	if err != nil {
		panic(errors.WithStack(err))
	}
	defer questionFile.Close()
	settings := game.Settings{
		QuestionReader: questionFile,
		PlayerInput:    os.Stdin,
		PlayerOutput:   os.Stdout,
	}
	err = game.Play(settings)
	if err != nil {
		panic(err)
	}
}
