package day03

import (
	"math"
)

func findLargestNumber(bank []byte, length int) int {
	num := 0
	lastIdx := -1

	for pow := length - 1; pow >= 0; pow-- {
		largestIdx := lastIdx + 1
		for i := largestIdx + 1; i < len(bank)-pow; i++ {
			if bank[i] > bank[largestIdx] {
				largestIdx = i
			}
		}

		lastIdx = largestIdx
		num += int(math.Pow10(pow)) * int(bank[largestIdx])
	}

	return num
}

func SolvePart1(input Input) int {
	sum := 0

	for _, bank := range input {
		largest := findLargestNumber(bank, 2)
		sum += largest
	}
	return sum
}

func SolvePart2(input Input) int {
	sum := 0

	for _, bank := range input {
		largest := findLargestNumber(bank, 12)
		sum += largest
	}
	return sum
}
