package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"

	"github.com/pituser/aoc-2020-go/util"
)

const (
	day = 2
)

type passwordWithPolicy struct {
	policy1, policy2 int
	letter           byte
	password         []byte
}

func readPuzzleInput() []passwordWithPolicy {
	input, err := util.ReadLinesFromFile("input.txt")
	util.CheckError(err, "error reading puzzle input file")

	return parsePasswordLines(input)
}

func parsePasswordLines(lines []string) []passwordWithPolicy {
	var passwordsWithPolicy []passwordWithPolicy
	re, err := regexp.Compile(`(\d+)-(\d+) (\w): (\w+)`)
	util.CheckError(err, "error compiling regexp")

	for _, line := range lines {
		parts := re.FindStringSubmatch(line)
		p1, err := strconv.Atoi(parts[1])
		util.CheckError(err, "error parsing int")
		p2, err := strconv.Atoi(parts[2])
		util.CheckError(err, "error parsing int")

		spec := passwordWithPolicy{
			policy1:  p1,
			policy2:  p2,
			letter:   (parts[3])[0], // by definition of the regexp this is only one byte
			password: []byte(parts[4]),
		}

		passwordsWithPolicy = append(passwordsWithPolicy, spec)
	}
	return passwordsWithPolicy
}

func checkPasswordPart1(spec passwordWithPolicy) bool {
	cnt := bytes.Count(spec.password, []byte{spec.letter})
	return cnt >= spec.policy1 && cnt <= spec.policy2
}

func checkPasswordPart2(spec passwordWithPolicy) bool {
	return (spec.password[spec.policy1-1] == spec.letter) != (spec.password[spec.policy2-1] == spec.letter)
}

func countCorrectPasswords(specs []passwordWithPolicy, checker func(passwordWithPolicy) bool) int {
	correctPasswordsCnt := 0
	for _, spec := range specs {
		if checker(spec) {
			correctPasswordsCnt++
		}
	}
	return correctPasswordsCnt
}

func solvePartOne(passwordsWithPolicy []passwordWithPolicy) int {
	return countCorrectPasswords(passwordsWithPolicy, checkPasswordPart1)
}

func solvePartTwo(passwordsWithPolicy []passwordWithPolicy) int {
	return countCorrectPasswords(passwordsWithPolicy, checkPasswordPart2)
}

func main() {
	fmt.Println("Advent of Code 2020 - Day ", day)

	passwordsWithPolicy := readPuzzleInput()

	fmt.Println("Part One: ", solvePartOne(passwordsWithPolicy))
	fmt.Println("Part Two: ", solvePartTwo(passwordsWithPolicy))
}
