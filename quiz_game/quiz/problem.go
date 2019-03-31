package quiz

type Problem struct {
	question      string
	correctAnswer string
}

func NewProblem(question string, correctAnswer string) Problem {
	return Problem{
		question:      question,
		correctAnswer: correctAnswer,
	}
}

func (p Problem) Question() string {
	return p.question
}

func (p Problem) IsCorrect(answer string) bool {
	return answer == p.correctAnswer
}
