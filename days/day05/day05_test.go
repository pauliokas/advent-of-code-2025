package day05

import (
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

const testInput = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

func readInput() io.ReadCloser {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func TestDay05Part1Example(t *testing.T) {
	input := strings.NewReader(testInput)
	solution := SolvePart1(ParseInput(input))

	answer := 3
	if solution != answer {
		t.Errorf("Part 1 example: expected %d, got %d", answer, solution)
	}
}

func TestDay05Part1(t *testing.T) {
	input := readInput()
	defer input.Close()
	solution := SolvePart1(ParseInput(input))

	answer := 640
	if solution != answer {
		t.Errorf("Part 1: expected %d, got %d", answer, solution)
	}
}

func TestDay05Part2Example(t *testing.T) {
	input := strings.NewReader(testInput)
	solution := SolvePart2(ParseInput(input))

	answer := 14
	if solution != answer {
		t.Errorf("Part 2 example: expected %d, got %d", answer, solution)
	}
}

func TestDay05Part2(t *testing.T) {
	input := readInput()
	defer input.Close()
	solution := SolvePart2(ParseInput(input))

	answer := 365804144481581
	if solution != answer {
		t.Errorf("Part 2: expected %d, got %d", answer, solution)
	}
}
