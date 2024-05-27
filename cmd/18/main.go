package main

import (
	"fmt"

	lib "github.com/thebenkogan/Advent-Of-Code-2016"
)

func countSafe(input string, numRows int) int {
	grid := make([][]bool, 0, numRows)
	grid = append(grid, make([]bool, 0, len(input)))
	for _, c := range input {
		grid[0] = append(grid[0], c == '^')
	}

	for row := 1; row < numRows; row++ {
		grid = append(grid, make([]bool, 0, len(input)))
		for i := 0; i < len(input); i++ {
			left := false
			if i > 0 {
				left = grid[row-1][i-1]
			}
			center := grid[row-1][i]
			right := false
			if i < len(input)-1 {
				right = grid[row-1][i+1]
			}
			isTrap := (left && center && !right) || (!left && center && right) || (left && !center && !right) || (!left && !center && right)
			grid[row] = append(grid[row], isTrap)
		}
	}

	total := 0
	for _, row := range grid {
		for _, isTrap := range row {
			if !isTrap {
				total++
			}
		}
	}

	return total
}

func main() {
	input := lib.GetInput()
	fmt.Println(countSafe(input, 40))
	fmt.Println(countSafe(input, 400000))
}
