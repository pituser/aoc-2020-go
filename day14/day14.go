package main

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"

	"github.com/pituser/aoc-2020-go/util"
)

const (
	day = "14"
)

const (
	opUpdateMask = iota
	opWriteMem
)

type instruction struct {
	opCode int
	mask   string
	addr   uint64
	value  uint64
}

func readPuzzleInput() []instruction {
	input, err := util.ReadLinesFromFile("input.txt")
	util.CheckError(err, "error reading puzzle input file")

	var instructions []instruction

	for _, line := range input {
		var instr instruction
		parts := strings.Split(line, " = ")
		if parts[0] == "mask" {
			instr.opCode = opUpdateMask
			instr.mask = parts[1]
		} else {
			instr.opCode = opWriteMem
			addr, err := strconv.Atoi(parts[0][4 : len(parts[0])-1])
			util.CheckError(err, "parse error")
			instr.addr = uint64(addr)
			value, err := strconv.Atoi((parts[1]))
			util.CheckError(err, "parse error")
			instr.value = uint64(value)
		}

		instructions = append(instructions, instr)
	}

	return instructions
}

func masksPartOne(mask string) (andMask, orMask uint64) {
	for i := 0; i < len(mask); i++ {
		switch mask[i] {
		case byte('1'):
			andMask++
			orMask++
		case byte('X'):
			andMask++
		}
		andMask = andMask << 1
		orMask = orMask << 1
	}

	return andMask >> 1, orMask >> 1
}

func solvePartOne(instructions []instruction) uint64 {
	mem := make(map[uint64]uint64)

	var andMask, orMask uint64
	for _, instr := range instructions {
		if instr.opCode == opUpdateMask {
			andMask, orMask = masksPartOne(instr.mask)
		} else {
			value := (instr.value & andMask) | orMask
			mem[instr.addr] = value
		}
	}
	sum := uint64(0)
	for _, v := range mem {
		sum += v
	}
	return sum
}

func floatingAddresses(address uint64, mask string) (adresses []uint64) {
	var addresses []uint64

	var masks []string
	work := list.New()
	work.PushBack(mask)

	for work.Len() > 0 {
		currentMask := work.Remove(work.Back()).(string)
		foundFloating := false
		for i := 0; i < len(currentMask); i++ {
			if currentMask[i] == byte('X') {
				var b []byte
				b = make([]byte, len(currentMask))
				copy(b, currentMask)
				b[i] = byte('z')
				work.PushBack(string(b))
				b = make([]byte, len(currentMask))
				copy(b, currentMask)
				b[i] = byte('x')
				work.PushBack(string(b))
				foundFloating = true
				break
			}
		}
		if !foundFloating {
			masks = append(masks, currentMask)
		}
	}

	for _, m := range masks {
		var andMask, orMask uint64
		for i := 0; i < len(m); i++ {
			switch m[i] {
			case byte('0'):
				andMask++
			case byte('1'), byte('x'):
				andMask++
				orMask++
			}
			orMask <<= 1
			andMask <<= 1
		}
		orMask >>= 1
		andMask >>= 1

		addresses = append(addresses, (address&andMask)|orMask)
	}

	return addresses
}

func solvePartTwo(instructions []instruction) uint64 {
	mem := make(map[uint64]uint64)

	var currentMask string
	for _, instr := range instructions {
		if instr.opCode == opUpdateMask {
			currentMask = instr.mask
		} else {
			for _, addr := range floatingAddresses(instr.addr, currentMask) {
				mem[addr] = instr.value
			}
		}
	}

	sum := uint64(0)
	for _, v := range mem {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println("Advent of Code 2020 - Day ", day)

	input := readPuzzleInput()

	fmt.Println("Part One: ", solvePartOne(input))
	fmt.Println("Part Two: ", solvePartTwo(input))
}
