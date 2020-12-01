package main

import (
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	var want int = 42
	got := solvePartOne()

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSolvePartTwo(t *testing.T) {
	var want int = -1
	got := solvePartTwo()

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
