package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
	"strconv"

	lib "github.com/thebenkogan/Advent-Of-Code-2016"
)

func stretchHash(hash string) string {
	for range 2016 {
		hashBytes := md5.Sum([]byte(hash))
		hash = hex.EncodeToString(hashBytes[:])
	}
	return hash
}

func findAllRepeating(s string, n int) []rune {
	good := make([]rune, 0, len(s))
	for i := 0; i < len(s)-n+1; i++ {
		valid := true
		for j := range n {
			if s[i] != s[i+j] {
				valid = false
				break
			}
		}
		if valid {
			good = append(good, rune(s[i]))
		}
	}
	return good
}

func get64thKey(salt string, withStretch bool) int {
	past := make(map[rune][]int)
	validKeyIndexes := make([]int, 0, 75)
	seenKeyIndexes := make(map[int]bool)
	for i := 0; ; i++ {
		combined := salt + strconv.Itoa(i)
		hashBytes := md5.Sum([]byte(combined))
		hex := hex.EncodeToString(hashBytes[:])
		if withStretch {
			hex = stretchHash(hex)
		}

		tripleChars := findAllRepeating(hex, 3)
		if len(tripleChars) == 0 {
			// there can't be quintuples without triples
			continue
		}
		tripleChar := tripleChars[0]

		for _, quintupleChar := range findAllRepeating(hex, 5) {
			for _, index := range past[quintupleChar] {
				if index >= i-1000 && !seenKeyIndexes[index] {
					seenKeyIndexes[index] = true
					validKeyIndexes = append(validKeyIndexes, index)
					if len(validKeyIndexes) == 75 {
						// we find 75 (chosen experimentally) valid keys, then sort and take the 64th
						// this is because this algorithm generates the keys out of order
						// since a quintuple might validate a later index first
						sort.Ints(validKeyIndexes)
						return validKeyIndexes[63]
					}
				}
			}
		}

		past[tripleChar] = append(past[tripleChar], i)
	}
}

func main() {
	salt := lib.GetInput()

	fmt.Println(get64thKey(salt, false))
	fmt.Println(get64thKey(salt, true))
}
