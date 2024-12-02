package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type ProblemInput struct {
	LeftList  []int
	RightList []int
}

func ReadInput() (ProblemInput, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return ProblemInput{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	leftList := []int{}
	rightList := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "   ")
		leftValue, err := strconv.Atoi(split[0])
		if err != nil {
			return ProblemInput{}, err
		}
		rightValue, err := strconv.Atoi(split[1])
		if err != nil {
			return ProblemInput{}, err
		}
		leftList = append(leftList, leftValue)
		rightList = append(rightList, rightValue)
	}

	return ProblemInput{LeftList: leftList, RightList: rightList}, nil
}
