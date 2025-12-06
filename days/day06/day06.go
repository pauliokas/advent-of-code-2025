package day06

import (
	"strconv"
	"strings"
)

func mul(a, b int) int {
	return a * b
}

func add(a, b int) int {
	return a + b
}

func SolvePart1(input Input) int {
	sum := 0

	operations := input[len(input)-1]
	numbers := input[:len(input)-1]

	lastIdx := len(operations)
	for i := len(operations) - 1; i >= 0; i-- {
		var op func(int, int) int
		var cumulator int

		switch operations[i] {
		case ' ':
			continue
		case '+':
			op = add
			cumulator = 0
		case '*':
			op = mul
			cumulator = 1
		}

		for _, nums := range numbers {
			num, _ := strconv.Atoi(strings.Trim(nums[i:lastIdx], " "))
			cumulator = op(cumulator, num)
		}
		sum += cumulator

		lastIdx = i - 1
	}

	return sum
}

func SolvePart2(input Input) int {
	sum := 0

	operations := input[len(input)-1]
	numbers := input[:len(input)-1]

	lastIdx := len(operations)
	for i := len(operations) - 1; i >= 0; i-- {
		var op func(int, int) int
		var cumulator int

		switch operations[i] {
		case ' ':
			continue
		case '+':
			op = add
			cumulator = 0
		case '*':
			op = mul
			cumulator = 1
		}

		for j := lastIdx - 1; j >= i; j-- {
			number := 0
			for _, nums := range numbers {
				if nums[j] == ' ' {
					continue
				}
				digit := nums[j] - '0'
				number = number*10 + int(digit)
			}

			cumulator = op(cumulator, number)
		}

		sum += cumulator

		lastIdx = i - 1
	}

	return sum
}
