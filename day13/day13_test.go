package main

import (
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	var want int = 2238
	got := solvePartOne(readPuzzleInput())

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSolvePartTwo(t *testing.T) {
	var want uint64 = 560214575859998
	_, busIDs := readPuzzleInput()
	got := solvePartTwo(busIDs)

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
