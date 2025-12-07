package main

import "testing"

const sample = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

const sample2 = `L150
L50`

const sample3 = `L150
R50`

const sample4 = `R150
L50`

const sample5 = `R150
R50`

func TestPart1(t *testing.T) {
	got := part1(sample)
	want := 3 // expected answer

	if got != want {
		t.Errorf("part1() = %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"sample", sample, 6},
		{"sample", sample2, 2},
		{"sample", sample3, 2},
		{"sample", sample4, 2},
		{"sample", sample5, 2},
		// Add more test cases here
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := part2(tt.input)
			if got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
