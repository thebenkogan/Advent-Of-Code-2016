package main

import (
	"fmt"
	"strings"

	lib "github.com/thebenkogan/Advent-Of-Code-2016"
)

func step(s string) string {
	out := strings.Builder{}
	out.WriteString(s)
	out.WriteRune('0')
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '0' {
			out.WriteRune('1')
		} else {
			out.WriteRune('0')
		}
	}
	return out.String()
}

func checksum(s string) string {
	if len(s)%2 == 1 {
		return s
	}

	out := strings.Builder{}
	for i := 0; i < len(s); i += 2 {
		if s[i] == s[i+1] {
			out.WriteRune('1')
		} else {
			out.WriteRune('0')
		}
	}

	return checksum(out.String())
}

func checksumFromInitialState(input string, diskLength int) string {
	for len(input) < diskLength {
		input = step(input)
	}
	input = input[:diskLength]
	return checksum(input)
}

func main() {
	input := lib.GetInput()
	fmt.Println(checksumFromInitialState(input, 272))
	fmt.Println(checksumFromInitialState(input, 35651584))
}
