package main

import (
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	var want int = 2030
	got := solvePartOne(readPuzzleInput())

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSolvePartTwo(t *testing.T) {
	var want uint64 = 42313823813632
	got := solvePartTwo(readPuzzleInput())

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
