package day04

import (
	"bufio"
	"io"
)

type Coords struct {
	X, Y int
}

type Grid struct {
	Width, Height int
	data          map[int]rune
}

func (g Grid) At(coords Coords) rune {
	if coords.X < 0 || coords.Y < 0 || coords.X >= g.Width || coords.Y >= g.Height {
		return '.'
	}

	v, ok := g.data[pair(coords.X, coords.Y)]
	if !ok {
		return '.'
	}
	return v
}

func (g Grid) Remove(coords Coords) {
	if g.At(coords) == '.' {
		return
	}
	delete(g.data, pair(coords.X, coords.Y))
}

func (g Grid) String() string {
	result := ""
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			result += string(g.At(Coords{x, y}))
		}
		result += "\n"
	}
	return result
}

type Input Grid

func pair(x, y int) int {
	return ((x+y)*(x+y+1))/2 + y
}

func ParseInput(input io.Reader) Input {
	scanner := bufio.NewScanner(input)
	data := make(map[int]rune)

	lineIdx := 0
	width := 0
	for scanner.Scan() {
		line := scanner.Text()

		width = len(line)

		for x := 0; x < len(line); x++ {
			if line[x] == '.' {
				continue
			}
			data[pair(x, lineIdx)] = rune(line[x])
		}
		lineIdx += 1
	}

	return Input{width, lineIdx, data}
}
