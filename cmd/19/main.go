package main

import (
	"fmt"
	"math"
	"strconv"

	lib "github.com/thebenkogan/Advent-Of-Code-2016"
)

type node struct {
	position int
	next     *node
}

func createRing(total int) *node {
	fst := &node{0, nil}
	tmp := fst
	for i := 1; i < total; i++ {
		next := &node{i, nil}
		tmp.next = next
		tmp = next
	}
	tmp.next = fst
	return fst
}

func partTwo(total int) int {
	curr := createRing(total)
	totalLeft := total
	for {
		if curr == curr.next {
			break
		}

		// how many places away is the next elf to steal from?
		// steps forward = ceil((totalLeft - 1) / 2)
		stepsForward := int(math.Ceil(float64(totalLeft-1) / 2))
		beforeToSteal := curr
		for range stepsForward - 1 {
			beforeToSteal = beforeToSteal.next
		}
		beforeToSteal.next = beforeToSteal.next.next

		totalLeft--
		curr = curr.next
	}
	return curr.position + 1
}

func main() {
	input := lib.GetInput()
	total, _ := strconv.Atoi(input)

	curr := createRing(total)
	for {
		if curr == curr.next {
			break
		}
		curr.next = curr.next.next
		curr = curr.next
	}

	fmt.Println(curr.position + 1)

	// run the partTwo function on smaller inputs to see the pattern, e.g.:
	_ = partTwo(5)
	// next answer = prev + 1 if previous is less than half of previous total
	// otherwise, next answer = prev + 2
	// when previous equals previous total, reset to 1

	prevAns := 1
	for totalElves := 3; totalElves <= total; totalElves++ {
		if prevAns == totalElves-1 {
			prevAns = 1
		} else if float64(prevAns) < (float64(totalElves-1) / 2) {
			prevAns++
		} else {
			prevAns += 2
		}
	}

	fmt.Println(prevAns)
}
