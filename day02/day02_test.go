package main

import (
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	var want int = 628
	got := solvePartOne(readPuzzleInput())

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSolvePartTwo(t *testing.T) {
	var want int = 705
	got := solvePartTwo(readPuzzleInput())

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestCheckPassword(t *testing.T) {
	var tests = []struct {
		input                  string
		isCorrect1, isCorrect2 bool
	}{
		{"1-3 a: abcde", true, true},
		{"1-3 b: cdefg", false, false},
		{"2-9 c: ccccccccc", true, false},
	}

	var inputLines []string
	for _, test := range tests {
		inputLines = append(inputLines, test.input)
	}
	passwords := parsePasswordLines(inputLines)

	for i, test := range tests {
		t.Run("checkPassword1", func(t *testing.T) {
			got := checkPasswordPart1(passwords[i])
			if got != test.isCorrect1 {
				t.Errorf("got %v want %v", got, test.isCorrect1)
			}
		})
		t.Run("checkPassword2", func(t *testing.T) {
			got := checkPasswordPart2(passwords[i])
			if got != test.isCorrect2 {
				t.Errorf("got %v want %v", got, test.isCorrect2)
			}
		})
	}
}
