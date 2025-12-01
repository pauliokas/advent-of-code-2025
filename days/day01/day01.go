package day01

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func SolvePart1(input Input) int {
	zeroes := 0
	pointer := 50

	for _, rotateBy := range input {
		pointer = (100 + (pointer+rotateBy)%100) % 100

		if pointer == 0 {
			zeroes += 1
		}
	}

	return zeroes
}

func SolvePart2(input Input) int {
	zeroes := 0
	pointer := 50

	for _, rotateBy := range input {
		zeroes += abs(rotateBy) / 100
		rotateBy = rotateBy % 100

		newPointer := pointer + rotateBy

		if pointer != 0 && (newPointer <= 0 || newPointer >= 100) {
			zeroes += 1
		}

		pointer = (100 + newPointer%100) % 100
	}

	return zeroes
}
