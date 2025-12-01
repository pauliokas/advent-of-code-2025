package day01

import (
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

const testInput = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

func readInput() io.ReadCloser {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func TestDay01Part1Example(t *testing.T) {
	input := strings.NewReader(testInput)
	solution := SolvePart1(ParseInput(input))

	answer := 3
	if solution != answer {
		t.Errorf("Part 1 example: expected %d, got %d", answer, solution)
	}
}

func TestDay01Part1(t *testing.T) {
	input := readInput()
	defer input.Close()
	solution := SolvePart1(ParseInput(input))

	answer := 1123
	if solution != answer {
		t.Errorf("Part 1: expected %d, got %d", answer, solution)
	}
}

func TestDay01Part2Example(t *testing.T) {
	input := strings.NewReader(testInput)
	solution := SolvePart2(ParseInput(input))

	answer := 6
	if solution != answer {
		t.Errorf("Part 2 example: expected %d, got %d", answer, solution)
	}
}

func TestDay01Part2(t *testing.T) {
	input := readInput()
	defer input.Close()
	solution := SolvePart2(ParseInput(input))

	answer := 6695
	if solution != answer {
		t.Errorf("Part 2: expected %d, got %d", answer, solution)
	}
}
