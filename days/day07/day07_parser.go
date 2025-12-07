package day07

import (
	"bufio"
	"io"
)

type Input struct {
	Grid  Grid
	Start Coords
}

func ParseInput(input io.Reader) Input {
	scanner := bufio.NewScanner(input)

	grid := make(map[int]rune)
	width := 0
	var start *Coords

	rowIdx := 0
	for scanner.Scan() {
		line := scanner.Text()
		width = len(line)

		for colIdx, char := range line {
			switch char {
			case 'S':
				start = &Coords{X: colIdx, Y: rowIdx}
				continue
			case '.':
				continue
			}
			grid[pair(colIdx, rowIdx)] = char
		}
		rowIdx += 1
	}

	return Input{Grid: Grid{Width: width, Height: rowIdx, data: grid}, Start: *start}
}
