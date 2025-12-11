package day10

import (
	"fmt"
	"iter"
	"slices"
	"strings"
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

func generateVectorCombinations(length int) iter.Seq[[]int] {
	return func(yield func([]int) bool) {
		vec := make([]int, length)

		var generate func(pos, maxVal int, hasMax bool) bool
		generate = func(pos, maxVal int, hasMax bool) bool {
			if pos >= length {
				if hasMax {
					result := make([]int, len(vec))
					copy(result, vec)
					return yield(result)
				}
				return true
			}

			for val := 0; val <= maxVal; val++ {
				vec[pos] = val
				if !generate(pos+1, maxVal, hasMax || val == maxVal) {
					return false
				}
			}
			return true
		}

		for maxVal := 0; ; maxVal++ {
			if !generate(0, maxVal, false) {
				return
			}
		}
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

// (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}

//  N = A + B + C + D + E
//
//  7 = A     + C + D
//  5 =             D + E
// 12 = A + B     + D + E
//  7 = A + B         + E
//  2 = A     + C     + E

type Matrix[T comparable] [][]T

func (m Matrix[T]) String() string {
	var sb strings.Builder
	for _, row := range m {
		for _, val := range row {
			sb.WriteString(fmt.Sprintf("%5v", val))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (m Matrix[T]) SortRows() {
	slices.SortStableFunc(m, func(a, b []T) int {
		var zero T
		for i := range a {
			if a[i] == zero && b[i] == zero {
				continue
			}
			if a[i] != zero {
				return -1
			}
			if b[i] != zero {
				return 1
			}
		}
		return 0
	})
}

func SolvePart2(input Input) int {
	buttonPresses := 0

	for _, machine := range input {
		matrix := make(Matrix[int], len(machine.Joltage))

		for i := range matrix {
			matrix[i] = make([]int, len(machine.Buttons)+1)
			matrix[i][len(machine.Buttons)] = -int(machine.Joltage[i])
		}

		for idx, button := range machine.Buttons {
			for _, joltageIdx := range button {
				matrix[joltageIdx][idx] = 1
			}
		}

		fmt.Printf("Matrix:\n%v\n", matrix)

	RowLoop:
		for rowIdx := 0; rowIdx < len(matrix)-1; rowIdx++ {
			matrix.SortRows()

			pivotIdx := slices.IndexFunc(matrix[rowIdx], func(val int) bool { return val != 0 })
			if pivotIdx < 0 {
				break RowLoop
			}
			pivotA := matrix[rowIdx][pivotIdx]

			for j := rowIdx + 1; j < len(matrix); j++ {
				if matrix[j][pivotIdx] == 0 {
					continue RowLoop
				}

				pivotB := matrix[j][pivotIdx]
				// avoid division, so multiply the row by the pivot of the "top" row
				for k := pivotIdx; k < len(matrix[rowIdx]); k++ {
					a := pivotB * matrix[rowIdx][k]
					b := pivotA * matrix[j][k]
					matrix[j][k] = a - b
				}
			}
		}

		// make leading coefficients positive
		for rowIdx := 0; rowIdx < len(matrix); rowIdx++ {
			pivotIdx := slices.IndexFunc(matrix[rowIdx], func(val int) bool { return val != 0 })
			if pivotIdx < 0 {
				break
			}

			if matrix[rowIdx][pivotIdx] < 0 {
				for k := pivotIdx; k < len(matrix[rowIdx]); k++ {
					matrix[rowIdx][k] = -matrix[rowIdx][k]
				}
			}
		}
		fmt.Printf("Echelon Matrix:\n%v\n", matrix)

		answerVector := make([]int, len(machine.Buttons)+1)
		answerVector[len(answerVector)-1] = 1
		for i := range answerVector[:len(answerVector)-1] {
			answerVector[i] = -1
		}

		// Back substitution
		for rowIdx := len(matrix) - 1; rowIdx >= 0; rowIdx-- {
			pivotIdx := slices.IndexFunc(matrix[rowIdx], func(val int) bool { return val != 0 })
			if pivotIdx < 0 || pivotIdx >= len(machine.Buttons) {
				continue
			}

			freeVariable := false
			sum := 0
			for colIdx := pivotIdx + 1; colIdx < len(answerVector); colIdx++ {
				if matrix[rowIdx][colIdx] != 0 && answerVector[colIdx] < 0 {
					freeVariable = true
					continue
				}
				sum += matrix[rowIdx][colIdx] * answerVector[colIdx]
			}

			if !freeVariable {
				answerVector[pivotIdx] = -sum / matrix[rowIdx][pivotIdx]
			}
		}

		freeVariableIndices := make([]int, 0, len(answerVector))
		for i, val := range answerVector[:len(answerVector)-1] {
			if val < 0 {
				freeVariableIndices = append(freeVariableIndices, i)
			}
		}
		fmt.Printf("Answer vector: %v\n", answerVector)

		for combination := range generateVectorCombinations(len(freeVariableIndices)) {
			for i, freeVarIdx := range freeVariableIndices {
				answerVector[freeVarIdx] = combination[i]
			}

			valid := true
			for rowIdx := 0; rowIdx < len(matrix); rowIdx++ {
				sum := 0
				for colIdx := 0; colIdx < len(answerVector); colIdx++ {
					sum += matrix[rowIdx][colIdx] * answerVector[colIdx]
				}
				if sum != 0 {
					valid = false
					break
				}
			}

			if valid {
				break
			}
		}
		fmt.Printf("Combination chosen: %v\n", answerVector)

		// expected [5, 0, 5, 1]

		// c = 5
		// b = 1 + d
		// a = 5 - b

		for _, presses := range answerVector[:len(answerVector)-1] {
			buttonPresses += int(presses)
		}
	}

	return buttonPresses
}
