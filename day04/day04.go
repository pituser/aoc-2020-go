package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pituser/aoc-2020-go/util"
)

const (
	day = "4"
)

type passport map[string]string

func readPuzzleInput() []passport {
	input, err := util.ReadLinesFromFile("input.txt")
	util.CheckError(err, "error reading puzzle input file")

	var passports []passport
	pass := make(passport)

	for _, line := range input {
		if line != "" {
			fields := strings.Split(line, " ")
			for _, field := range fields {
				data := strings.Split(field, ":")
				pass[data[0]] = data[1]
			}
		} else {
			passports = append(passports, pass)
			pass = make(passport)
		}
	}
	passports = append(passports, pass)
	return passports
}

func hasField(pass passport, field string) bool {
	_, ok := pass[field]
	return ok
}

func checkPassportPartOne(pass passport) bool {
	fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, field := range fields {
		if !hasField(pass, field) {
			return false
		}
	}
	return true
}

func checkInt(s string, min, max int) bool {
	val, err := strconv.Atoi(s)
	return err == nil && val >= min && val <= max
}

var heightRE *regexp.Regexp = regexp.MustCompile(`^(\d+)(in|cm)$`)

func checkHeight(s string) bool {
	parts := heightRE.FindStringSubmatch(s)
	if len(parts) != 3 {
		return false
	}

	unit := parts[2]
	if unit == "cm" {
		val, err := strconv.Atoi(parts[1])
		return err == nil && val >= 150 && val <= 193
	}
	if unit == "in" {
		val, err := strconv.Atoi(parts[1])
		return err == nil && val >= 59 && val <= 76
	}
	return false
}

var hairRE *regexp.Regexp = regexp.MustCompile(`^#[0-9a-f]{6}$`)

func checkHairColor(s string) bool {
	return hairRE.MatchString(s)
}

func checkEyeColor(s string) bool {
	colors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

	for _, color := range colors {
		if s == color {
			return true
		}
	}

	return false
}

var pidRE *regexp.Regexp = regexp.MustCompile(`^[0-9]{9}$`)

func checkPid(s string) bool {
	return pidRE.MatchString(s)
}

var f func(int) int = func(i int) int { return 12 * i }

var checker = map[string](func(string) bool){
	"byr": func(s string) bool { return checkInt(s, 1920, 2002) },
	"iyr": func(s string) bool { return checkInt(s, 2010, 2020) },
	"eyr": func(s string) bool { return checkInt(s, 2020, 2030) },
	"hgt": checkHeight,
	"hcl": checkHairColor,
	"ecl": checkEyeColor,
	"pid": checkPid,
}

func checkPassportPartTwo(pass passport) bool {
	fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, field := range fields {
		val, ok := pass[field]
		if !(ok && checker[field](val)) {
			return false
		}
	}
	return true
}

func checkPassports(passports []passport, check func(passport) bool) int {
	cnt := 0
	for _, pass := range passports {
		if check(pass) {
			cnt++
		}
	}
	return cnt
}

func solvePartOne(passports []passport) int {
	return checkPassports(passports, checkPassportPartOne)
}

func solvePartTwo(passports []passport) int {
	return checkPassports(passports, checkPassportPartTwo)
}

func main() {
	fmt.Println("Advent of Code 2020 - Day ", day)

	input := readPuzzleInput()

	fmt.Println("Part One: ", solvePartOne(input))
	fmt.Println("Part Two: ", solvePartTwo(input))
}
