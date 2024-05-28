package main

import (
	"fmt"
	"slices"
	"strings"

	lib "github.com/thebenkogan/Advent-Of-Code-2016"
)

func swapPos(s []rune, x, y int) []rune {
	s[x], s[y] = s[y], s[x]
	return s
}

func swapLetter(s []rune, x, y rune) []rune {
	for i, c := range s {
		if c == x {
			s[i] = y
		} else if c == y {
			s[i] = x
		}
	}
	return s
}

func rotateLeft(s []rune, x int) []rune {
	for range x {
		s = append(s[1:], s[0])
	}
	return s
}

func rotateRight(s []rune, x int) []rune {
	for range x {
		s = append([]rune{s[len(s)-1]}, s[:len(s)-1]...)
	}
	return s
}

func rotateBasedOnLetter(s []rune, x rune) []rune {
	i := slices.Index(s, x)
	if i >= 4 {
		i++
	}
	return rotateRight(s, i+1)
}

func reverse(s []rune, x, y int) []rune {
	slices.Reverse(s[x : y+1])
	return s
}

func move(s []rune, x, y int) []rune {
	c := s[x]
	s = append(s[:x], s[x+1:]...)
	return append(s[:y], append([]rune{c}, s[y:]...)...)
}

func run(s string, instructions []string) string {
	rs := []rune(s)
	for _, line := range instructions {
		switch {
		case strings.HasPrefix(line, "swap position"):
			nums := lib.ParseNums(line)
			rs = swapPos(rs, nums[0], nums[1])
		case strings.HasPrefix(line, "swap letter"):
			split := strings.Split(line, " ")
			rs = swapLetter(rs, rune(split[2][0]), rune(split[5][0]))
		case strings.HasPrefix(line, "rotate left"):
			steps := lib.ParseNums(line)[0]
			rs = rotateLeft(rs, steps)
		case strings.HasPrefix(line, "rotate right"):
			steps := lib.ParseNums(line)[0]
			rs = rotateRight(rs, steps)
		case strings.HasPrefix(line, "rotate based on position"):
			split := strings.Split(line, " ")
			rs = rotateBasedOnLetter(rs, rune(split[6][0]))
		case strings.HasPrefix(line, "reverse positions"):
			nums := lib.ParseNums(line)
			rs = reverse(rs, nums[0], nums[1])
		case strings.HasPrefix(line, "move position"):
			nums := lib.ParseNums(line)
			rs = move(rs, nums[0], nums[1])
		}
	}
	return string(rs)
}

func main() {
	input := lib.GetInput()
	instructions := strings.Split(input, "\n")

	fmt.Println(run("abcdefgh", instructions))

	for _, perm := range lib.StringPermutations("abcdefgh") {
		if run(perm, instructions) == "fbgdceah" {
			fmt.Println(perm)
			break
		}
	}
}
