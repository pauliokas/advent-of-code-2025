package day09

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func SolvePart1(input Input) int {
	maxArea := 0
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			width := abs(input[i].X-input[j].X) + 1
			height := abs(input[i].Y-input[j].Y) + 1
			area := width * height

			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func SolvePart2(input Input) int {
	var boundsMax Coords

	for _, vertex := range input {
		if vertex.X > boundsMax.X {
			boundsMax.X = vertex.X
		}
		if vertex.Y > boundsMax.Y {
			boundsMax.Y = vertex.Y
		}
	}

	grid := make([][]bool, boundsMax.Y+1)
	for y := range grid {
		grid[y] = make([]bool, boundsMax.X+1)
	}

	for _, vertex := range input {
		grid[vertex.Y][vertex.X] = true
	}

	for i := 0; i < len(input); i++ {
		vertex := input[i]
		next := input[(i+1)%len(input)]

		for y := min(vertex.Y, next.Y); y <= max(vertex.Y, next.Y); y++ {
			for x := min(vertex.X, next.X); x <= max(vertex.X, next.X); x++ {
				grid[y][x] = true
			}
		}
	}

	for y, line := range grid {
		fill := false
		wall := false
		for x, cell := range line {
			if wall && cell {
				continue
			}

			wall = cell

			if cell { // we're hit a wall, flip the "colour"
				fill = !fill
				continue
			}

			grid[y][x] = fill
		}
	}

	maxArea := 0
	for i := 0; i < len(input); i++ {

	OuterLoop:
		for j := i + 1; j < len(input); j++ {
			// test if edges go outside the shape
			for n := min(input[i].X, input[j].X); n <= max(input[i].X, input[j].X); n++ {
				if !grid[input[i].Y][n] {
					continue OuterLoop
				}
			}
			for n := min(input[i].Y, input[j].Y); n <= max(input[i].Y, input[j].Y); n++ {
				if !grid[n][input[i].X] {
					continue OuterLoop
				}
			}
			for n := min(input[i].X, input[j].X); n <= max(input[i].X, input[j].X); n++ {
				if !grid[input[j].Y][n] {
					continue OuterLoop
				}
			}
			for n := min(input[i].Y, input[j].Y); n <= max(input[i].Y, input[j].Y); n++ {
				if !grid[n][input[j].X] {
					continue OuterLoop
				}
			}

			width := abs(input[i].X-input[j].X) + 1
			height := abs(input[i].Y-input[j].Y) + 1
			area := width * height

			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}
