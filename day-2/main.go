package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	input, err := ReadInput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part one answer: %d\n", solve(input, false))
	fmt.Printf("Part two answer: %d\n", solve(input, true))
}

func solve(input ProblemInput, faultTolerable bool) int {
	sum := 0

	for _, row := range input.Matrix {
		if isRowSafe(row, row[0] < row[1], faultTolerable) {
			sum++
		}
	}

	return sum
}

func isRowSafe(row []int, ascending bool, canFail bool) bool {
	if len(row) <= 1 {
		return true
	}

	diff := math.Abs(float64(row[0] - row[1]))
	correctOrder := (row[0] < row[1]) == ascending

	if diff > 3 || diff == 0 || !correctOrder {
		if canFail {
			tmpRow := make([]int, len(row))
			copy(tmpRow, row)
			tmpRow[1] = tmpRow[0]
			return isRowSafe(row[1:], ascending, false) || isRowSafe(tmpRow[1:], ascending, false)
		}
		return false
	}

	return isRowSafe(row[1:], ascending, canFail)
}
