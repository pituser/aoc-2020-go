package main

import (
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	var want int = 2164
	got := solvePartOne(readPuzzleInput())

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSolvePartTwo(t *testing.T) {
	var want int = 1974
	got := solvePartTwo(readPuzzleInput())

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
