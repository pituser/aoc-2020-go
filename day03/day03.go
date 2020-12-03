package main

import (
	"fmt"

	"github.com/pituser/aoc-2020-go/util"
)

const (
	day = "3"
)

func readPuzzleInput() []string {
	input, err := util.ReadLinesFromFile("input.txt")
	util.CheckError(err, "error reading puzzle input file")

	return input
}

type pos struct {
	x, y int
}

func countTreesOnPath(treeMap []string, slope pos) int {
	tree := byte('#')

	width := len(treeMap[0])
	height := len(treeMap)

	numTrees := 0
	currentPos := pos{0, 0}

	for currentPos.y < height {
		if treeMap[currentPos.y][currentPos.x] == tree {
			numTrees++
		}
		currentPos.x = (currentPos.x + slope.x) % width
		currentPos.y += slope.y
	}

	return numTrees
}

func solvePartOne(treeMap []string) int {
	return countTreesOnPath(treeMap, pos{3, 1})
}

func solvePartTwo(treeMap []string) int {
	slopes := []pos{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	result := 1
	for _, slope := range slopes {
		result *= countTreesOnPath(treeMap, slope)
	}

	return result
}

func main() {
	fmt.Println("Advent of Code 2020 - Day ", day)

	input := readPuzzleInput()

	fmt.Println("Part One: ", solvePartOne(input))
	fmt.Println("Part Two: ", solvePartTwo(input))
}
