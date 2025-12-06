package util

import (
	"os"
	"strconv"
	"strings"
)

// ReadInput reads input.txt from the same directory
func ReadInput() string {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return string(data)
}

// Lines splits input into lines and trims whitespace
func Lines(input string) []string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
	}
	return lines
}

// Ints converts string slice to ints
func Ints(lines []string) []int {
	nums := make([]int, len(lines))
	for i, line := range lines {
		n, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			panic(err)
		}
		nums[i] = n
	}
	return nums
}

// Abs returns absolute value
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
