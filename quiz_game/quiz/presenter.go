package quiz

import (
	"fmt"
	"time"

	"github.com/antonio-muniz/gophercises/quiz_game/stdio"
)

func PresentWelcomeScreen() {
	stdio.UpdateText("\nWelcome to the Quiz Game!\n")
	stdio.AwaitConfirmation()
}

func PresentProblems(problems []Problem, timeLimitSeconds int) []string {
	answers := make([]string, len(problems))

	timeLimit := time.Duration(timeLimitSeconds)
	timer := time.After(timeLimit * time.Second)

	func() {
		for index, problem := range problems {
			select {
			case <-timer:
				stdio.UpdateText("\nTime is up!\n")
				stdio.AwaitConfirmation()
				return
			case answer := <-presentProblem(index+1, problem.Question()):
				answers[index] = answer
			}
		}
	}()

	return answers
}

func PresentScore(score int, maxScore int) {
	result := fmt.Sprintf("\nYour score: %d/%d.\nCongratulations!\n", score, maxScore)
	stdio.UpdateText(result)
	stdio.AwaitConfirmation()
}

func presentProblem(number int, question string) chan string {
	prompt := fmt.Sprintf("\nProblem #%d: %s\nAnswer: ", number, question)
	stdio.UpdateText(prompt)
	answerIn := make(chan string)
	go func() {
		answerIn <- stdio.RequestWord()
	}()
	return answerIn
}
