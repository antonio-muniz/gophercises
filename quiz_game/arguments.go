package main

import (
	"flag"
	"os"
)

type arguments struct {
	problemFile string
}

func parseArguments() *arguments {
	var showHelp = flag.Bool(
		"help",
		false,
		"Display usage.",
	)

	var problemFile = flag.String(
		"file",
		"problems.csv",
		"The path to the .csv file containing the problems in \"question,answer\" format.",
	)

	flag.Parse()

	if *showHelp {
		flag.Usage()
		os.Exit(0)
	}

	return &arguments{
		problemFile: *problemFile,
	}
}
