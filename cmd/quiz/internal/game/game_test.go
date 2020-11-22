package game_test

import (
	"strings"
	"testing"

	"github.com/antonio-muniz/gophercises/cmd/quiz/internal/game"
	"github.com/stretchr/testify/require"
)

func TestPlay(t *testing.T) {
	t.Run("player gets all questions right", func(t *testing.T) {
		output := &strings.Builder{}
		settings := game.Settings{
			QuestionReader: strings.NewReader(strings.Join([]string{
				"5 + 5 = ?,10",
				"7 + 3 = ?,10",
				"1 + 1 = ?,2",
				"8 + 3 = ?,11",
				"1 + 2 = ?,3",
				"",
			}, "\n")),
			PlayerInput: strings.NewReader(strings.Join([]string{
				"10",
				"10",
				"2",
				"11",
				"3",
			}, "\n")),
			PlayerOutput: output,
		}
		err := game.Play(settings)
		require.NoError(t, err)
		require.Equal(t, strings.Join([]string{
			"Question #1: 5 + 5 = ? Answer: ",
			"Question #2: 7 + 3 = ? Answer: ",
			"Question #3: 1 + 1 = ? Answer: ",
			"Question #4: 8 + 3 = ? Answer: ",
			"Question #5: 1 + 2 = ? Answer: ",
			"You scored 5/5! Congratulations!",
			"",
		}, "\n"), output.String())
	})

	t.Run("player gets a question wrong", func(t *testing.T) {
		output := &strings.Builder{}
		settings := game.Settings{
			QuestionReader: strings.NewReader(strings.Join([]string{
				"5 + 5 = ?,10",
				"7 + 3 = ?,10",
				"1 + 1 = ?,2",
				"8 + 3 = ?,11",
				"1 + 2 = ?,3",
				"",
			}, "\n")),
			PlayerInput: strings.NewReader(strings.Join([]string{
				"10",
				"10",
				"11",
				"11",
				"3",
			}, "\n")),
			PlayerOutput: output,
		}
		err := game.Play(settings)
		require.NoError(t, err)
		require.Equal(t, strings.Join([]string{
			"Question #1: 5 + 5 = ? Answer: ",
			"Question #2: 7 + 3 = ? Answer: ",
			"Question #3: 1 + 1 = ? Answer: ",
			"Question #4: 8 + 3 = ? Answer: ",
			"Question #5: 1 + 2 = ? Answer: ",
			"You scored 4/5!",
			"",
		}, "\n"), output.String())
	})
}
