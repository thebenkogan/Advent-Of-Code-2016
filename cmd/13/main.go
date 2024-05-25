package main

import (
	"fmt"
	"strconv"

	lib "github.com/thebenkogan/Advent-Of-Code-2016"
)

// adjust until it works
const GRID_SIZE = 50

func main() {
	input := lib.GetInput()
	favNum, _ := strconv.Atoi(input)

	grid := make([][]rune, GRID_SIZE)
	for i := range grid {
		grid[i] = make([]rune, GRID_SIZE)
	}

	for y := range grid {
		for x := range grid[y] {
			a := x*x + 3*x + 2*x*y + y + y*y + favNum
			bin := strconv.FormatInt(int64(a), 2)
			numOnes := 0
			for _, c := range bin {
				if c == '1' {
					numOnes++
				}
			}
			if numOnes%2 == 0 {
				grid[y][x] = '.'
			} else {
				grid[y][x] = '#'
			}
		}
	}

	type node struct {
		x     int
		y     int
		steps int
	}

	queue := []node{{x: 1, y: 1, steps: 0}}
	seen := map[string]bool{"1,1": true}
	totalUnder50 := 0
	var stepsToTarget int
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr.x == 31 && curr.y == 39 {
			stepsToTarget = curr.steps
		}
		if curr.steps <= 50 {
			totalUnder50++
		}
		for _, dir := range lib.DIRS {
			next := node{x: curr.x + dir.X, y: curr.y + dir.Y, steps: curr.steps + 1}
			if next.x < 0 || next.y < 0 || next.x >= len(grid[0]) || next.y >= len(grid) || grid[next.y][next.x] == '#' {
				continue
			}
			key := fmt.Sprintf("%d,%d", next.x, next.y)
			if seen[key] {
				continue
			}
			seen[key] = true
			queue = append(queue, next)
		}
	}

	fmt.Println(stepsToTarget)
	fmt.Println(totalUnder50)
}
