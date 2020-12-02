package main

import (
	"aoc2020/util"
	"fmt"
)

const (
	day = "XX"
)

func readPuzzleInput() []string {
	input, err := util.ReadLinesFromFile("input.txt")
	util.CheckError(err, "error reading puzzle input file")

	return input
}

func solvePartOne() int {
	return 42
}

func solvePartTwo() int {
	return -1
}

func main() {
	fmt.Println("Advent of Code 2020 - Day ", day)

	input := readPuzzleInput()
	fmt.Println(input)

	fmt.Println("Part One: ", solvePartOne())
	fmt.Println("Part Two: ", solvePartTwo())
}
