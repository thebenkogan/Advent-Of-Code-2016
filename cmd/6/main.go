package main

import (
	"fmt"
	"math"
	"strings"

	lib "github.com/thebenkogan/Advent-Of-Code-2016"
)

func main() {
	input := lib.GetInput()
	split := strings.Split(input, "\n")
	message1 := ""
	message2 := ""
	for col := range len(split[0]) {
		counts := make(map[rune]int)
		for _, row := range split {
			counts[rune(row[col])]++
		}
		max := 0
		var maxChar rune
		min := math.MaxInt32
		var minChar rune
		for char, count := range counts {
			if count > max {
				max = count
				maxChar = char
			}
			if count < min {
				min = count
				minChar = char
			}
		}
		message1 += string(maxChar)
		message2 += string(minChar)
	}

	fmt.Println(message1)
	fmt.Println(message2)
}
