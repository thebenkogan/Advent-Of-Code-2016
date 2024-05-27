package main

import (
	"fmt"

	lib "github.com/thebenkogan/Advent-Of-Code-2016"
)

func countSafe(input string, numRows int) int {
	totalSafe := 0

	lastRow := make([]bool, 0, len(input))
	for _, c := range input {
		isTrap := c == '^'
		lastRow = append(lastRow, isTrap)
		if !isTrap {
			totalSafe++
		}
	}

	for row := 1; row < numRows; row++ {
		nextRow := make([]bool, 0, len(input))
		for i := 0; i < len(input); i++ {
			left := false
			if i > 0 {
				left = lastRow[i-1]
			}
			center := lastRow[i]
			right := false
			if i < len(input)-1 {
				right = lastRow[i+1]
			}
			isTrap := (left && center && !right) || (!left && center && right) || (left && !center && !right) || (!left && !center && right)
			nextRow = append(nextRow, isTrap)
			if !isTrap {
				totalSafe++
			}
		}
		lastRow = nextRow
	}

	return totalSafe
}

func main() {
	input := lib.GetInput()
	fmt.Println(countSafe(input, 40))
	fmt.Println(countSafe(input, 400000))
}
