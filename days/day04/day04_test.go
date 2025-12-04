package day04

import (
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

const testInput = `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

func readInput() io.ReadCloser {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func TestDay04Part1Example(t *testing.T) {
	input := strings.NewReader(testInput)
	solution := SolvePart1(ParseInput(input))

	answer := 13
	if solution != answer {
		t.Errorf("Part 1 example: expected %d, got %d", answer, solution)
	}
}

func TestDay04Part1(t *testing.T) {
	input := readInput()
	defer input.Close()
	solution := SolvePart1(ParseInput(input))

	answer := 1493
	if solution != answer {
		t.Errorf("Part 1: expected %d, got %d", answer, solution)
	}
}

func TestDay04Part2Example(t *testing.T) {
	input := strings.NewReader(testInput)
	solution := SolvePart2(ParseInput(input))

	answer := 43
	if solution != answer {
		t.Errorf("Part 2 example: expected %d, got %d", answer, solution)
	}
}

func TestDay04Part2(t *testing.T) {
	input := readInput()
	defer input.Close()
	solution := SolvePart2(ParseInput(input))

	answer := 9194
	if solution != answer {
		t.Errorf("Part 2: expected %d, got %d", answer, solution)
	}
}
