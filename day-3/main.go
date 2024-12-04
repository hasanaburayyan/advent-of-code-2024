package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, err := ReadInput()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("problem one: %d\n", solveProblemOne(input.Output))
	fmt.Printf("problem one: %d\n", solveProblemTwo(input.Output))
}

func solveProblemOne(input string) int {
	re := regexp.MustCompile(`mul\(((-)?[0-9]+\,((-)?[0-9]+))\)`)
	matches := re.FindAll([]byte(input), -1)

	sum := 0

	for _, match := range matches {
		numString := strings.Split(string(match[4:len(match)-1]), ",")

		left, _ := strconv.Atoi(numString[0])
		right, _ := strconv.Atoi(numString[1])
		sum += right * left
	}

	return sum
}

func solveProblemTwo(input string) int {
	re := regexp.MustCompile(`(mul\((-?\d+),(-?\d+)\)|do\(\)|don't\(\))`)
	matches := re.FindAllStringSubmatch(input, -1)

	doEnabled := true
	sum := 0

	for _, match := range matches {
		if strings.HasPrefix(match[0], "do()") {
			doEnabled = true
		} else if strings.HasPrefix(match[0], "don't()") {
			doEnabled = false
		} else if strings.HasPrefix(match[0], "mul") && doEnabled {
			// Extract numbers from the mul instruction
			left, _ := strconv.Atoi(match[2])
			right, _ := strconv.Atoi(match[3])
			sum += left * right
		}
	}

	return sum
}
