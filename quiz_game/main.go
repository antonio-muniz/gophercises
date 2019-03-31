package main

import (
	"fmt"
	"os"

	"github.com/antonio-muniz/gophercises/quiz_game/quiz"
)

func main() {
	args := parseArguments()

	quiz.PresentWelcomeScreen()

	problems, err := quiz.LoadProblems(args.problemFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading problems: %s\n", err)
		os.Exit(1)
	}

	answers := quiz.PresentProblems(problems, args.timeLimit)

	score := quiz.ComputeScore(problems, answers)

	quiz.PresentScore(score, len(problems))
}
