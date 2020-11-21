package main

import (
	"os"

	"github.com/antonio-muniz/gophercises/cmd/quiz/game"
)

func main() {
	game.Play(os.Stdout, os.Stdin)
}
