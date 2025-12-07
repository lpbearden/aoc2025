package main

import (
	"fmt"
	"strconv"

	"aoc2025/util"
)

func part1(input string) int {
	lines := util.Lines(input)
	_ = lines

	lockPosition := 50
	zc := 0

	for _, value := range lines {
		dir := value[0]
		steps, err := strconv.Atoi(value[1:])
		if err != nil {
			panic(err)
		}

		if dir == 'R' {
			lockPosition += steps
			for lockPosition > 99 {
				lockPosition = lockPosition - 100
			}
		}

		if dir == 'L' {
			lockPosition -= steps
			for lockPosition < 0 {
				lockPosition = 100 + lockPosition
			}
		}

		if lockPosition == 0 {
			zc++
		}
	}

	return zc
}

func part2(input string) int {
	lines := util.Lines(input)

	lockPosition := 50
	zc := 0

	for _, value := range lines {
		dir := value[0]
		steps, err := strconv.Atoi(value[1:])
		if err != nil {
			panic(err)
		}
		old := lockPosition

		if dir == 'R' {
			// Going right from `old` by `steps`, we cross 0 each time we pass a multiple of 100
			zc += (old + steps) / 100

			lockPosition = (old + steps) % 100
		}

		if dir == 'L' {
			if old == 0 {
				// Starting from 0, we only hit it again after full rotations
				zc += steps / 100
			} else if steps >= old {
				// We hit 0 once when we reach it, plus once for each additional 100 steps
				zc += 1 + (steps-old)/100
			}

			newPos := (old - steps) % 100
			if newPos < 0 {
				newPos += 100
			}
			lockPosition = newPos
		}
	}

	return zc
}

func main() {
	input := util.ReadInput()

	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}
