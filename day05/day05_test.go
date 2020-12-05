package main

import (
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	var want int = 930
	got := solvePartOne(readPuzzleInput())

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSolvePartTwo(t *testing.T) {
	var want int = 515
	got := solvePartTwo(readPuzzleInput())

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestNewSeatId(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{"FBFBBFFRLR", 357},
		{"FFFBBBFRRR", 119},
		{"BBFFBBFRLL", 820},
	}

	for _, test := range tests {
		t.Run("newSeat", func(t *testing.T) {
			got := newSeat(test.input).id
			if got != test.want {
				t.Errorf("got %v want %v", got, test.want)
			}
		})

	}
}
