package main

import (
	"os"
)

type ProblemInput struct {
	Output string
}

func ReadInput() (ProblemInput, error) {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		return ProblemInput{}, err
	}

	return ProblemInput{
		Output: string(file),
	}, nil
}
