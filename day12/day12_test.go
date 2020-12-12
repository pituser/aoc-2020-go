package main

import (
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	var want int = 2297
	got := solvePartOne(readPuzzleInput())

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSolvePartTwo(t *testing.T) {
	var want int = 89984
	got := solvePartTwo(readPuzzleInput())

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
