package day12

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
	"strings"
)

var dimensionsPattern *regexp.Regexp = regexp.MustCompile(`(\d+x\d+): `)

type Shape [][]bool

func (s Shape) String() string {
	var sb strings.Builder
	for _, row := range s {
		for _, cell := range row {
			if cell {
				sb.WriteString("#")
			} else {
				sb.WriteString(".")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

type Region struct {
	Width      int
	Height     int
	Quantities []int
}

type Input struct {
	Shapes  []Shape
	Regions []Region
}

func ParseInput(input io.Reader) Input {
	scanner := bufio.NewScanner(input)

	shapes := make([]Shape, 0)
	regions := make([]Region, 0)
	for scanner.Scan() {
		// scanner.Text()

		if len(shapes) < 6 {
			shape := make(Shape, 0, 3)
			for range 3 {
				scanner.Scan()
				line := scanner.Text()
				row := make([]bool, 0, 3)
				for _, char := range line {
					if char == '#' {
						row = append(row, true)
					} else {
						row = append(row, false)
					}
				}
				shape = append(shape, row)
			}
			shapes = append(shapes, shape)
			scanner.Scan()
		} else {
			line := scanner.Text()
			dimensionsRes := dimensionsPattern.FindStringSubmatch(line)
			separatorIdx := strings.Index(dimensionsRes[1], "x")
			width, _ := strconv.Atoi(dimensionsRes[1][:separatorIdx])
			height, _ := strconv.Atoi(dimensionsRes[1][separatorIdx+1:])

			quantities := make([]int, 0)
			for _, num := range strings.Split(line[len(dimensionsRes[0]):], " ") {
				val, _ := strconv.Atoi(num)
				quantities = append(quantities, val)
			}

			regions = append(regions, Region{
				Width:      width,
				Height:     height,
				Quantities: quantities,
			})
		}
	}

	return Input{Shapes: shapes, Regions: regions}
}
