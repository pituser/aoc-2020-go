package main

import (
	"fmt"

	"github.com/ernestosuarez/itertools"
	"github.com/pituser/aoc-2020-go/util"
)

const (
	day = 1
)

func readPuzzleInput() []int {
	input, err := util.ReadLinesFromFile("input.txt")
	util.CheckError(err, "error reading input file")
	expenses, err := util.ParseIntList(input)
	util.CheckError(err, "error parsing input file")
	return expenses
}

func findSubListInt(list []int, length int, match func([]int) bool) ([]int, bool) {
	for v := range itertools.CombinationsInt(list, length) {
		if match(v) {
			return v, true
		}
	}
	return make([]int, 0), false
}

func solve(expenses []int, length int) int {
	list, _ := findSubListInt(expenses, length, func(l []int) bool {
		return util.ReduceInt(l, 0, func(a, b int) int {
			return a + b
		}) == 2020
	})
	return util.ReduceInt(list, 1, func(a, b int) int {
		return a * b
	})
}

func solvePartOne(expenses []int) int {
	return solve(expenses, 2)
}

func solvePartTwo(expenses []int) int {
	return solve(expenses, 3)
}

func main() {
	fmt.Println("Advent of Code 2020 - Day ", day)

	expenses := readPuzzleInput()

	fmt.Println("Part One: ", solvePartOne(expenses))
	fmt.Println("Part Two: ", solvePartTwo(expenses))
}
