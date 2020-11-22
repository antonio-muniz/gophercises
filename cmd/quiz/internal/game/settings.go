package game

import "io"

type Settings struct {
	QuestionReader io.Reader
	PlayerInput    io.Reader
	PlayerOutput   io.Writer
}
