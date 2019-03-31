package main

import (
	"flag"
	"os"
)

type arguments struct {
	problemFile string
	timeLimit   int
}

func parseArguments() *arguments {
	showHelp := flag.Bool(
		"help",
		false,
		"Display usage.",
	)
	problemFile := flag.String(
		"file",
		"problems.csv",
		"The path to the .csv file containing the problems in \"question,answer\" format.",
	)
	timeLimit := flag.Int(
		"time",
		10,
		"The time limit of the quiz, in seconds.",
	)

	flag.Parse()

	if *showHelp {
		flag.Usage()
		os.Exit(0)
	}

	return &arguments{
		problemFile: *problemFile,
		timeLimit:   *timeLimit,
	}
}
