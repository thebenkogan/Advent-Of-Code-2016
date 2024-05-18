package main

import (
	"fmt"
	"strings"

	lib "github.com/thebenkogan/Advent-Of-Code-2016"
)

func rect(screen [][]bool, a, b int) {
	for i := 0; i < a; i++ {
		for j := 0; j < b; j++ {
			screen[j][i] = true
		}
	}
}

func rotateRow(screen [][]bool, a, b int) {
	newRow := make([]bool, len(screen[a]))
	for i, v := range screen[a] {
		newRow[(i+b)%len(screen[a])] = v
	}
	screen[a] = newRow
}

func rotateCol(screen [][]bool, a, b int) {
	newCol := make([]bool, len(screen))
	for i, row := range screen {
		newCol[(i+b)%len(screen)] = row[a]
	}
	for i, row := range screen {
		row[a] = newCol[i]
	}
}

func main() {
	input := lib.GetInput()

	screen := make([][]bool, 6)
	for i := range screen {
		screen[i] = make([]bool, 50)
	}

	for _, line := range strings.Split(input, "\n") {
		nums := lib.ParseNums(line)
		switch {
		case strings.HasPrefix(line, "rect"):
			rect(screen, nums[0], nums[1])
		case strings.HasPrefix(line, "rotate row"):
			rotateRow(screen, nums[0], nums[1])
		case strings.HasPrefix(line, "rotate column"):
			rotateCol(screen, nums[0], nums[1])
		}
	}

	total := 0
	for _, row := range screen {
		for _, v := range row {
			if v {
				total++
			}
		}
	}

	fmt.Println(total)

	for _, row := range screen {
		s := strings.Builder{}
		for _, v := range row {
			if v {
				s.WriteRune('#')
			} else {
				s.WriteRune('.')
			}
		}
		fmt.Println(s.String())
	}
}
