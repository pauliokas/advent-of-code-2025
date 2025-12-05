package day04

import (
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

func readInput() string {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	return string(data)
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
	input := strings.NewReader(readInput())
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
	input := strings.NewReader(readInput())
	solution := SolvePart2(ParseInput(input))

	answer := 9194
	if solution != answer {
		t.Errorf("Part 2: expected %d, got %d", answer, solution)
	}
}

func BenchmarkDay04Part1(b *testing.B) {
	input := readInput()

	for b.Loop() {
		SolvePart1(ParseInput(strings.NewReader(input)))
	}
}

func BenchmarkDay04Part2(b *testing.B) {
	input := readInput()

	for b.Loop() {
		SolvePart2(ParseInput(strings.NewReader(input)))
	}
}
