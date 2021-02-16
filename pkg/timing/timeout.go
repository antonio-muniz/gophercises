package timing

import (
	"errors"
	"time"
)

var ErrTimeout = errors.New("timeout")

func WithTimeout(timeoutDuration time.Duration, function func(chan error)) error {
	finish := make(chan error)
	go function(finish)
	if timeoutDuration == 0 {
		return <-finish
	} else {
		select {
		case err := <-finish:
			return err
		case <-time.After(timeoutDuration):
			return ErrTimeout
		}
	}
}
