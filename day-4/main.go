package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	input, err := ReadInput()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Part one solution: %d\n", solvePartOne(input.Puzzle))
	fmt.Printf("Part two solution: %d\n", solvePartTwo(input.Puzzle))
}

// solvePartOne counts the occurrences of "XMAS" in all directions
func solvePartOne(grid [][]byte) int {
	word := "XMAS"
	count := 0

	directions := []Direction{
		{1, 0},   // Right
		{-1, 0},  // Left
		{0, 1},   // Down
		{0, -1},  // Up
		{1, 1},   // Diagonal Down-Right
		{-1, -1}, // Diagonal Up-Left
		{1, -1},  // Diagonal Up-Right
		{-1, 1},  // Diagonal Down-Left
	}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == word[0] {
				count += countWord(grid, x, y, word, directions)
			}
		}
	}
	return count
}

func solvePartTwo(grid [][]byte) int {
	word := "MAS"
	count := 0

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == word[1] {
				if checkForCross(grid, x, y, word) {
					count++
				}
			}
		}
	}

	return count
}

type Direction struct {
	dx int
	dy int
}

func countWord(grid [][]byte, x, y int, word string, directions []Direction) int {
	count := 0
	for _, dir := range directions {
		if checkForWord(grid, x, y, word, dir) {
			count++
		}
	}
	return count
}

func checkForWord(grid [][]byte, x int, y int, word string, dir Direction) bool {
	for i := 0; i < len(word); i++ {
		xOffSet := dir.dx * i
		yOffSet := dir.dy * i
		b, err := getPoint(grid, x+xOffSet, y+yOffSet)
		if err != nil || b != word[i] {
			return false
		}
	}
	return true
}

// getPoint retrieves the character at (x, y) or returns an error if out of bounds
func getPoint(grid [][]byte, x, y int) (byte, error) {
	if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[y]) {
		return 0, errors.New("out of bounds")
	}
	return grid[y][x], nil
}

func checkForCross(grid [][]byte, x, y int, word string) bool {
	// use check for word
	// x and y signify the center point

	// check left cross
	directions := []Direction{
		{1, 1},   // Diagonal Down-Right
		{-1, -1}, // Diagonal Up-Left
		{1, -1},  // Diagonal Up-Right
		{-1, 1},  // Diagonal Down-Left
	}

	type Point struct {
		x int
		y int
	}

	topLeft := Point{x: x - 1, y: y - 1}
	topRight := Point{x: x + 1, y: y - 1}
	bottomLeft := Point{x: x - 1, y: y + 1}
	bottomRight := Point{x: x + 1, y: y + 1}

	sum := 0
	sum += countWord(grid, topLeft.x, topLeft.y, word, []Direction{directions[0]})
	sum += countWord(grid, topRight.x, topRight.y, word, []Direction{directions[3]})
	sum += countWord(grid, bottomLeft.x, bottomLeft.y, word, []Direction{directions[2]})
	sum += countWord(grid, bottomRight.x, bottomRight.y, word, []Direction{directions[1]})

	return sum >= 2
}
