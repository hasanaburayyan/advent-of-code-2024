package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type ProblemInput struct {
	Matrix [][]int
}

func ReadInput() (ProblemInput, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return ProblemInput{}, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	input := ProblemInput{}

	for scanner.Scan() {
		line := scanner.Text()

		vals := strings.Split(line, " ")

		row := []int{}

		for _, val := range vals {
			n, err := strconv.Atoi(val)
			if err != nil {
				return ProblemInput{}, nil
			}

			row = append(row, n)
		}

		input.Matrix = append(input.Matrix, row)
	}

	return input, nil
}
