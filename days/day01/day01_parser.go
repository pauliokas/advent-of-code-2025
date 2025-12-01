package day01

import (
	"bufio"
	"io"
	"strconv"
)

type Input []int

func ParseInput(input io.Reader) Input {
	scanner := bufio.NewScanner(input)
	var rotations []int

	for scanner.Scan() {
		line := scanner.Text()

		var multiplier int
		switch line[0] {
		case 'L':
			multiplier = -1
		case 'R':
			multiplier = +1
		}

		steps, _ := strconv.Atoi(line[1:])
		rotations = append(rotations, steps*multiplier)
	}

	return rotations
}
