package day03

import (
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

const testInput = `987654321111111
811111111111119
234234234234278
818181911112111`

func readInput() io.ReadCloser {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func TestDay03Part1Example(t *testing.T) {
	input := strings.NewReader(testInput)
	solution := SolvePart1(ParseInput(input))

	answer := 357
	if solution != answer {
		t.Errorf("Part 1 example: expected %d, got %d", answer, solution)
	}
}

func TestDay03Part1(t *testing.T) {
	input := readInput()
	defer input.Close()
	solution := SolvePart1(ParseInput(input))

	answer := 16887
	if solution != answer {
		t.Errorf("Part 1: expected %d, got %d", answer, solution)
	}
}

func TestDay03Part2Example(t *testing.T) {
	input := strings.NewReader(testInput)
	solution := SolvePart2(ParseInput(input))

	answer := 3121910778619
	if solution != answer {
		t.Errorf("Part 2 example: expected %d, got %d", answer, solution)
	}
}

func TestDay03Part2(t *testing.T) {
	input := readInput()
	defer input.Close()
	solution := SolvePart2(ParseInput(input))

	answer := 167302518850275
	if solution != answer {
		t.Errorf("Part 2: expected %d, got %d", answer, solution)
	}
}
