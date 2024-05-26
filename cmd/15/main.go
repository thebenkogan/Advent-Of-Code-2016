package main

import "fmt"

func main() {
	// Disc #1 has 17 positions; at time=0, it is at position 15. -> 15 + t + 1 == 0 (mod 17)
	// Disc #2 has 3 positions; at time=0, it is at position 2. ->  2 + t + 2 == 0 (mod 3)
	// Disc #3 has 19 positions; at time=0, it is at position 4. -> 4 + t + 3 == 0 (mod 19)
	// Disc #4 has 13 positions; at time=0, it is at position 2. -> 2 + t + 4 == 0 (mod 13)
	// Disc #5 has 7 positions; at time=0, it is at position 2. -> 2 + t + 5 == 0 (mod 7)
	// Disc #6 has 5 positions; at time=0, it is at position 0. -> t + 6 == 0 (mod 5)
	// plug in to chinese remainder theorem calculator, get 400589
	fmt.Println(400589)

	// add one more disc:
	// Disc #7 has 11 positions; at time=0, it is at position 0. -> t + 7 == 0 (mod 11)
	fmt.Println(3045959)
}
