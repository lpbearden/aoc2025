package main

import "testing"

const sample = `paste sample input here`

func TestPart1(t *testing.T) {
	got := part1(sample)
	want := 0 // expected answer

	if got != want {
		t.Errorf("part1() = %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := part2(sample)
	want := 0 // expected answer

	if got != want {
		t.Errorf("part2() = %v, want %v", got, want)
	}
}
