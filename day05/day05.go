package main

import (
	"fmt"

	"github.com/pituser/aoc-2020-go/util"
)

const (
	day = "5"
)

func readPuzzleInput() []string {
	input, err := util.ReadLinesFromFile("input.txt")
	util.CheckError(err, "error reading puzzle input file")

	return input
}

func getRow(seatCode string) int {
	row := struct {
		low, up int
	}{
		0, 127,
	}

	for i := 0; i < 7; i++ {
		median := (row.up - row.low + 1) / 2
		if seatCode[i] == byte('F') {
			row.up -= median
		} else if seatCode[i] == byte('B') {
			row.low += median
		}
	}
	return row.up
}

func getCol(seatCode string) int {
	col := struct {
		low, up int
	}{
		0, 7,
	}

	for i := 7; i < 10; i++ {
		median := (col.up - col.low + 1) / 2
		if seatCode[i] == byte('L') {
			col.up -= median
		} else if seatCode[i] == byte('R') {
			col.low += median
		}
	}
	return col.up
}

func getSeatID(seatCode string) int {
	return getRow(seatCode)*8 + getCol(seatCode)
}

func solvePartOne(steatCodes []string) int {
	maxID := 0

	for _, seatCode := range steatCodes {
		id := getSeatID(seatCode)
		if id > maxID {
			maxID = id
		}
	}

	return maxID
}

func solvePartTwo(seatCodes []string) int {
	var seats [128 * 8]bool

	for _, seatCode := range seatCodes {
		seats[getSeatID(seatCode)] = true
	}

	for id := 0; id < 128*8; id++ {
		if seats[id] == false {
			if id > 0 && id < 128*8 && seats[id-1] == true && seats[id+1] == true {
				return id
			}
		}
	}

	return -1
}

func main() {
	fmt.Println("Advent of Code 2020 - Day ", day)

	input := readPuzzleInput()

	fmt.Println("Part One: ", solvePartOne(input))
	fmt.Println("Part Two: ", solvePartTwo(input))
}
