package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/pituser/aoc-2020-go/util"
)

const (
	day = "13"
)

func readPuzzleInput() (depart int, busIDs []int) {
	input, err := util.ReadLinesFromFile("input.txt")
	util.CheckError(err, "error reading puzzle input file")

	depart, err = strconv.Atoi(input[0])
	util.CheckError(err, "parse error")

	ids := strings.Split(input[1], ",")
	busIDs = make([]int, len(ids))
	for i, id := range ids {
		if id != "x" {
			busIDs[i], err = strconv.Atoi(id)
			util.CheckError(err, "parse error")
		} else {
			busIDs[i] = -1
		}

	}
	return depart, busIDs
}

func solvePartOne(depart int, busIDs []int) int {
	minWaitTime := math.MaxInt32
	minWaitID := -1
	for i := 0; i < len(busIDs); i++ {
		if busIDs[i] != -1 {
			waitTime := int(math.Ceil(float64(depart)/float64(busIDs[i])))*busIDs[i] - depart
			if waitTime < minWaitTime {
				minWaitTime = waitTime
				minWaitID = busIDs[i]
			}
		}
	}

	return minWaitID * int(minWaitTime)
}

func solvePartTwo(busIDs []int) uint64 {
	first := true
	fmt.Println("Let Wolfram alpha solve this: ")
	for i := 0; i < len(busIDs); i++ {
		if busIDs[i] != -1 {
			if !first {
				fmt.Print(", ")
			} else {
				first = false
			}

			fmt.Printf("(n+%d)%%%d = 0", i, busIDs[i])
		}
	}
	fmt.Println()
	// Result will be
	// n = 581610429053251 m + 560214575859998, m element Z

	return 560214575859998
}

func main() {
	fmt.Println("Advent of Code 2020 - Day ", day)

	depart, busIDs := readPuzzleInput()

	fmt.Println("Part One: ", solvePartOne(depart, busIDs))
	fmt.Println("Part Two: ", solvePartTwo(busIDs))
}
