package day02

import (
	"iter"
	"math"
)

func iterator(ranges []Range) iter.Seq[int] {
	return func(yield func(int) bool) {
		for _, r := range ranges {
			for i := r.Min; i <= r.Max; i++ {
				if !yield(i) {
					return
				}
			}
		}
	}
}

func getPattern(num int) (int, int) {
	digits := int(math.Floor(math.Log10(float64(num)))) + 1

	for length := digits / 2; length >= 1; length-- {
		if (digits % length) != 0 {
			continue
		}

		pattern := num % int(math.Pow10(length))

		allMatch := false
		for segment := 1; segment < digits/length; segment++ {
			allMatch = true

			part := num % int(math.Pow10((segment+1)*length))
			part /= int(math.Pow10(segment * length))

			if part != pattern {
				allMatch = false
				break
			}
		}

		if allMatch {
			return pattern, digits / length
		}
	}

	return -1, -1
}

func SolvePart1(input Input) int {
	sum := 0

	for i := range iterator(input) {
		_, count := getPattern(i)
		if count == 2 {
			sum += i
		}
	}

	return sum
}

func SolvePart2(input Input) int {
	sum := 0

	for i := range iterator(input) {
		_, count := getPattern(i)
		if count > 0 {
			sum += i
		}
	}

	return sum
}
