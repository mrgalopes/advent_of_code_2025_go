package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	file := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`
	lines := strings.Split(file, "\n")

	want := 13
	got := Part1(lines)

	if want != got {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestPart2(t *testing.T) {
	file := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`
	lines := strings.Split(file, "\n")
	var matrix [][]rune
	for _, line := range lines {
		matrix = append(matrix, []rune(line))
	}

	want := 43
	got := Part2(matrix)

	if want != got {
		t.Fatalf("want %v, got %v", want, got)
	}
}
