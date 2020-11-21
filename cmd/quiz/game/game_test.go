package game_test

import (
	"strings"
	"testing"

	"github.com/antonio-muniz/gophercises/cmd/quiz/game"
	"github.com/stretchr/testify/require"
)

func TestPlay(t *testing.T) {
	writer := &strings.Builder{}
	reader := strings.NewReader(strings.Join([]string{
		"10",
		"10",
		"2",
		"11",
		"3",
	}, "\n"))
	err := game.Play(writer, reader)
	require.NoError(t, err)
	require.Equal(t, strings.Join([]string{
		"Question #1: \"5+5\" Answer: ",
		"Question #2: \"7+3\" Answer: ",
		"Question #3: \"1+1\" Answer: ",
		"Question #4: \"8+3\" Answer: ",
		"Question #5: \"1+2\" Answer: ",
		"",
	}, "\n"), writer.String())
}
