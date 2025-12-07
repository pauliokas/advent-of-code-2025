package day07

type Stack[T any] []*T

func (stack *Stack[T]) Push(value *T) {
	*stack = append(*stack, value)
}

func (stack *Stack[T]) Pop() *T {
	l := len(*stack)
	value := (*stack)[l-1]
	*stack = (*stack)[:l-1]
	return value
}

func SolvePart1(input Input) int {
	grid := input.Grid
	stack := make(Stack[Coords], 0, grid.Height)
	visited := make(map[Coords]bool)

	stack.Push(&input.Start)

	splitCount := 0

	for len(stack) > 0 {
		current := *stack.Pop()

		next := Coords{X: current.X, Y: current.Y + 1}
		for next.Y < grid.Height && grid.At(next) != '^' {
			next.Y += 1
		}

		if grid.At(next) == '^' {
			if _, ok := visited[next]; ok {
				continue
			}
			visited[next] = true

			stack.Push(&Coords{next.X + 1, next.Y})
			stack.Push(&Coords{next.X - 1, next.Y})
			splitCount += 1
		}
	}

	return splitCount
}

type StackValue struct {
	Path []Coords
	Beam Coords
}

func SolvePart2(input Input) int {
	grid := input.Grid
	stack := make(Stack[StackValue], 0, grid.Height)
	cache := make(map[Coords]int)

	stack.Push(&StackValue{Path: []Coords{input.Start}, Beam: input.Start})

	for len(stack) > 0 {
		stackHead := *stack.Pop()
		path := stackHead.Path
		current := stackHead.Beam

		next := Coords{X: current.X, Y: current.Y + 1}
		for next.Y < grid.Height && grid.At(next) != '^' {
			next.Y += 1
		}

		if cacheVal, cacheHit := cache[next]; next.Y == grid.Height || cacheHit {
			if !cacheHit {
				cacheVal = 1
			}
			for _, coord := range path {
				cache[coord] += cacheVal
			}
			continue
		}

		if grid.At(next) == '^' {
			newPath := append([]Coords(nil), path...)
			newPath = append(newPath, next)

			stack.Push(&StackValue{Path: newPath, Beam: Coords{next.X + 1, next.Y}})
			stack.Push(&StackValue{Path: newPath, Beam: Coords{next.X - 1, next.Y}})
		}
	}

	return cache[input.Start]
}
