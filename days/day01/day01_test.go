package day01

import (
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

func readInput() string {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	return string(data)
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
	input := strings.NewReader(readInput())
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
	input := strings.NewReader(readInput())
	solution := SolvePart2(ParseInput(input))

	answer := 6695
	if solution != answer {
		t.Errorf("Part 2: expected %d, got %d", answer, solution)
	}
}

func BenchmarkDay01Part1(b *testing.B) {
	input := readInput()

	for b.Loop() {
		SolvePart1(ParseInput(strings.NewReader(input)))
	}
}

func BenchmarkDay01Part2(b *testing.B) {
	input := readInput()

	for b.Loop() {
		SolvePart2(ParseInput(strings.NewReader(input)))
	}
}
