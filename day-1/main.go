package main

import (
	"fmt"
	"log"
	"math"
	"sort"
)

func main() {
	input, err := ReadInput()
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	fmt.Printf("Part one: %d\n", partOneSimpleApproach(input))
	fmt.Printf("Part two: %d\n", partTwoSimpleApproach(input))
}

func partOneSimpleApproach(input ProblemInput) int {
	// sort each list
	sort.Ints(input.LeftList)
	sort.Ints(input.RightList)

	// Now O(n) iteration
	sum := 0.0

	for i := 0; i < len(input.LeftList); i++ {
		diff := input.LeftList[i] - input.RightList[i]
		sum += math.Abs(float64(diff))
	}

	return int(sum)
}

func partTwoSimpleApproach(input ProblemInput) int {
	// Create a frequency map of the right list
	freqMap := make(map[int]int)

	for _, val := range input.RightList {
		if existingCount, ok := freqMap[val]; ok {
			freqMap[val] = existingCount + 1
			continue
		}
		freqMap[val] = 1
	}

	// now iterate left list checking freq map
	sum := 0
	for _, val := range input.LeftList {
		if count, ok := freqMap[val]; ok {
			sum += count * val
		}
	}

	return sum
}
