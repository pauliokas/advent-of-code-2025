package day08

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Coords struct {
	X, Y, Z int
}

func (c Coords) String() string {
	return fmt.Sprintf("(%d,%d,%d)", c.X, c.Y, c.Z)
}

type Input []Coords

func ParseInput(input io.Reader) Input {
	scanner := bufio.NewScanner(input)

	var coords []Coords = make([]Coords, 0)
	for scanner.Scan() {
		line := scanner.Text()

		splitStr := strings.Split(line, ",")

		x, _ := strconv.Atoi(splitStr[0])
		y, _ := strconv.Atoi(splitStr[1])
		z, _ := strconv.Atoi(splitStr[2])

		coords = append(coords, Coords{X: x, Y: y, Z: z})
	}

	return coords
}
