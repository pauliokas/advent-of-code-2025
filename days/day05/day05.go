package day05

import (
	"cmp"
	"slices"
)

func isFresh(ranges []Range, product uint) bool {
	for _, r := range ranges {
		if product >= r.Start && product <= r.End {
			return true
		}
	}
	return false
}

func SolvePart1(input Input) int {
	freshCount := 0
	for _, product := range input.Products {
		if isFresh(input.Ranges, product) {
			freshCount++
		}
	}
	return freshCount
}

func max(a, b uint) uint {
	if a > b {
		return a
	}
	return b
}

func SolvePart2(input Input) int {
	ranges := input.Ranges

	slices.SortFunc(ranges, func(a, b Range) int {
		return cmp.Compare(a.Start, b.Start)
	})

	for i := 0; i < len(ranges)-1; i++ {
		current := ranges[i]
		next := ranges[i+1]

		if current.End+1 >= next.Start {
			merged := Range{
				Start: current.Start,
				End:   max(current.End, next.End),
			}

			ranges[i] = merged
			ranges = append(ranges[:i+1], ranges[i+2:]...)

			i -= 1
		}
	}

	var totalFreshProducts uint = 0
	for _, r := range ranges {
		totalFreshProducts += r.End - r.Start + 1
	}

	return int(totalFreshProducts)
}
