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

func countPushes(joltages []int, buttons [][]int) int {
	zeroes := true
	negative := false
	lights := 0
	for i, joltage := range joltages {
		if joltage != 0 {
			zeroes = false
		}
		if joltage < 0 {
			negative = true
		}
		lights |= (joltage % 2) << i
	}

	if zeroes {
		return 0
	}

	if negative {
		return -1
	}

	if lights == 0 {
		for i, joltage := range joltages {
			joltages[i] = joltage / 2
		}
		result := countPushes(joltages, buttons)
		if result < 0 {
			return -1
		}
		return 2 * result
	}

	minButtonPresses := -1
	for permutation := range generatePermutations(buttons) {
		result := 0
		for _, buttonGroup := range permutation {
			for _, button := range buttonGroup {
				result ^= 1 << button
			}
		}

		if result == lights {
			newJoltages := make([]int, len(joltages))
			copy(newJoltages, joltages)
			for _, buttonGroup := range permutation {
				for _, button := range buttonGroup {
					newJoltages[button] -= 1
				}
			}

			presses := countPushes(newJoltages, buttons)
			if presses < 0 {
				continue
			}

			total := len(permutation) + presses
			if minButtonPresses < 0 || total < minButtonPresses {
				minButtonPresses = total
			}
		}
	}

	if minButtonPresses < 0 {
		return -1
	}

	return minButtonPresses
}

func SolvePart1(input Input) int {
	buttonPresses := 0

	for _, machine := range input {
		minPresses := 0

		ligths := 0
		for i, light := range machine.Lights {
			ligths |= light << i
		}

		for permutation := range generatePermutations(machine.Buttons) {
			result := 0
			for _, buttonGroup := range permutation {
				for _, button := range buttonGroup {
					result ^= 1 << button
				}
			}

			if result == ligths {
				if minPresses == 0 || len(permutation) < minPresses {
					minPresses = len(permutation)
				}
			}
		}

		buttonPresses += minPresses
	}

	return buttonPresses
}

func SolvePart2(input Input) int {
	buttonPresses := 0

	for _, machine := range input {
		buttonPresses += countPushes(machine.Joltage, machine.Buttons)
	}

	return buttonPresses
}
