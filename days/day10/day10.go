package day10

import (
	"iter"
)

func generatePermutations[T any](items []T) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		var gen func(start int, subset []T) bool
		gen = func(start int, subset []T) bool {
			if start >= len(items) {
				return true
			}

			for i := start; i < len(items); i++ {
				newArr := make([]T, len(subset), len(subset)+1)
				copy(newArr, subset)
				newArr = append(newArr, items[i])

				if !yield(newArr) {
					return false
				}

				if !gen(i+1, newArr) {
					return false
				}
			}
			return true
		}
		gen(0, []T{})
	}
}

func SolvePart1(input Input) int {
	buttonPresses := 0

	for _, machine := range input {
		var mostEfficientPesses [][]byte
		result := uint16(0)
		for permutation := range generatePermutations(machine.Buttons) {
			for _, buttonGroup := range permutation {
				for _, button := range buttonGroup {
					result ^= 1 << button
				}
			}

			if result == machine.Lights {
				if len(mostEfficientPesses) == 0 || len(permutation) < len(mostEfficientPesses) {
					mostEfficientPesses = permutation
				}
			}

			result = 0
		}

		buttonPresses += len(mostEfficientPesses)
	}

	return buttonPresses
}

func SolvePart2(input Input) int {
	return 0
}
