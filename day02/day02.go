package main

import (
	"aoc2020/util"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	day = 2
)

type passwordSpec struct {
	first, second int
	char          string
	password      string
}

func readPuzzleInput() []passwordSpec {
	input, err := util.ReadLinesFromFile("input.txt")
	if err != nil {
		log.Fatalf("error reading input file: %q", err)
	}
	return parsePasswordSpecs(input)
}

func parsePasswordSpecs(lines []string) []passwordSpec {
	var passwordSpecs []passwordSpec
	re, _ := regexp.Compile(`(\d+)-(\d+) (\w): (\w+)`)

	for _, line := range lines {
		var spec passwordSpec
		parts := re.FindStringSubmatch(line)
		spec.first, _ = strconv.Atoi(parts[1])
		spec.second, _ = strconv.Atoi(parts[2])
		spec.char = parts[3]
		spec.password = parts[4]
		passwordSpecs = append(passwordSpecs, spec)
	}
	return passwordSpecs
}

func checkPassword1(spec passwordSpec) bool {
	cnt := strings.Count(spec.password, spec.char)
	return cnt >= spec.first && cnt <= spec.second
}

func checkPassword2(spec passwordSpec) bool {
	return (spec.password[spec.first-1] == spec.char[0]) != (spec.password[spec.second-1] == spec.char[0])
}

func countCorrectPasswords(specs []passwordSpec, checker func(passwordSpec) bool) int {
	correctPasswords := 0
	for _, spec := range specs {
		if checker(spec) {
			correctPasswords++
		}
	}
	return correctPasswords
}

func solvePartOne(specs []passwordSpec) int {
	return countCorrectPasswords(specs, checkPassword1)
}

func solvePartTwo(specs []passwordSpec) int {
	return countCorrectPasswords(specs, checkPassword2)
}

func main() {
	fmt.Println("Advent of Code 2020 - Day ", day)

	passwordSpecs := readPuzzleInput()

	fmt.Println("Part One: ", solvePartOne(passwordSpecs))
	fmt.Println("Part Two: ", solvePartTwo(passwordSpecs))
}
