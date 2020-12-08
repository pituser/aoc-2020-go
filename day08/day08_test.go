package main

import (
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	var want int = 1394
	got := solvePartOne(readPuzzleInput())

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSolvePartTwo(t *testing.T) {
	var want int = 1626
	got := solvePartTwo(readPuzzleInput())

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
