package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pituser/aoc-2020-go/util"
)

const (
	day = "7"
)

type Color string

type BagContents struct {
	num   int
	color Color
}

type BagRules map[Color][]BagContents

var ruleRegExp *regexp.Regexp = regexp.MustCompile(`(\d+) ([a-z ]+) bags?`)

func readPuzzleInput() BagRules {
	input, err := util.ReadLinesFromFile("input.txt")
	util.CheckError(err, "error reading puzzle input file")

	rules := make(BagRules)

	for _, line := range input {
		rule := strings.Split(line, " bags contain ")

		color := Color(rule[0])
		var bagList []BagContents
		contents := ruleRegExp.FindAllStringSubmatch(rule[1], -1)

		for _, c := range contents {
			num, _ := strconv.Atoi(c[1])
			bagList = append(bagList, BagContents{num, Color(c[2])})
		}

		rules[color] = bagList
	}

	return rules
}

func canContain(container Color, bag Color, rules BagRules) bool {
	containedBags := rules[container]
	if len(containedBags) == 0 {
		return false
	}
	for _, b := range containedBags {
		if b.color == bag {
			return true
		} else if canContain(b.color, bag, rules) {
			return true
		}
	}
	return false
}

func countContainedBags(bag Color, rules BagRules) int {
	containedBags := rules[bag]
	if len(containedBags) == 0 {
		return 0
	}

	num := 0
	for _, b := range containedBags {
		num += b.num + b.num*countContainedBags(b.color, rules)
	}

	return num
}

func solvePartOne(rules BagRules) int {
	myBag := Color("shiny gold")

	num := 0
	for bag := range rules {
		if canContain(bag, myBag, rules) {
			num++
		}
	}

	return num
}

func solvePartTwo(rules BagRules) int {
	myBag := Color("shiny gold")
	return countContainedBags(myBag, rules)
}

func main() {
	fmt.Println("Advent of Code 2020 - Day ", day)

	rules := readPuzzleInput()

	fmt.Println("Part One: ", solvePartOne(rules))
	fmt.Println("Part Two: ", solvePartTwo(rules))
}
