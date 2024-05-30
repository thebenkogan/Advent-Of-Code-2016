package main

import (
	"fmt"
	"strconv"
	"strings"

	lib "github.com/thebenkogan/Advent-Of-Code-2016"
)

// compute shortest path between all pairs of points
// generate all possible paths, pick shortest

type point struct {
	x int
	y int
}

func (p point) hash() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

func shortestDist(input []string, p1 point, p2 point) int {
	type node struct {
		p     point
		steps int
	}

	queue := []node{{p: p1, steps: 0}}
	seen := map[string]bool{p1.hash(): true}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr.p.x == p2.x && curr.p.y == p2.y {
			return curr.steps
		}
		for _, dir := range lib.DIRS {
			next := point{curr.p.x + dir.X, curr.p.y + dir.Y}
			if input[next.y][next.x] == '#' {
				continue
			}
			if _, ok := seen[next.hash()]; !ok {
				seen[next.hash()] = true
				queue = append(queue, node{next, curr.steps + 1})
			}
		}
	}

	panic("no path found")
}

type location struct {
	p   point
	num int
}

func main() {
	input := lib.GetInput()
	lines := strings.Split(input, "\n")

	locations := make([]location, 0)
	for y, line := range lines {
		for x, char := range line {
			n, err := strconv.Atoi(string(char))
			if err == nil {
				locations = append(locations, location{p: point{x, y}, num: n})
			}
		}
	}

	distances := make(map[int]map[int]int)
	for i := 0; i < len(locations); i++ {
		for j := i + 1; j < len(locations); j++ {
			d := shortestDist(lines, locations[i].p, locations[j].p)
			if _, ok := distances[locations[i].num]; !ok {
				distances[locations[i].num] = make(map[int]int)
			}
			distances[locations[i].num][locations[j].num] = d
			if _, ok := distances[locations[j].num]; !ok {
				distances[locations[j].num] = make(map[int]int)
			}
			distances[locations[j].num][locations[i].num] = d
		}
	}

	nums := ""
	for i := 1; i < len(locations); i++ {
		nums += strconv.Itoa(i)
	}
	combos := lib.StringPermutations(nums)

	shortestPath := 100000000
	for _, path := range combos {
		path = "0" + path
		dist := 0
		for i := 0; i < len(path)-1; i++ {
			a, _ := strconv.Atoi(string(path[i]))
			b, _ := strconv.Atoi(string(path[i+1]))
			dist += distances[a][b]
		}
		shortestPath = min(shortestPath, dist)
	}

	fmt.Println(shortestPath)

	// now all paths require 0 at the end

	shortestPath = 100000000
	for _, path := range combos {
		path = "0" + path + "0"
		dist := 0
		for i := 0; i < len(path)-1; i++ {
			a, _ := strconv.Atoi(string(path[i]))
			b, _ := strconv.Atoi(string(path[i+1]))
			dist += distances[a][b]
		}
		shortestPath = min(shortestPath, dist)
	}

	fmt.Println(shortestPath)
}
