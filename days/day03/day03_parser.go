package day03

import (
	"bufio"
	"io"
)

type Range struct {
	Min int
	Max int
}

type Input [][]byte

func ParseInput(input io.Reader) Input {
	scanner := bufio.NewScanner(input)
	var banks [][]byte

	for scanner.Scan() {
		line := scanner.Text()

		bank := make([]byte, len(line))
		for i := 0; i < len(line); i++ {
			bank[i] = line[i] - '0'
		}

		banks = append(banks, bank)
	}

	return banks
}
