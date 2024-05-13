package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	lib "github.com/thebenkogan/Advent-Of-Code-2016"
)

func main() {
	id := lib.GetInput()
	password1 := make([]rune, 0)
	password2 := make([]rune, 8)
	seenIndexes := make(map[int]bool)
	i := 0
	for len(seenIndexes) < 8 {
		hash := md5.New()
		hash.Write([]byte(id + strconv.Itoa(i)))
		sum := hash.Sum(nil)
		hex := hex.EncodeToString(sum)
		if strings.HasPrefix(hex, "00000") {
			if len(password1) < 8 {
				password1 = append(password1, rune(hex[5]))
			}
			index, err := strconv.Atoi(string(hex[5]))
			if err == nil && index < 8 && !seenIndexes[index] {
				password2[index] = rune(hex[6])
				seenIndexes[index] = true
			}
		}
		i += 1
	}

	fmt.Println(string(password1))
	fmt.Println(string(password2))
}
