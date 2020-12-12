package main

import (
	"fmt"

	"github.com/pituser/aoc-2020-go/util"
)

const (
	day = "11"
)

const (
	Floor = iota
	Empty
	Occ
)

func readPuzzleInput() [][]int {
	input, err := util.ReadLinesFromFile("input.txt")
	util.CheckError(err, "error reading puzzle input file")

	var seats [][]int

	for _, line := range input {
		var row []int
		for j := 0; j < len(line); j++ {
			var code int
			switch line[j] {
			case byte('.'):
				code = Floor
			case byte('L'):
				code = Empty
			case byte('#'):
				code = Occ
			default:
				panic("unknown seat code")
			}
			row = append(row, code)
		}
		seats = append(seats, row)
	}

	return seats
}

func countOccupied(seats [][]int, row, col int, dist int) int {
	h, w := len(seats), len(seats[0])
	numOcc := 0
	dir := [3]int{-1, 0, 1}

	for _, dirY := range dir {
		for _, dirX := range dir {
			if !(dirX == 0 && dirY == 0) {
				for i, j, d := row+dirX, col+dirY, 0; i >= 0 && i < h && j >= 0 && j < w && d < dist; i, j, d = i+dirX, j+dirY, d+1 {
					if seats[i][j] == Occ {
						numOcc++
						break
					}
					if seats[i][j] == Empty {
						break
					}
				}
			}
		}
	}

	return numOcc
}

func simulateStep(seats [][]int, maxDist int, maxOcc int) (occupiedSeats int) {
	newSeats := make([][]int, len(seats))
	for i := 0; i < len(seats); i++ {
		newSeats[i] = make([]int, len(seats[0]))
	}

	for i, row := range seats {
		for j, pos := range row {
			if pos == Floor {
				newSeats[i][j] = Floor
			} else {
				numOcc := countOccupied(seats, i, j, maxDist)
				if pos == Empty {
					if numOcc == 0 {
						newSeats[i][j] = Occ
						occupiedSeats++
					} else {
						newSeats[i][j] = Empty
					}
				}
				if pos == Occ {
					if numOcc >= maxOcc {
						newSeats[i][j] = Empty
					} else {
						newSeats[i][j] = Occ
						occupiedSeats++
					}
				}
			}
		}
	}

	for i := 0; i < len(seats); i++ {
		copy(seats[i], newSeats[i])
	}
	return occupiedSeats
}

func solvePartOne(seats [][]int) int {
	changed := true
	occupiedSeats := 0
	for changed {
		newOcc := simulateStep(seats, 1, 4)
		if newOcc == occupiedSeats {
			changed = false
		}
		occupiedSeats = newOcc
	}
	return occupiedSeats
}

func solvePartTwo(seats [][]int) int {
	changed := true
	occupiedSeats := 0
	maxDist := util.MaxInt([]int{len(seats), len(seats[0])})
	for changed {
		newOcc := simulateStep(seats, maxDist, 5)
		if newOcc == occupiedSeats {
			changed = false
		}
		occupiedSeats = newOcc
	}
	return occupiedSeats
}

func main() {
	fmt.Println("Advent of Code 2020 - Day ", day)

	input := readPuzzleInput()
	fmt.Println("Part One: ", solvePartOne(input))

	input = readPuzzleInput()
	fmt.Println("Part Two: ", solvePartTwo(input))
}
