package day07

import "fmt"

type Coords struct {
	X, Y int
}

func (c Coords) String() string {
	return fmt.Sprintf("(%d, %d)", c.X, c.Y)
}

type Grid struct {
	Width, Height int
	data          map[int]rune
}

func (g *Grid) At(coords Coords) rune {
	if coords.X < 0 || coords.Y < 0 || coords.X >= g.Width || coords.Y >= g.Height {
		return '.'
	}

	v, ok := g.data[pair(coords.X, coords.Y)]
	if !ok {
		return '.'
	}
	return v
}

func (g *Grid) String() string {
	result := ""
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			result += string(g.At(Coords{x, y}))
		}
		result += "\n"
	}
	return result
}

func pair(x, y int) int {
	return ((x+y)*(x+y+1))/2 + y
}
