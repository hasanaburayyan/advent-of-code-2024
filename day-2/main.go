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

	fmt.Printf("Part one answer: %d\n", partOneSolution(input, 0))
	fmt.Printf("Part two answer: %d\n", partOneSolution(input, 1))
}

type OrderFunc func([]int, int, int) bool

func partOneSolution(input ProblemInput, faultTolerance int) int {

	var orderCheck OrderFunc

	sum := 0

	for _, row := range input.Matrix {
		if isAscending(row, 0, 1) {
			orderCheck = isAscending
		} else {
			orderCheck = isDescending
		}

		safe := true
		rowFaultTolerance := faultTolerance
		for i := 0; i < len(row)-1; i++ {
			if !isSafe(row, i, i+1, orderCheck) {
				// use fault tolerance if any
				if rowFaultTolerance > 0 {
					rowFaultTolerance--
					continue
				}
				safe = false
				break
			}
		}

		// check last elements
		if !orderCheck(row, len(row)-2, len(row)-1) {
			if rowFaultTolerance > 0 {
				safe = true
			} else {
				safe = false
			}
		}

		if safe {
			sum++
		}
	}

	return sum
}

// func checkRow(row []int, checkOrder OrderFunc) bool {
// 	if !checkOrder(row[0], row[1]) {
// 		return false
// 	}
// }

func isSafe(row []int, i int, j int, orderCheck OrderFunc) bool {
	if !orderCheck(row, i, j) {
		return false
	}

	diff := math.Abs(float64(row[i] - row[j]))

	return diff <= 3 && diff > 0
}

func isAscending(row []int, i int, j int) bool {
	if j >= len(row) {
		return true
	}

	if row[i] <= row[j] {
		return true
	}

	return false
}

func isDescending(row []int, i int, j int) bool {
	if j >= len(row) {
		return true
	}

	return !isAscending(row, i, j)
}
