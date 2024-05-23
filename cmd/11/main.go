package main

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	lib "github.com/thebenkogan/Advent-Of-Code-2016"
)

type square int

const (
	generator square = iota
	microchip
)

type item struct {
	name string
	typ  square
}

type state struct {
	elevator int
	floors   [4][]item
}

func (s *state) isSafe() bool {
	for _, floor := range s.floors {
		for _, item := range floor {
			if item.typ == microchip {
				hasSameGen := false
				hasDiffGen := false
				for _, other := range floor {
					if other.typ == generator {
						if other.name == item.name {
							hasSameGen = true
						} else {
							hasDiffGen = true
						}
					}
				}
				if !hasSameGen && hasDiffGen {
					return false
				}
			}
		}
	}
	return true
}

func (s *state) isDone() bool {
	for i := 0; i < 3; i++ {
		if len(s.floors[i]) > 0 {
			return false
		}
	}
	return true
}

func (s *state) hash() string {
	hash := strings.Builder{}
	hash.WriteString(strconv.Itoa(s.elevator))
	for _, floor := range s.floors {
		hash.WriteRune('|')
		slices.SortFunc(floor, func(i, j item) int {
			if i.name == j.name {
				return int(i.typ - j.typ)
			}
			return strings.Compare(i.name, j.name)
		})
		for _, item := range floor {
			hash.WriteString(item.name)
			hash.WriteRune(' ')
			hash.WriteString(strconv.Itoa(int(item.typ)))
		}
	}
	return hash.String()
}

func (s *state) copy() *state {
	newState := *s
	newState.floors = [4][]item{}
	for i, floor := range s.floors {
		newState.floors[i] = make([]item, len(floor))
		copy(newState.floors[i], floor)
	}
	return &newState
}

func (s *state) nextStates() []*state {
	var nexts []*state
	for i, floorItem := range s.floors[s.elevator] {
		for _, delta := range []int{-1, 1} {
			if s.elevator+delta < 0 || s.elevator+delta >= 4 {
				continue
			}

			// Try moving just this item
			next := s.copy()
			next.elevator += delta
			next.floors[s.elevator+delta] = append(next.floors[s.elevator+delta], floorItem)
			next.floors[s.elevator] = append(next.floors[s.elevator][:i], next.floors[s.elevator][i+1:]...)
			if next.isSafe() {
				nexts = append(nexts, next)
			}

			// Try moving two items
			for j := i + 1; j < len(s.floors[s.elevator]); j++ {
				next := s.copy()
				next.elevator += delta
				next.floors[s.elevator+delta] = append(next.floors[s.elevator+delta], floorItem)
				next.floors[s.elevator+delta] = append(next.floors[s.elevator+delta], next.floors[s.elevator][j])
				next.floors[s.elevator] = append(next.floors[s.elevator][:i], next.floors[s.elevator][i+1:]...)
				next.floors[s.elevator] = append(next.floors[s.elevator][:j-1], next.floors[s.elevator][j:]...)
				if next.isSafe() {
					nexts = append(nexts, next)
				}
			}
		}
	}

	return nexts
}

var MicrochipRegex = regexp.MustCompile(`a (\w+)-compatible microchip`)
var GeneratorRegex = regexp.MustCompile(`a (\w+) generator`)

func getMinSteps(s *state) int {
	type node struct {
		state *state
		steps int
	}

	queue := []node{{s, 0}}
	seen := map[string]bool{s.hash(): true}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr.state.isDone() {
			return curr.steps
		}
		for _, state := range curr.state.nextStates() {
			if _, ok := seen[state.hash()]; !ok {
				seen[state.hash()] = true
				queue = append(queue, node{state, curr.steps + 1})
			}
		}
	}

	panic("No solution found")
}

func main() {
	input := lib.GetInput()

	var initialState state
	initialState.floors = [4][]item{}
	for i, line := range strings.Split(input, "\n") {
		microchips := MicrochipRegex.FindAllStringSubmatch(line, -1)
		generators := GeneratorRegex.FindAllStringSubmatch(line, -1)
		for _, match := range microchips {
			initialState.floors[i] = append(initialState.floors[i], item{match[1], microchip})
		}
		for _, match := range generators {
			initialState.floors[i] = append(initialState.floors[i], item{match[1], generator})
		}
	}

	nextState := initialState.copy()
	fmt.Println(getMinSteps(&initialState))

	nextState.floors[0] = append(nextState.floors[0], item{"elerium", generator})
	nextState.floors[0] = append(nextState.floors[0], item{"elerium", microchip})
	nextState.floors[0] = append(nextState.floors[0], item{"dilithium", generator})
	nextState.floors[0] = append(nextState.floors[0], item{"dilithium", microchip})

	fmt.Println(getMinSteps(nextState))
}
