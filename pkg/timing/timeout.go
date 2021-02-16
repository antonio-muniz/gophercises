package timing

import (
	"errors"
	"time"
)

var ErrTimeout = errors.New("timeout")

func WithTimeout(timeoutMilliseconds int, function func(chan error)) error {
	finish := make(chan error)
	go function(finish)
	if timeoutMilliseconds == 0 {
		return <-finish
	} else {
		timeout := time.After(time.Duration(timeoutMilliseconds) * time.Millisecond)
		select {
		case err := <-finish:
			return err
		case <-timeout:
			return ErrTimeout
		}
	}
}
