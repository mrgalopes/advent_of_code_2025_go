package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	input := strings.NewReader(`3-5
10-14
16-20
12-18

1
5
8
11
17
32`)
	want := 3
	got, _, err := Part1(input)

	if err != nil {
		t.Fatalf("want err nil, got %v", err)
	}

	if got != want {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestMinRanges(t *testing.T) {
	ranges := []Range{
		{3, 5}, {10, 14}, {16, 20}, {12, 18},
	}

	want := []Range{
		{3, 5}, {10, 20},
	}

	got := MinRanges(ranges)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %+v, got %+v", want, got)
	}
}

func TestPart2(t *testing.T) {
	minRanges := []Range{
		{3, 5}, {10, 20},
	}
	want := 14

	got := Part2(minRanges)

	if want != got {
		t.Fatalf("want %v, got %v", want, got)
	}
}
