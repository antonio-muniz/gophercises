package game_test

import (
	"testing"

	"github.com/antonio-muniz/gophercises/cmd/quiz/internal/game"
	"github.com/stretchr/testify/require"
)

func TestComputeScore(t *testing.T) {
	t.Run("counts correct answers", func(t *testing.T) {
		questions := []game.Question{
			{CorrectAnswer: "A"},
			{CorrectAnswer: "B"},
			{CorrectAnswer: "C"},
			{CorrectAnswer: "D"},
		}
		answers := []string{"A", "B", "X", "D"}

		score, err := game.ComputeScore(questions, answers)

		require.NoError(t, err)
		require.Equal(t, 3, score)
	})

	t.Run("fails if there are more answers than questions", func(t *testing.T) {
		questions := []game.Question{
			{CorrectAnswer: "A"},
			{CorrectAnswer: "B"},
			{CorrectAnswer: "C"},
			{CorrectAnswer: "D"},
		}
		answers := []string{"A", "B", "C", "D", "E"}

		_, err := game.ComputeScore(questions, answers)

		require.EqualError(t, err, "there are more answers than questions")
	})

	t.Run("counts in order if there are less answers than questions", func(t *testing.T) {
		questions := []game.Question{
			{CorrectAnswer: "A"},
			{CorrectAnswer: "B"},
			{CorrectAnswer: "C"},
			{CorrectAnswer: "D"},
		}
		answers := []string{"A", "X", "C"}

		score, err := game.ComputeScore(questions, answers)
		require.NoError(t, err)
		require.Equal(t, 2, score)
	})
}
