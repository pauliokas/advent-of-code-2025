package day11

import "fmt"

func findPaths(grid Grid, start, end string, cache map[string]uint64) uint64 {
	var dfs func(string) uint64
	dfs = func(current string) uint64 {
		if current == end {
			return 1
		}

		cacheKey := fmt.Sprintf("%s|%s", current, end)
		if val, found := cache[cacheKey]; found {
			return val
		}

		pathCount := uint64(0)
		for neighbor := range grid[current] {
			pathCount += dfs(neighbor)
		}

		cache[cacheKey] = pathCount

		return pathCount
	}

	return dfs(start)
}

func SolvePart1(input Input) uint64 {
	cache := make(map[string]uint64, 100)

	return findPaths(Grid(input), "you", "out", cache)
}

func SolvePart2(input Input) uint64 {
	cache := make(map[string]uint64, 2000)

	svrToFft := findPaths(Grid(input), "svr", "fft", cache)
	fftToDac := findPaths(Grid(input), "fft", "dac", cache)
	dacToOut := findPaths(Grid(input), "dac", "out", cache)

	svrToDac := findPaths(Grid(input), "svr", "dac", cache)
	dacToFft := findPaths(Grid(input), "dac", "fft", cache)
	fftToOut := findPaths(Grid(input), "fft", "out", cache)

	return svrToFft*fftToDac*dacToOut + svrToDac*dacToFft*fftToOut
}
