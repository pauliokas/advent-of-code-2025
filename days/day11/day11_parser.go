package day11

import (
	"bufio"
	"io"
	"strings"
)

type Grid map[string](map[string]bool)

type Input Grid

func ParseInput(input io.Reader) Input {
	scanner := bufio.NewScanner(input)

	grid := make(Grid)
	for scanner.Scan() {
		line := scanner.Text()

		separatorIdx := strings.IndexRune(line, ':')
		vertex := line[:separatorIdx]
		edgesStr := line[separatorIdx+2:]

		edges := make(map[string]bool)
		for edge := range strings.SplitSeq(edgesStr, " ") {
			edges[edge] = true
		}

		grid[vertex] = edges
	}

	return Input(grid)
}
