package day12

func SolvePart1(input Input) int {
	fittingRegions := 0
	for _, region := range input.Regions {
		area := region.Width * region.Height
		for _, qty := range region.Quantities {
			area -= qty * 9
		}
		if area >= 0 {
			fittingRegions += 1
		}
	}

	return fittingRegions
}
