package main

import (
	"fmt"

	lib "github.com/thebenkogan/Advent-Of-Code-2016"
)

func getDecompressedLength(s string, v2 bool) int {
	i := 0
	decompressedLength := 0
	parenRegion := ""
	for i < len(s) {
		switch s[i] {
		case '(':
			parenRegion += "("
		case ')':
			nums := lib.ParseNums(parenRegion)
			repeated := s[i+1 : i+1+nums[0]]
			var repeatedLength int
			if v2 {
				repeatedLength = getDecompressedLength(repeated, true)
			} else {
				repeatedLength = len(repeated)
			}
			decompressedLength += repeatedLength * nums[1]
			i += nums[0]
			parenRegion = ""
		default:
			if len(parenRegion) > 0 {
				parenRegion += string(s[i])
			} else {
				decompressedLength += 1
			}
		}
		i++
	}
	return decompressedLength
}

func main() {
	input := lib.GetInput()
	fmt.Println(getDecompressedLength(input, false))
	fmt.Println(getDecompressedLength(input, true))
}
