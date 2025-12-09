package day08

import (
	"cmp"
	"slices"
)

type Line struct {
	A, B Coords
}

func getDistance(a, b Coords) int {
	diffX := a.X - b.X
	diffY := a.Y - b.Y
	diffZ := a.Z - b.Z
	return diffX*diffX + diffY*diffY + diffZ*diffZ
}

func getSortedLines(coords []Coords) []Line {
	lines := make([]Line, 0, (len(coords)*(len(coords)+1))/2)
	for i := 0; i < len(coords); i++ {
		for j := i + 1; j < len(coords); j++ {
			lines = append(lines, Line{A: coords[i], B: coords[j]})
		}
	}

	slices.SortFunc(lines, func(a, b Line) int { return cmp.Compare(getDistance(a.A, a.B), getDistance(b.A, b.B)) })

	return lines
}

func SolvePart1(input Input, connections int) int {
	chains := make(map[Coords]*[]Coords, len(input))
	for _, coord := range input {
		chain := make([]Coords, 0)
		chain = append(chain, coord)
		chains[coord] = &chain
	}

	for _, minLine := range getSortedLines(input)[:connections] {
		if chains[minLine.A] == chains[minLine.B] {
			continue
		}

		newChain := append(*chains[minLine.A], *chains[minLine.B]...)
		for _, coord := range newChain {
			chains[coord] = &newChain
		}
	}

	distinctChainsMap := make(map[*[]Coords]bool)
	for _, chain := range chains {
		distinctChainsMap[chain] = true
	}

	distinctChains := make([][]Coords, 0, len(distinctChainsMap))
	for chain := range distinctChainsMap {
		distinctChains = append(distinctChains, *chain)
	}
	slices.SortFunc(distinctChains, func(a, b []Coords) int { return len(b) - len(a) })

	product := 1
	for _, chain := range distinctChains[:3] {
		product *= len(chain)
	}

	return product
}

func SolvePart2(input Input) int {
	chains := make(map[Coords]*[]Coords, len(input))
	for _, coord := range input {
		chain := make([]Coords, 0)
		chain = append(chain, coord)
		chains[coord] = &chain
	}

	for _, minLine := range getSortedLines(input) {
		if chains[minLine.A] == chains[minLine.B] {
			continue
		}

		newChain := append(*chains[minLine.A], *chains[minLine.B]...)

		if len(newChain) == len(input) {
			return minLine.A.X * minLine.B.X
		}

		for _, coord := range newChain {
			chains[coord] = &newChain
		}

	}

	return -1
}
