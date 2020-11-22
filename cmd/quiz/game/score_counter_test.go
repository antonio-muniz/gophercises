package game_test

import (
	"testing"

	"github.com/antonio-muniz/gophercises/cmd/quiz/game"
	"github.com/stretchr/testify/require"
)

func TestCountScore(t *testing.T) {
	t.Run("counts correct answers", func(t *testing.T) {
		questions := []game.Question{
			{CorrectAnswer: "A"},
			{CorrectAnswer: "B"},
			{CorrectAnswer: "C"},
			{CorrectAnswer: "D"},
		}
		answers := []string{"A", "B", "X", "D"}

		score, err := game.CountScore(questions, answers)

		require.NoError(t, err)
		require.Equal(t, 3, score)
	})

	t.Run("fails if there are too many answers", func(t *testing.T) {
		questions := []game.Question{
			{CorrectAnswer: "A"},
			{CorrectAnswer: "B"},
			{CorrectAnswer: "C"},
			{CorrectAnswer: "D"},
		}
		answers := []string{"A", "B", "C", "D", "E"}

		_, err := game.CountScore(questions, answers)

		require.EqualError(t, err, "there are more answers than questions")
	})

	t.Run("fails if there are too few answers", func(t *testing.T) {
		questions := []game.Question{
			{CorrectAnswer: "A"},
			{CorrectAnswer: "B"},
			{CorrectAnswer: "C"},
			{CorrectAnswer: "D"},
		}
		answers := []string{"A", "B", "C"}

		_, err := game.CountScore(questions, answers)

		require.EqualError(t, err, "there are less answers than questions")
	})
}
