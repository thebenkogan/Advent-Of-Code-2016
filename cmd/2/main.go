package main

import (
	"fmt"
	"strconv"
	"strings"

	lib "github.com/thebenkogan/Advent-Of-Code-2016"
)

func main() {
	input := lib.GetInput()
	lines := strings.Split(input, "\n")

	num := 5
	code := ""
	for _, line := range lines {
		for _, c := range line {
			switch c {
			case 'U':
				if num-3 > 0 {
					num -= 3
				}
			case 'D':
				if num+3 < 10 {
					num += 3
				}
			case 'L':
				if num%3 != 1 {
					num -= 1
				}
			case 'R':
				if num%3 != 0 {
					num += 1
				}
			}
		}
		code += strconv.Itoa(num)
	}

	fmt.Println(code)

	curr := "5"
	code = ""
	for _, line := range lines {
		for _, c := range line {
			if next, ok := graph[curr][c]; ok {
				curr = next
			}
		}
		code += curr
	}

	fmt.Println(code)
}

var graph = map[string]map[rune]string{
	"1": {
		'D': "3",
	},
	"2": {
		'D': "6",
		'R': "3",
	},
	"3": {
		'U': "1",
		'D': "7",
		'L': "2",
		'R': "4",
	},
	"4": {
		'D': "8",
		'L': "3",
	},
	"5": {
		'R': "6",
	},
	"6": {
		'U': "2",
		'D': "A",
		'L': "5",
		'R': "7",
	},
	"7": {
		'U': "3",
		'D': "B",
		'L': "6",
		'R': "8",
	},
	"8": {
		'U': "4",
		'D': "C",
		'L': "7",
		'R': "9",
	},
	"9": {
		'L': "8",
	},
	"A": {
		'U': "6",
		'R': "B",
	},
	"B": {
		'U': "7",
		'D': "D",
		'L': "A",
		'R': "C",
	},
	"C": {
		'U': "8",
		'L': "B",
	},
	"D": {
		'U': "B",
	},
}
