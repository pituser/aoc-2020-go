package main

import (
	"fmt"

	"github.com/pituser/aoc-2020-go/util"
)

const (
	day = "6"
)

type groupForm struct {
	numMembers int
	answers    map[byte]int
}

func newGroupForm() *groupForm {
	return &groupForm{0, make(map[byte]int)}
}

func readPuzzleInput() []groupForm {
	input, err := util.ReadLinesFromFile("input.txt")
	util.CheckError(err, "error reading puzzle input file")

	var answers []groupForm

	groupResult := newGroupForm()
	numMembers := 0
	for _, line := range input {
		if line != "" {
			numMembers++
			for i := 0; i < len(line); i++ {
				groupResult.answers[line[i]]++
			}
		} else {
			groupResult.numMembers = numMembers
			answers = append(answers, *groupResult)
			groupResult = newGroupForm()
			numMembers = 0
		}
	}
	groupResult.numMembers = numMembers
	answers = append(answers, *groupResult)

	return answers
}

func solvePartOne(answers []groupForm) int {
	sum := 0
	for _, answer := range answers {
		sum += len(answer.answers)
	}
	return sum
}

func solvePartTwo(answers []groupForm) int {
	sum := 0
	for _, answer := range answers {
		for _, v := range answer.answers {
			if v == answer.numMembers {
				sum++
			}
		}
	}
	return sum
}

func main() {
	fmt.Println("Advent of Code 2020 - Day ", day)

	input := readPuzzleInput()
	fmt.Println(input)

	fmt.Println("Part One: ", solvePartOne(input))
	fmt.Println("Part Two: ", solvePartTwo(input))
}
