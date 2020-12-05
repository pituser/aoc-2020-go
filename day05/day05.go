package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/pituser/aoc-2020-go/util"
)

const (
	day = "5"
)

func readPuzzleInput() []Seat {
	input, err := util.ReadLinesFromFile("input.txt")
	util.CheckError(err, "error reading puzzle input file")
	var seats []Seat
	for _, line := range input {
		seat := newSeat(line)
		if seat == nil {
			util.Error("error parsing seat codes")
		} else {
			seats = append(seats, *seat)
		}
	}

	return seats
}

const (
	lower = iota
	upper
)

type Seat struct {
	code     [10]byte
	row, col int
	id       int
}

func binSearch(codes []byte) int {
	l := len(codes)
	r := struct{ low, up int }{0, int(math.Exp2(float64(l)))}

	for i := 0; i < l; i++ {
		median := (r.up - r.low + 1) / 2
		if codes[i] == lower {
			r.up -= median
		} else {
			r.low += median
		}
	}
	return r.low
}

var seatRexExp *regexp.Regexp = regexp.MustCompile(`^[FB]{7}[LR]{3}$`)

func newSeat(inputCode string) *Seat {
	var code [10]byte
	inputCode = strings.TrimSpace(inputCode)
	if !seatRexExp.MatchString(inputCode) {
		return nil
	}
	for i := 0; i < len(inputCode); i++ {
		switch inputCode[i] {
		case byte('F'), byte('L'):
			code[i] = lower
		default:
			code[i] = upper
		}
	}
	var seat Seat
	seat.code = code
	seat.row = binSearch(code[0:7])
	seat.col = binSearch(code[7:10])
	seat.id = 8*seat.row + seat.col
	return &seat
}

func solvePartOne(seats []Seat) int {
	maxID := 0

	for _, seat := range seats {
		id := seat.id
		if id > maxID {
			maxID = id
		}
	}

	return maxID
}

func solvePartTwo(seats []Seat) int {
	const maxSeats = 128 * 8
	var seatMap [maxSeats]bool

	for _, seat := range seats {
		seatMap[seat.id] = true
	}

	for id := 0; id < maxSeats; id++ {
		if seatMap[id] == false {
			if id > 0 && id < maxSeats && seatMap[id-1] && seatMap[id+1] {
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
