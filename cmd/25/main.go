package main

import "fmt"

// the program takes the input, adds 15*170, then repeatedly floor divides by 2 until zero
// after each division, it should change from even to odd or odd to even
// at zero, it wraps around, so the last division should be odd (since first must be even to transmit 0)
// brute force and find this even number, it's not that big anyways

func main() {
	for n := 2; ; n += 2 {
		a := n + 15*170
		even := true
		good := true
		for a > 0 {
			if a%2 == 1 && even || a%2 == 0 && !even {
				good = false
				break
			}
			a = a / 2
			even = !even
		}
		if good {
			fmt.Println(n)
			break
		}
	}
}
