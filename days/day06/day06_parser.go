package day06

import (
	"bufio"
	"io"
)

type Input []string

func ParseInput(input io.Reader) Input {
	scanner := bufio.NewScanner(input)

	var lines []string = make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
