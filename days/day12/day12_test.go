package day12

import (
	"os"
	"strings"
	"testing"
)

const testInput = `0:
###
##.
##.

1:
###
##.
.##

2:
.##
###
##.

3:
##.
###
##.

4:
###
#..
###

5:
###
.#.
###

4x4: 0 0 0 0 2 0
12x5: 1 0 1 0 2 2
12x5: 1 0 1 0 3 2
`

func readInput() string {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	return string(data)
}

func TestDay12Part1(t *testing.T) {
	input := strings.NewReader(readInput())
	solution := SolvePart1(ParseInput(input))

	answer := 437
	if solution != answer {
		t.Errorf("Part 1: expected %d, got %d", answer, solution)
	}
}

func BenchmarkDay12Part1(b *testing.B) {
	input := readInput()

	for b.Loop() {
		SolvePart1(ParseInput(strings.NewReader(input)))
	}
}
