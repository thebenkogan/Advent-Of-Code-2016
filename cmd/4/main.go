package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	lib "github.com/thebenkogan/Advent-Of-Code-2016"
)

func rotate(c rune, n int) rune {
	if c == '-' {
		if n%2 == 0 {
			return '-'
		} else {
			return ' '
		}
	}
	return rune((int(c)-'a'+n)%26 + 'a')
}

func main() {
	input := lib.GetInput()
	total := 0
	northPoleSectorId := 0
	for _, line := range strings.Split(input, "\n") {
		split := strings.Split(line, "-")
		name := strings.Join(split[:len(split)-1], "")
		counts := make(map[rune]int)
		for _, c := range name {
			counts[c]++
		}
		chars := make([]rune, 0)
		for k := range counts {
			chars = append(chars, k)
		}
		slices.SortFunc(chars, func(a, b rune) int {
			if counts[a] == counts[b] {
				return int(a - b)
			}
			return counts[b] - counts[a]
		})

		sectorId, _ := strconv.Atoi(strings.Split(split[len(split)-1], "[")[0])
		checksum := strings.Split(split[len(split)-1], "[")[1]
		checksum = strings.TrimSuffix(checksum, "]")

		if string(chars[:5]) == checksum {
			total += sectorId
		}

		dashedName := strings.Join(split[:len(split)-1], "-")
		decryptedBuilder := strings.Builder{}
		for _, c := range dashedName {
			decryptedBuilder.WriteRune(rotate(c, sectorId))
		}

		if decryptedBuilder.String() == "northpole-object-storage" {
			northPoleSectorId = sectorId
		}
	}

	fmt.Println(total)
	fmt.Println(northPoleSectorId)
}
