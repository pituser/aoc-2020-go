package main

import (
	"fmt"
	"sort"

	"github.com/pituser/aoc-2020-go/util"
)

const (
	day = "XX"
)

func readPuzzleInput() []int {
	input, err := util.ReadLinesFromFile("input.txt")
	util.CheckError(err, "error reading puzzle input file")

	list, err := util.ParseIntList(input)
	util.CheckError(err, "parse error")
	return list
}

func solvePartOne(joltages []int) int {
	joltages = append(joltages, 0)                       // outlet
	joltages = append(joltages, util.MaxInt(joltages)+3) // device

	sort.Ints(joltages)

	diff1, diff3 := 0, 0

	for i := 1; i < len(joltages); i++ {
		diff := joltages[i] - joltages[i-1]
		if diff == 1 {
			diff1++
		} else if diff == 3 {
			diff3++
		} else {
			panic("unexpected diff")
		}
	}

	return diff1 * diff3
}

func countWays(joltMap map[int]int, cache []uint64, current, max int) (ways uint64) {
	if cache[current] != 0 {
		return cache[current]
	}
	if joltMap[current] == max {
		cache[current] = uint64(1)
		return uint64(1)
	}
	for k := 1; k <= 3; k++ {
		if joltMap[current+k] != 0 {
			ways += countWays(joltMap, cache, current+k, max)
		}
	}
	cache[current] = ways

	return ways
}

func solvePartTwo(joltages []int) uint64 {
	joltages = append(joltages, 0)                       // outlet
	joltages = append(joltages, util.MaxInt(joltages)+3) // device

	sort.Ints(joltages)
	last := len(joltages) - 1
	max := joltages[last]

	joltMap := make(map[int]int)

	for i := 0; i <= last; i++ {
		joltMap[joltages[i]] = joltages[i]
	}
	cache := make([]uint64, max+1)
	return countWays(joltMap, cache, 0, max)
}

func main() {
	fmt.Println("Advent of Code 2020 - Day ", day)

	input := readPuzzleInput()

	fmt.Println("Part One: ", solvePartOne(input))
	fmt.Println("Part Two: ", solvePartTwo(input))
}
