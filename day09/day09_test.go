package main

import (
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	var want int = 552655238
	got := solvePartOne(readPuzzleInput())

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSolvePartTwo(t *testing.T) {
	var want int = 70672245
	got := solvePartTwo(readPuzzleInput(), 552655238)

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
