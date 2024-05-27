package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	lib "github.com/thebenkogan/Advent-Of-Code-2016"
)

type Range struct {
	start int
	end   int
}

func main() {
	input := lib.GetInput()

	ranges := make([]Range, 0)
	for _, line := range strings.Split(input, "\n") {
		split := strings.Split(line, "-")
		n1, _ := strconv.Atoi(split[0])
		n2, _ := strconv.Atoi(split[1])
		ranges = append(ranges, Range{n1, n2})
	}
	slices.SortFunc(ranges, func(r1, r2 Range) int {
		return r1.start - r2.start
	})

	prevMax := 0
	for _, r := range ranges {
		if r.start <= prevMax+1 {
			prevMax = max(prevMax, r.end)
			continue
		}
		fmt.Println(prevMax + 1)
		break
	}

	prevMax = 0
	total := 0
	for _, r := range ranges {
		if r.start <= prevMax+1 {
			prevMax = max(prevMax, r.end)
			continue
		}
		total += r.start - prevMax - 1
		prevMax = r.end
	}

	fmt.Println(total)
}
