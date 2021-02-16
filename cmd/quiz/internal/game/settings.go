package game

import (
	"io"
	"time"
)

type Settings struct {
	QuestionReader io.Reader
	TimerDuration  time.Duration
	PlayerInput    io.Reader
	PlayerOutput   io.Writer
}
