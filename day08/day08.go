package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pituser/aoc-2020-go/util"
)

const (
	day = "8"
)

const (
	NOP = iota
	ACC
	JMP
	ILLEGAL = 99
)

type Op struct {
	code  int
	param int
}

func readPuzzleInput() []Op {
	input, err := util.ReadLinesFromFile("input.txt")
	util.CheckError(err, "error reading puzzle input file")

	var program []Op

	for _, line := range input {
		var op Op

		parts := strings.Split(line, " ")
		switch parts[0] {
		case "nop":
			op.code = NOP
		case "acc":
			op.code = ACC
		case "jmp":
			op.code = JMP
		default:
			op.code = ILLEGAL
		}

		p, _ := strconv.Atoi(parts[1])
		op.param = p
		program = append(program, op)
	}

	return program
}

func run(program []Op) (acc int, infinite bool) {
	executed := make([]bool, len(program))

	acc = 0
	ip := 0

	var op *Op
	op = &program[ip]
	for ip < len(program) && executed[ip] == false {
		executed[ip] = true
		switch op.code {
		case NOP:
			ip++
		case ACC:
			acc += op.param
			ip++
		case JMP:
			ip += op.param
		default:
			util.Error("illegal opcode")
		}
		if ip < len(program) {
			op = &program[ip]
		}
	}

	return acc, ip < len(program)
}

func solvePartOne(program []Op) int {
	acc, _ := run(program)
	return acc
}

func switchNopJmpOpcode(op *Op) bool {
	switched := false
	if op.code == NOP {
		op.code = JMP
		switched = true
	} else if op.code == JMP {
		op.code = NOP
		switched = true
	}
	return switched
}

func solvePartTwo(program []Op) int {
	for i := range program {
		if switchNopJmpOpcode(&program[i]) {
			acc, infinite := run(program)
			if !infinite {
				return acc
			}
			switchNopJmpOpcode(&program[i])
		}
	}
	return -1
}

func main() {
	fmt.Println("Advent of Code 2020 - Day ", day)

	input := readPuzzleInput()
	program := make([]Op, len(input))

	copy(program, input)
	fmt.Println("Part One: ", solvePartOne(input))
	copy(program, input)
	fmt.Println("Part Two: ", solvePartTwo(input))

}
