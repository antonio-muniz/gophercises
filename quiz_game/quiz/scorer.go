package quiz

func ComputeScore(problems []Problem, answers []string) int {
	score := 0
	for index, problem := range problems {
		answer := answers[index]
		if problem.IsCorrect(answer) {
			score++
		}
	}
	return score
}
