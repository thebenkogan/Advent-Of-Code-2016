package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	lib "github.com/thebenkogan/Advent-Of-Code-2016"
)

func getDoorStatus(passcode string, path string) map[string]bool {
	hash := md5.Sum([]byte(passcode + path))
	hex := hex.EncodeToString(hash[:])
	return map[string]bool{
		"U": hex[0] > 'a',
		"D": hex[1] > 'a',
		"L": hex[2] > 'a',
		"R": hex[3] > 'a',
	}
}

func translateDir(v lib.Vector) string {
	switch {
	case v.X == 1 && v.Y == 0:
		return "R"
	case v.X == -1 && v.Y == 0:
		return "L"
	case v.X == 0 && v.Y == 1:
		return "D"
	case v.X == 0 && v.Y == -1:
		return "U"
	}
	panic("Invalid direction")
}

type node struct {
	x    int
	y    int
	path string
}

func main() {
	passcode := lib.GetInput()

	queue := []node{{x: 0, y: 0, path: ""}}
	found := false
	longestPath := 0
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr.x == 3 && curr.y == 3 {
			if !found {
				fmt.Println(curr.path)
				found = true
			}
			longestPath = len(curr.path)
			continue
		}
		doorStatus := getDoorStatus(passcode, curr.path)
		for _, dir := range lib.DIRS {
			letter := translateDir(dir)
			next := node{x: curr.x + dir.X, y: curr.y + dir.Y, path: curr.path + letter}
			if next.x < 0 || next.y < 0 || next.x > 3 || next.y > 3 || !doorStatus[letter] {
				continue
			}
			queue = append(queue, next)
		}
	}
	fmt.Println(longestPath)
}
