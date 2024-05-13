package main

import (
	"fmt"
	"strings"

	lib "github.com/thebenkogan/Advent-Of-Code-2016"
)

func isValid(triangle []int) bool {
	sum := 0
	for _, n := range triangle {
		sum += n
	}
	for _, n := range triangle {
		if sum-n <= n {
			return false
		}
	}
	return true
}

func main() {
	input := lib.GetInput()
	split := strings.Split(input, "\n")
	nums := make([][]int, len(split))
	for i, line := range split {
		nums[i] = lib.ParseNums(line)
	}

	total := 0
	for _, t := range nums {
		if isValid(t) {
			total++
		}
	}

	fmt.Println(total)

	total2 := 0
	for col := range 3 {
		for i := 0; i < len(nums); i += 3 {
			triangle := []int{nums[i][col], nums[i+1][col], nums[i+2][col]}
			if isValid(triangle) {
				total2++
			}
		}
	}

	fmt.Println(total2)
}
