package game_test

import (
	"testing"

	"github.com/antonio-muniz/gophercises/cmd/quiz/game"
	"github.com/stretchr/testify/require"
)

func TestParseQuestions(t *testing.T) {
	t.Run("converts all rows into questions", func(t *testing.T) {
		questionRows := [][]string{
			{"How old are you?", "26"},
			{"What are you doing?", "Coding"},
		}

		questions, err := game.ParseQuestions(questionRows)

		require.NoError(t, err)
		require.Equal(t, []game.Question{
			{
				Number:        1,
				Text:          "How old are you?",
				CorrectAnswer: "26",
			},
			{
				Number:        2,
				Text:          "What are you doing?",
				CorrectAnswer: "Coding",
			},
		}, questions)
	})

	t.Run("fails to convert rows with less than 2 columns", func(t *testing.T) {
		questionRows := [][]string{
			{"How old are you?", "26"},
			{"What are you doing?"},
		}

		_, err := game.ParseQuestions(questionRows)
		require.EqualError(t, err, "too few columns at question row 1")
	})
}
