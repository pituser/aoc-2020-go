package main

import (
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	var want uint64 = 7440382076205
	got := solvePartOne(readPuzzleInput())

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSolvePartTwo(t *testing.T) {
	var want uint64 = 4200656704538
	got := solvePartTwo(readPuzzleInput())

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
