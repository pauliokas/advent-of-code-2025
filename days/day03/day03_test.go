package day03

import (
	"os"
	"strings"
	"testing"
)

const testInput = `987654321111111
811111111111119
234234234234278
818181911112111`

func readInput() string {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	return string(data)
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
	input := strings.NewReader(readInput())
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
	input := strings.NewReader(readInput())
	solution := SolvePart2(ParseInput(input))

	answer := 167302518850275
	if solution != answer {
		t.Errorf("Part 2: expected %d, got %d", answer, solution)
	}
}

func BenchmarkDay03Part1(b *testing.B) {
	input := readInput()

	for b.Loop() {
		SolvePart1(ParseInput(strings.NewReader(input)))
	}
}

func BenchmarkDay03Part2(b *testing.B) {
	input := readInput()

	for b.Loop() {
		SolvePart2(ParseInput(strings.NewReader(input)))
	}
}
