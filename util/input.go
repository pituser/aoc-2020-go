package util

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// ReadLinesFromFile reads a given file and returns all lines as string slice.
func ReadLinesFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

// ParseIntList takes a list of strings and makes a list of integers.
func ParseIntList(lines []string) ([]int, error) {
	var result []int
	for _, line := range lines {
		value, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			return make([]int, 0), err
		}
		result = append(result, int(value))
	}
	return result, nil
}
