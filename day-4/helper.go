package main

import (
	"bytes"
	"os"
)

type ProblemInput struct {
	Puzzle [][]byte
}

func ReadInput() (ProblemInput, error) {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		return ProblemInput{}, err
	}

	lines := bytes.Split(file, []byte("\n"))

	matrix := make([][]byte, 0, len(lines))
	matrix = append(matrix, lines...)

	return ProblemInput{Puzzle: matrix}, nil
}
