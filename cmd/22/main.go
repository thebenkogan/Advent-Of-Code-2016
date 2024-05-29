package main

import (
	"fmt"
	"strings"

	lib "github.com/thebenkogan/Advent-Of-Code-2016"
)

type node struct {
	size  int
	used  int
	avail int
}

func countViable(adj [][]*node) int {
	total := 0
	for y := 0; y < len(adj); y++ {
		for x := 0; x < len(adj[0]); x++ {
			for oy := 0; oy < len(adj); oy++ {
				for ox := 0; ox < len(adj[0]); ox++ {
					if y == oy && x == ox {
						continue
					}
					if adj[y][x].used == 0 {
						continue
					}
					if adj[y][x].used <= adj[oy][ox].avail {
						total++
					}
				}
			}
		}
	}
	return total
}

func printGrid(adj [][]*node, zero *node) {
	for y := 0; y < len(adj); y++ {
		for x := 0; x < len(adj[0]); x++ {
			if adj[y][x].used > zero.size {
				fmt.Print("#")
			} else if adj[y][x] == zero {
				fmt.Print("_")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func main() {
	input := lib.GetInput()

	lines := strings.Split(input, "\n")[2:]
	adj := make([][]*node, 0) // weird data structure because I was thinking about graph traversal
	var zero *node
	for _, line := range lines {
		nums := lib.ParseNums(line)
		fileNode := node{
			size:  nums[2],
			used:  nums[3],
			avail: nums[4],
		}
		y := nums[1]
		if fileNode.used == 0 {
			zero = &fileNode
		}
		if y >= len(adj) {
			adj = append(adj, make([]*node, 0))
		}
		adj[y] = append(adj[y], &fileNode)
	}

	fmt.Println(countViable(adj))

	// for part 2, don't actually do anything hard
	// use this function to print the grid:
	printGrid(adj, zero)
	// then, notice that we just have to move the empty node to the top right
	// and then from there, shuffle the goal data to the left
	// there is an intentional wall placed between the empty node and the top right
	// so first calculate how many steps it takes to get around the wall
	// then starting at the top right, it takes 5 moves for each shuffle, times the
	// distance to the left edge minus 1. finally, add 1 for moving the goal data to the origin
	fmt.Println(236)
}
