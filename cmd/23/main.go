package main

import (
	"fmt"
	"strconv"
	"strings"

	lib "github.com/thebenkogan/Advent-Of-Code-2016"
)

type instruction interface {
	execute(registers map[string]int, instructions []instruction, pc int) int
	toggle() instruction
}

type cpy struct {
	src  string
	dest string
}

func (c cpy) execute(registers map[string]int, instructions []instruction, pc int) int {
	if _, err := strconv.Atoi(c.dest); err == nil {
		return pc + 1
	}
	n, err := strconv.Atoi(c.src)
	if err == nil {
		registers[c.dest] = n
	} else {
		registers[c.dest] = registers[c.src]
	}
	return pc + 1
}

func (c cpy) toggle() instruction {
	return jnz{test: c.src, offset: c.dest}
}

type jnz struct {
	test   string
	offset string
}

func (j jnz) execute(registers map[string]int, instructions []instruction, pc int) int {
	n, err := strconv.Atoi(j.test)
	offset, offsetErr := strconv.Atoi(j.offset)
	if err == nil && n != 0 || registers[j.test] != 0 {
		if offsetErr == nil {
			return pc + offset
		} else {
			return pc + registers[j.offset]
		}
	}
	return pc + 1
}

func (j jnz) toggle() instruction {
	return cpy{src: j.test, dest: j.offset}
}

type inc struct {
	reg string
}

func (i inc) execute(registers map[string]int, instructions []instruction, pc int) int {
	registers[i.reg]++
	return pc + 1
}

func (i inc) toggle() instruction {
	return dec(i)
}

type dec struct {
	reg string
}

func (d dec) execute(registers map[string]int, instructions []instruction, pc int) int {
	registers[d.reg]--
	return pc + 1
}

func (d dec) toggle() instruction {
	return inc(d)
}

type tgl struct {
	offsetReg string
}

func (t tgl) execute(registers map[string]int, instructions []instruction, pc int) int {
	offset := pc + registers[t.offsetReg]
	if offset >= 0 && offset < len(instructions) {
		instructions[offset] = instructions[offset].toggle()
	}
	return pc + 1
}

func (t tgl) toggle() instruction {
	return inc{reg: t.offsetReg}
}

type mul struct {
	dest string
	src1 string
	src2 string
}

func (m mul) execute(registers map[string]int, instructions []instruction, pc int) int {
	registers[m.dest] = registers[m.src1] * registers[m.src2]
	return pc + 1
}

func (m mul) toggle() instruction {
	panic("idk")
}

func run(input string, a int) int {
	instructions := make([]instruction, 0)
	for _, line := range strings.Split(input, "\n") {
		split := strings.Split(line, " ")
		switch {
		case strings.HasPrefix(line, "cpy"):
			instructions = append(instructions, cpy{src: split[1], dest: split[2]})
		case strings.HasPrefix(line, "inc"):
			instructions = append(instructions, inc{reg: split[1]})
		case strings.HasPrefix(line, "dec"):
			instructions = append(instructions, dec{reg: split[1]})
		case strings.HasPrefix(line, "jnz"):
			instructions = append(instructions, jnz{test: split[1], offset: split[2]})
		case strings.HasPrefix(line, "tgl"):
			instructions = append(instructions, tgl{offsetReg: split[1]})
		case strings.HasPrefix(line, "mul"):
			instructions = append(instructions, mul{dest: split[1], src1: split[2], src2: split[3]})
		}
	}

	registers := map[string]int{"a": a, "b": 0, "c": 0, "d": 0}
	pc := 0
	for pc >= 0 && pc < len(instructions) {
		pc = instructions[pc].execute(registers, instructions, pc)
	}

	return registers["a"]
}

func main() {
	input := lib.GetInput()
	fmt.Println(run(input, 7))
	fmt.Println(run(input, 12))
}
