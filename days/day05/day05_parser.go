package day05

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

type Range struct {
	Start, End uint
}

func (r Range) String() string {
	return fmt.Sprintf("%d-%d", r.Start, r.End)
}

type Input struct {
	Ranges   []Range
	Products []uint
}

func ParseInput(input io.Reader) Input {
	scanner := bufio.NewScanner(input)

	var ranges []Range
	var products []uint

	parse := "ranges"
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			parse = "products"
			continue
		}

		switch parse {
		case "ranges":
			var start, end uint
			fmt.Sscanf(line, "%d-%d", &start, &end)
			ranges = append(ranges, Range{start, end})
		case "products":
			val, _ := strconv.ParseUint(line, 10, 64)
			products = append(products, uint(val))
		}
	}

	return Input{ranges, products}
}
