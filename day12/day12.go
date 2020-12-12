package main

import (
	"fmt"
	"strconv"

	"github.com/pituser/aoc-2020-go/util"
)

const (
	day = "12"
)

const (
	N = byte('N')
	S = byte('S')
	E = byte('E')
	W = byte('W')
	L = byte('L')
	R = byte('R')
	F = byte('F')
)

type instruction struct {
	action byte
	param  int
}

type position struct {
	x, y int
}

func readPuzzleInput() []instruction {
	input, err := util.ReadLinesFromFile("input.txt")
	util.CheckError(err, "error reading puzzle input file")

	var instructions []instruction
	for _, line := range input {
		var instr instruction
		instr.action = line[0]
		val, err := strconv.Atoi(line[1:])
		util.CheckError(err, "parse error")
		instr.param = val
		instructions = append(instructions, instr)
	}

	return instructions
}

func turn(p *position, direction byte, degree int) {
	for i := 0; i < degree/90; i++ {
		if direction == R {
			p.x, p.y = p.y, -p.x
		} else {
			p.x, p.y = -p.y, p.x
		}
	}
}

func driveShip(instructions []instruction, waypoint position, waypointMode bool) position {
	shipPos := position{0, 0}

	var posToMove *position
	if waypointMode {
		posToMove = &waypoint
	} else {
		posToMove = &shipPos
	}

	for _, instr := range instructions {
		switch instr.action {
		case N:
			posToMove.y += instr.param
		case S:
			posToMove.y -= instr.param
		case E:
			posToMove.x += instr.param
		case W:
			posToMove.x -= instr.param
		case L, R:
			turn(&waypoint, instr.action, instr.param)
		case F:
			shipPos.x += instr.param * waypoint.x
			shipPos.y += instr.param * waypoint.y
		default:
			panic("unknown instruction")
		}
	}
	return shipPos
}

func solvePartOne(instructions []instruction) int {
	waypoint := position{1, 0}
	finalPos := driveShip(instructions, waypoint, false)
	return util.AbsInt(finalPos.x) + util.AbsInt(finalPos.y)
}

func solvePartTwo(instructions []instruction) int {

	waypoint := position{10, 1}
	finalPos := driveShip(instructions, waypoint, true)
	return util.AbsInt(finalPos.x) + util.AbsInt(finalPos.y)
}

func main() {
	fmt.Println("Advent of Code 2020 - Day ", day)

	input := readPuzzleInput()
	fmt.Println("Part One: ", solvePartOne(input))
	fmt.Println("Part Two: ", solvePartTwo(input))
}
