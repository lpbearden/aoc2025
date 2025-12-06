#!/bin/bash
day=$(printf "%02d" $1)
mkdir -p day$day
cd day$day

cat > main.go <<'EOF'
package main

import (
	"fmt"
	"aoc2025/util"
)

func part1(input string) int {
	lines := util.Lines(input)
	_ = lines
	return 0
}

func part2(input string) int {
	lines := util.Lines(input)
	_ = lines
	return 0
}

func main() {
	input := util.ReadInput()
	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}
EOF

cat > main_test.go <<'EOF'
package main

import "testing"

const sample = ``

func TestPart1(t *testing.T) {
	got := part1(sample)
	want := 0
	if got != want {
		t.Errorf("part1() = %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := part2(sample)
	want := 0
	if got != want {
		t.Errorf("part2() = %v, want %v", got, want)
	}
}
EOF

touch input.txt
cd ..
echo "Created day$day"