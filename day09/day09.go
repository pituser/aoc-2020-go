package main

import (
	"fmt"

	"github.com/pituser/aoc-2020-go/util"
)

const (
	day = "9"
)

func readPuzzleInput() []int {
	input, err := util.ReadLinesFromFile("input.txt")
	util.CheckError(err, "error reading puzzle input file")
	list, err := util.ParseIntList(input)
	util.CheckError(err, "parse error")
	return list
}

func findFirstNonSummable(numbers []int, preamble int) int {
	var found bool
	for n := preamble; n < len(numbers); n++ {
		found = false
		for i := n - preamble; i < n; i++ {
			for j := n - preamble + 1; j < n; j++ {
				if i != j && numbers[i]+numbers[j] == numbers[n] {
					found = true
					break
				}
			}
		}
		if !found {
			return numbers[n]
		}
	}
	return -1
}

func solvePartOne(numbers []int) int {
	return findFirstNonSummable(numbers, 25)
}

func solvePartTwo(numbers []int, target int) int {
	var min, max int

	for start := 0; start < len(numbers)-1; start++ {
		for end := start + 2; end < len(numbers); end++ {
			sum := 0
			min = numbers[start]
			max = numbers[start]
			for i := start; i < end; i++ {
				sum += numbers[i]
				if sum > target {
					break
				}
				if numbers[i] < min {
					min = numbers[i]
				}
				if numbers[i] > max {
					max = numbers[i]
				}
			}
			if sum == target {
				return min + max
			}
		}
	}
	return -1
}

func main() {
	fmt.Println("Advent of Code 2020 - Day ", day)

	input := readPuzzleInput()

	target := solvePartOne(input)
	fmt.Println("Part One: ", target)
	fmt.Println("Part Two: ", solvePartTwo(input, target))
}
