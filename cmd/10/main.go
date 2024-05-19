package main

import (
	"fmt"
	"slices"
	"strings"

	lib "github.com/thebenkogan/Advent-Of-Code-2016"
)

type toType int

const (
	botBin toType = iota
	outputBin
)

type destination struct {
	to  int
	typ toType
}

func main() {
	input := lib.GetInput()

	bots := make(map[int][]int)
	output := make(map[int][]int)
	edges := make(map[int][]destination)
	for _, line := range strings.Split(input, "\n") {
		if strings.HasPrefix(line, "value") {
			nums := lib.ParseNums(line)
			bots[nums[1]] = append(bots[nums[1]], nums[0])
			slices.Sort(bots[nums[1]])
		}

		if strings.HasPrefix(line, "bot") {
			nums := lib.ParseNums(line)
			low := destination{to: nums[1]}
			if strings.Contains(line, fmt.Sprintf("to output %d", nums[1])) {
				low.typ = outputBin
			} else {
				low.typ = botBin
			}
			high := destination{to: nums[2]}
			if strings.Contains(line, fmt.Sprintf("to output %d", nums[2])) {
				high.typ = outputBin
			} else {
				high.typ = botBin
			}
			edges[nums[0]] = []destination{low, high}
		}
	}

	seen := make(map[int]bool)
	for {
		done := true
		for num, bot := range bots {
			if !seen[num] && len(bot) == 2 {
				if bot[0] == 17 && bot[1] == 61 {
					fmt.Println(num)
				}

				done = false
				seen[num] = true
				edge := edges[num]
				for i := range 2 {
					if edge[i].typ == outputBin {
						output[edge[i].to] = append(output[edge[i].to], bot[i])
						slices.Sort(output[edge[i].to])
					} else {
						bots[edge[i].to] = append(bots[edge[i].to], bot[i])
						slices.Sort(bots[edges[num][i].to])
					}
				}
				break
			}
		}
		if done {
			break
		}
	}

	fmt.Println(output[0][0] * output[1][0] * output[2][0])
}
