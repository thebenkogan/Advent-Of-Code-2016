package main

import (
	"fmt"
	"strconv"
	"strings"

	lib "github.com/thebenkogan/Advent-Of-Code-2016"
)

func run(instructions []string, registers map[string]int) int {
	pc := 0
	for pc >= 0 && pc < len(instructions) {
		line := instructions[pc]
		split := strings.Split(line, " ")
		offset := 1
		switch {
		case strings.HasPrefix(line, "cpy"):
			n, err := strconv.Atoi(split[1])
			if err == nil {
				registers[split[2]] = n
			} else {
				registers[split[2]] = registers[split[1]]
			}
		case strings.HasPrefix(line, "inc"):
			registers[split[1]]++
		case strings.HasPrefix(line, "dec"):
			registers[split[1]]--
		case strings.HasPrefix(line, "jnz"):
			n, err := strconv.Atoi(split[1])
			jmp, _ := strconv.Atoi(split[2])
			if err == nil && n != 0 || registers[split[1]] != 0 {
				offset = jmp
			}
		}
		pc += offset
	}

	return registers["a"]
}

func main() {
	input := lib.GetInput()
	instructions := strings.Split(input, "\n")

	registers := map[string]int{"a": 0, "b": 0, "c": 0, "d": 0}
	fmt.Println(run(instructions, registers))

	registers = map[string]int{"a": 0, "b": 0, "c": 1, "d": 0}
	fmt.Println(run(instructions, registers))
}
