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

func printSeats(seats [][]int) {
	for i := 0; i < len(seats); i++ {
		for j := 0; j < len(seats[0]); j++ {
			switch seats[i][j] {
			case Floor:
				fmt.Print(".")
			case Empty:
				fmt.Print("L")
			case Occ:
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func countAdjacentOccupied(seats [][]int, row, col int) (numOcc int) {
	h, w := len(seats), len(seats[0])

	if row > 0 {
		if seats[row-1][col] == Occ {
			numOcc++
		}
		if col > 0 {
			if seats[row-1][col-1] == Occ {
				numOcc++
			}
		}
		if col < w-1 {
			if seats[row-1][col+1] == Occ {
				numOcc++
			}
		}
	}
	if row < h-1 {
		if seats[row+1][col] == Occ {
			numOcc++
		}
		if col > 0 {
			if seats[row+1][col-1] == Occ {
				numOcc++
			}
		}
		if col < w-1 {
			if seats[row+1][col+1] == Occ {
				numOcc++
			}
		}
	}
	if col > 0 {
		if seats[row][col-1] == Occ {
			numOcc++
		}
	}
	if col < w-1 {
		if seats[row][col+1] == Occ {
			numOcc++
		}
	}
	return numOcc
}

func countVisibleOccupied(seats [][]int, row, col int) (numOcc int) {
	h, w := len(seats), len(seats[0])

	// left up
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if seats[i][j] == Occ {
			numOcc++
			break
		}
		if seats[i][j] == Empty {
			break
		}
	}
	// up
	for i := row - 1; i >= 0; i = i - 1 {
		if seats[i][col] == Occ {
			numOcc++
			break
		}
		if seats[i][col] == Empty {
			break
		}
	}
	// right up
	for i, j := row-1, col+1; i >= 0 && j < w; i, j = i-1, j+1 {
		if seats[i][j] == Occ {
			numOcc++
			break
		}
		if seats[i][j] == Empty {
			break
		}
	}
	// left
	for j := col - 1; j >= 0; j = j - 1 {
		if seats[row][j] == Occ {
			numOcc++
			break
		}
		if seats[row][j] == Empty {
			break
		}
	}
	// right
	for j := col + 1; j < w; j = j + 1 {
		if seats[row][j] == Occ {
			numOcc++
			break
		}
		if seats[row][j] == Empty {
			break
		}
	}
	// left down
	for i, j := row+1, col-1; i < h && j >= 0; i, j = i+1, j-1 {
		if seats[i][j] == Occ {
			numOcc++
			break
		}
		if seats[i][j] == Empty {
			break
		}
	}
	// down
	for i := row + 1; i < h; i = i + 1 {
		if seats[i][col] == Occ {
			numOcc++
			break
		}
		if seats[i][col] == Empty {
			break
		}
	}
	// right down
	for i, j := row+1, col+1; i < h && j < w; i, j = i+1, j+1 {
		if seats[i][j] == Occ {
			numOcc++
			break
		}
		if seats[i][j] == Empty {
			break
		}
	}

	return numOcc
}

func simulateStep(seats [][]int, occCounter func([][]int, int, int) int, maxOcc int) (occupiedSeats int) {
	newSeats := make([][]int, len(seats))
	for i := 0; i < len(seats); i++ {
		newSeats[i] = make([]int, len(seats[0]))
	}

	for i, row := range seats {
		for j, pos := range row {
			if pos == Floor {
				newSeats[i][j] = Floor
			} else {
				numOcc := occCounter(seats, i, j)
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
		newOcc := simulateStep(seats, countAdjacentOccupied, 4)
		//fmt.Println(seats)
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
	for changed {
		newOcc := simulateStep(seats, countVisibleOccupied, 5)
		//printSeats(seats)
		if newOcc == occupiedSeats {
			changed = false
		}
		occupiedSeats = newOcc
	}
	return occupiedSeats
}

func main() {
	fmt.Println("Advent of Code 2020 - Day ", day)

	/*
		input := readPuzzleInput()
		printSeats(input)
		fmt.Println("Part One: ", solvePartOne(input))
	*/

	input := readPuzzleInput()
	printSeats(input)
	fmt.Println("Part Two: ", solvePartTwo(input))
}
