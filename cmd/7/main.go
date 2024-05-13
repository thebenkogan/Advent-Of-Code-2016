package main

import (
	"fmt"
	"strings"

	lib "github.com/thebenkogan/Advent-Of-Code-2016"
)

func main() {
	input := lib.GetInput()
	total := 0
	for _, line := range strings.Split(input, "\n") {
		inHypernet := false
		numInside := 0
		numOutside := 0
		for i := 0; i < len(line)-3; i++ {
			switch rune(line[i]) {
			case '[':
				inHypernet = true
			case ']':
				inHypernet = false
			default:
				if line[i] == line[i+3] && line[i+1] == line[i+2] && line[i] != line[i+1] {
					if inHypernet {
						numInside += 1
					} else {
						numOutside += 1
					}
				}
			}
		}

		if numOutside > 0 && numInside == 0 {
			total += 1
		}
	}

	fmt.Println(total)

	total = 0
	for _, line := range strings.Split(input, "\n") {
		inHypernet := false
		inside := make(map[string]bool)
		outside := make(map[string]bool)
		for i := 0; i < len(line)-2; i++ {
			switch rune(line[i]) {
			case '[':
				inHypernet = true
			case ']':
				inHypernet = false
			default:
				if line[i] == line[i+2] && line[i] != line[i+1] {
					if inHypernet {
						inside[line[i:i+3]] = true
					} else {
						outside[line[i:i+3]] = true
					}
				}
			}
		}

		for key := range inside {
			if outside[string(key[1])+string(key[0])+string(key[1])] {
				total += 1
				break
			}
		}
	}

	fmt.Println(total)
}
