package day02

import (
	"io"
	"strconv"
	"strings"
)

type Range struct {
	Min int
	Max int
}

type Input []Range

func ParseInput(input io.Reader) Input {
	inputBytes, _ := io.ReadAll(input)

	splitRanges := strings.Split(strings.ReplaceAll(string(inputBytes), "\n", ""), ",")

	ranges := make([]Range, 0, len(splitRanges))

	for _, r := range splitRanges {
		bounds := strings.Split(r, "-")
		min, _ := strconv.Atoi(bounds[0])
		max, _ := strconv.Atoi(bounds[1])
		ranges = append(ranges, Range{Min: min, Max: max})
	}
	return ranges
}
