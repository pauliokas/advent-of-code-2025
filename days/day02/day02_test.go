package day02

import (
	"os"
	"strings"
	"testing"
)

const testInput = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,
1698522-1698528,446443-446449,38593856-38593862,565653-565659,
824824821-824824827,2121212118-2121212124`

func readInput() string {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	return string(data)
}

func TestDay02Part1Example(t *testing.T) {
	input := strings.NewReader(testInput)
	solution := SolvePart1(ParseInput(input))

	answer := 1227775554
	if solution != answer {
		t.Errorf("Part 1 example: expected %d, got %d", answer, solution)
	}
}

func TestDay02Part1(t *testing.T) {
	input := strings.NewReader(readInput())
	solution := SolvePart1(ParseInput(input))

	answer := 13919717792
	if solution != answer {
		t.Errorf("Part 1: expected %d, got %d", answer, solution)
	}
}

func TestDay02Part2Example(t *testing.T) {
	input := strings.NewReader(testInput)
	solution := SolvePart2(ParseInput(input))

	answer := 4174379265
	if solution != answer {
		t.Errorf("Part 2 example: expected %d, got %d", answer, solution)
	}
}

func TestDay02Part2(t *testing.T) {
	input := strings.NewReader(readInput())
	solution := SolvePart2(ParseInput(input))

	answer := 14582313461
	if solution != answer {
		t.Errorf("Part 2: expected %d, got %d", answer, solution)
	}
}

func BenchmarkDay02Part1(b *testing.B) {
	input := readInput()

	for b.Loop() {
		SolvePart1(ParseInput(strings.NewReader(input)))
	}
}

func BenchmarkDay02Part2(b *testing.B) {
	input := readInput()

	for b.Loop() {
		SolvePart2(ParseInput(strings.NewReader(input)))
	}
}
