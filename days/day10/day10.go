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

type Matrix [][]float64

func NewMatrix(rows, cols int) Matrix {
	m := make(Matrix, rows)
	for i := range m {
		m[i] = make([]float64, cols)
	}
	return m
}

func (m Matrix) String() string {
	var sb strings.Builder
	for _, row := range m {
		for _, val := range row {
			sb.WriteString(fmt.Sprintf("%5v", val))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (m Matrix) SortRows() {
	slices.SortStableFunc(m, func(a, b []float64) int {
		for i := range a {
			if a[i] == 0 && b[i] == 0 {
				continue
			}
			if a[i] != 0 {
				return -1
			}
			if b[i] != 0 {
				return 1
			}
		}
		return 0
	})
}

func IterateGuesses(equation []float64, solutions []float64, maxSum *int) iter.Seq[[]float64] {
	unsolvedInEquation := make([]int, 0, len(equation))
	for i, val := range equation[:len(equation)-1] {
		if val != 0 && solutions[i] < 0 {
			unsolvedInEquation = append(unsolvedInEquation, i)
		}
	}

	if len(unsolvedInEquation) == 0 {
		return func(yield func([]float64) bool) {
			sum := equation[len(equation)-1] // constant term
			for i, val := range equation[:len(equation)-1] {
				if val != 0 {
					sum += val * solutions[i]
				}
			}
			if sum == 0 {
				result := make([]float64, len(solutions))
				copy(result, solutions)
				yield(result)
			}
		}
	}

	constantTerm := equation[len(equation)-1]

	currentSum := 0
	for i := 0; i < len(solutions)-1; i++ {
		if solutions[i] >= 0 {
			currentSum += int(solutions[i])
		}
	}

	return func(yield func([]float64) bool) {
		state := make([]float64, len(solutions))
		copy(state, solutions)

		fixedSum := 0.0
		for i := 0; i < len(equation)-1; i++ {
			if equation[i] != 0 && state[i] >= 0 {
				fixedSum += equation[i] * state[i]
			}
		}

		pivotIdx := unsolvedInEquation[0]
		pivotCoeff := equation[pivotIdx]

		freeVars := unsolvedInEquation[1:]

		var gen func(depth int, runningSum int) bool
		gen = func(depth int, runningSum int) bool {
			if maxSum != nil && *maxSum >= 0 && runningSum > *maxSum {
				return true
			}

			if depth == len(freeVars) {
				sum := fixedSum + constantTerm
				for _, idx := range freeVars {
					sum += equation[idx] * state[idx]
				}

				pivot := -sum / pivotCoeff
				rounded := float64(int(pivot + 0.5))
				if pivot < 0 || pivot != rounded {
					return true
				}

				if maxSum != nil && *maxSum >= 0 && runningSum+int(pivot) > *maxSum {
					return true
				}

				state[pivotIdx] = pivot
				result := make([]float64, len(state))
				copy(result, state)
				return yield(result)
			}

			freeIdx := freeVars[depth]
			freeCoeff := equation[freeIdx]

			var maxVal float64 = -1
			if maxSum != nil && *maxSum >= 0 {
				maxVal = float64(*maxSum - runningSum)
			}

			for val := 0.0; maxVal < 0 || val <= maxVal; val++ {
				state[freeIdx] = val

				if !gen(depth+1, runningSum+int(val)) {
					return false
				}

				sum := fixedSum + constantTerm
				for _, idx := range freeVars {
					sum += equation[idx] * state[idx]
				}
				pivot := -sum / pivotCoeff

				if pivot < 0 && (freeCoeff/pivotCoeff) > 0 {
					break
				}
			}

			state[freeIdx] = -1
			return true
		}

		gen(0, currentSum)
	}
}

func SolvePart2(input Input) int {
	totalButtonPresses := 0

	for _, machine := range input {
		matrix := NewMatrix(len(machine.Joltage), len(machine.Buttons)+1)

		for i := range matrix {
			matrix[i][len(machine.Buttons)] = -float64(machine.Joltage[i])
		}

		for idx, button := range machine.Buttons {
			for _, joltageIdx := range button {
				matrix[joltageIdx][idx] = 1
			}
		}

	RowLoop:
		for rowIdx := 0; rowIdx < len(matrix)-1; rowIdx++ {
			matrix.SortRows()

			pivotIdx := slices.IndexFunc(matrix[rowIdx], func(val float64) bool { return val != 0 })
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

		// Start with unknown solution vector (-1 means unsolved)
		initialSolution := make([]float64, len(machine.Buttons)+1)
		initialSolution[len(initialSolution)-1] = 1 // constant multiplier
		for i := range initialSolution[:len(initialSolution)-1] {
			initialSolution[i] = -1
		}

		// Find the last non-zero row
		lastNonZeroRow := len(matrix) - 1
		for lastNonZeroRow >= 0 {
			allZero := true
			for _, val := range matrix[lastNonZeroRow] {
				if val != 0 {
					allZero = false
					break
				}
			}
			if !allZero {
				break
			}
			lastNonZeroRow--
		}

		numCols := len(matrix[0])
		initialBound := 0
		for row := 0; row <= lastNonZeroRow; row++ {
			constantTerm := matrix[row][numCols-1]
			if constantTerm < 0 {
				initialBound += int(-constantTerm)
			} else {
				initialBound += int(constantTerm)
			}
		}
		minForMachine := initialBound
		foundSolution := false
		var solveFromRow func(rowIdx int, currentSolution []float64) bool
		solveFromRow = func(rowIdx int, currentSolution []float64) bool {
			if minForMachine >= 0 {
				partialSum := 0
				for i := 0; i < len(currentSolution)-1; i++ {
					if currentSolution[i] >= 0 {
						partialSum += int(currentSolution[i])
					}
				}
				if partialSum > minForMachine {
					return true
				}
			}

			if rowIdx < 0 {
				total := 0
				for i := 0; i < len(currentSolution)-1; i++ {
					total += int(currentSolution[i])
				}
				foundSolution = true
				if minForMachine < 0 || total < minForMachine {
					minForMachine = total
				}
				return true
			}

			for solution := range IterateGuesses(matrix[rowIdx], currentSolution, &minForMachine) {
				solveFromRow(rowIdx-1, solution)
			}
			return true
		}

		solveFromRow(lastNonZeroRow, initialSolution)
		if foundSolution && minForMachine > 0 {
			totalButtonPresses += minForMachine
		}
	}

	return totalButtonPresses
}
