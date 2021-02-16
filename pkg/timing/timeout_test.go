package timing_test

import (
	"errors"
	"testing"
	"time"

	"github.com/antonio-muniz/gophercises/pkg/timing"
	"github.com/stretchr/testify/require"
)

func TestWithTimeout(t *testing.T) {
	t.Run("function finishes successfully before timeout", func(t *testing.T) {
		var result string
		err := timing.WithTimeout(200*time.Millisecond, func(finish chan error) {
			time.Sleep(100 * time.Millisecond)
			result = "done"
			finish <- nil
		})
		require.NoError(t, err)
		require.Equal(t, "done", result)
	})

	t.Run("function finishes with error before timeout", func(t *testing.T) {
		var result string
		err := timing.WithTimeout(200*time.Millisecond, func(finish chan error) {
			time.Sleep(100 * time.Millisecond)
			finish <- errors.New("oops!")
		})
		require.EqualError(t, err, "oops!")
		require.Equal(t, "", result)
	})

	t.Run("function times out", func(t *testing.T) {
		var result string
		err := timing.WithTimeout(200*time.Millisecond, func(finish chan error) {
			time.Sleep(300 * time.Millisecond)
			result = "done"
			finish <- nil
		})
		require.Equal(t, timing.ErrTimeout, err)
		require.Equal(t, "", result)
	})
}
