package game

import "io"

type Settings struct {
	QuestionReader    io.Reader
	TimerMilliseconds int
	PlayerInput       io.Reader
	PlayerOutput      io.Writer
}
