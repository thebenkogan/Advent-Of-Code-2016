package main

import "fmt"

func findNumSatisfying(discs ...func(t int) bool) int {
	for t := 1; ; t++ {
		valid := true
		for _, disc := range discs {
			if !disc(t) {
				valid = false
				break
			}
		}
		if valid {
			return t
		}
	}
}

func main() {
	// Disc #1 has 17 positions; at time=0, it is at position 15. -> 15 + t + 1 == 0 (mod 17)
	// Disc #2 has 3 positions; at time=0, it is at position 2. ->  2 + t + 2 == 0 (mod 3)
	// Disc #3 has 19 positions; at time=0, it is at position 4. -> 4 + t + 3 == 0 (mod 19)
	// Disc #4 has 13 positions; at time=0, it is at position 2. -> 2 + t + 4 == 0 (mod 13)
	// Disc #5 has 7 positions; at time=0, it is at position 2. -> 2 + t + 5 == 0 (mod 7)
	// Disc #6 has 5 positions; at time=0, it is at position 0. -> t + 6 == 0 (mod 5)
	// plug into Chinese Remainder Theorem solver, or just brute force it, numbers aren't that big

	validDisc1 := func(t int) bool {
		return (t+1+15)%17 == 0
	}

	validDisc2 := func(t int) bool {
		return (t+2+2)%3 == 0
	}

	validDisc3 := func(t int) bool {
		return (t+3+4)%19 == 0
	}

	validDisc4 := func(t int) bool {
		return (t+4+2)%13 == 0
	}

	validDisc5 := func(t int) bool {
		return (t+5+2)%7 == 0
	}

	validDisc6 := func(t int) bool {
		return (t+6)%5 == 0
	}

	fmt.Println(findNumSatisfying(validDisc1, validDisc2, validDisc3, validDisc4, validDisc5, validDisc6))

	// add one more disc:
	// Disc #7 has 11 positions; at time=0, it is at position 0. -> t + 7 == 0 (mod 11)

	validDisc7 := func(t int) bool {
		return (t+7)%11 == 0
	}

	fmt.Println(findNumSatisfying(validDisc1, validDisc2, validDisc3, validDisc4, validDisc5, validDisc6, validDisc7))
}
