package day04

var neighbours = []Coords{
	{-1, -1},
	{0, -1},
	{1, -1},
	{-1, 0},
	{1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
}

func SolvePart1(input Input) int {
	grid := Grid(input)

	accessibleRolls := 0
	for x := 0; x < grid.Width; x++ {
		for y := 0; y < grid.Height; y++ {
			if grid.At(Coords{x, y}) != '@' {
				continue
			}

			neighouringRolls := 0

			for _, neighbour := range neighbours {
				if grid.At(Coords{x + neighbour.X, y + neighbour.Y}) == '@' {
					neighouringRolls++
				}
			}

			if neighouringRolls < 4 {
				accessibleRolls++
			}
		}
	}

	return accessibleRolls
}

func SolvePart2(input Input) int {
	grid := Grid(input)

	removableRolls := 0

	for {
		var toRemove []Coords = make([]Coords, 0)

		for x := 0; x < grid.Width; x++ {
			for y := 0; y < grid.Height; y++ {
				if grid.At(Coords{x, y}) != '@' {
					continue
				}

				neighouringRolls := 0

				for _, neighbour := range neighbours {
					if grid.At(Coords{x + neighbour.X, y + neighbour.Y}) == '@' {
						neighouringRolls++
					}
				}

				if neighouringRolls < 4 {
					toRemove = append(toRemove, Coords{x, y})
				}
			}
		}

		if len(toRemove) == 0 {
			break
		}

		removableRolls += len(toRemove)
		for _, coords := range toRemove {
			grid.Remove(coords)
		}
	}

	return removableRolls
}
