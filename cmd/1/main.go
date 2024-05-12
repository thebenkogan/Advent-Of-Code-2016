package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	lib "github.com/thebenkogan/Advent-Of-Code-2016"
)

func main() {
	input := lib.GetInput()
	split := strings.Split(input, ", ")

	x, y := 0, 0
	dx, dy := 0, 1
	for _, s := range split {
		if s[0] == 'R' {
			dx, dy = dy, -dx
		} else {
			dx, dy = -dy, dx
		}

		n, _ := strconv.Atoi(s[1:])
		x += n * dx
		y += n * dy
	}

	fmt.Println(math.Abs(float64(x)) + math.Abs(float64(y)))

	x, y = 0, 0
	dx, dy = 0, 1
	seen := map[string]bool{}
	for _, s := range split {
		if s[0] == 'R' {
			dx, dy = dy, -dx
		} else {
			dx, dy = -dy, dx
		}

		n, _ := strconv.Atoi(s[1:])
		for range n {
			x += dx
			y += dy
			hash := fmt.Sprintf("%d,%d", x, y)
			if seen[hash] {
				fmt.Println(math.Abs(float64(x)) + math.Abs(float64(y)))
				return
			} else {
				seen[hash] = true
			}
		}
	}
}
