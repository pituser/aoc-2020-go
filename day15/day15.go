package main

import (
	"fmt"
)

const (
	day = "15"
)

func readPuzzleInput() []int {
	return []int{1, 20, 8, 12, 0, 14}
}

func run(startingNumbers []int, maxTurn int) int {
	// map from spoken number to turn in which this number was most recently spoken
	memory := make(map[int]int)

	// starting numbers
	turn := 1
	for i := 0; i < len(startingNumbers); i++ {
		memory[startingNumbers[i]] = turn
		turn++
	}
	lastSpokenNumber := startingNumbers[len(startingNumbers)-1]
	firstTime := true

	// play
	for turn <= maxTurn {
		var nextNumber int

		if firstTime {
			nextNumber = 0
		} else {
			nextNumber = turn - memory[lastSpokenNumber] - 1
		}

		firstTime = (memory[nextNumber] == 0)
		memory[lastSpokenNumber] = turn - 1
		lastSpokenNumber = nextNumber
		turn++
	}

	return lastSpokenNumber

}

func solvePartOne(startingNumbers []int) int {
	return run(startingNumbers, 2020)
}

func solvePartTwo(startingNumbers []int) int {
	return run(startingNumbers, 30000000)
}

func main() {
	fmt.Println("Advent of Code 2020 - Day ", day)

	input := readPuzzleInput()

	fmt.Println("Part One: ", solvePartOne(input))
	fmt.Println("Part Two: ", solvePartTwo(input))
}
