package main

import (
	"strings"
	"testing"
)

func TestJoltage1(t *testing.T) {
	tests := []struct{
		battery string
		answer int
	} {
		{
			battery: "987654321111111",
			answer:  98,
		},
		{
			battery: "811111111111119",
			answer:  89,
		},
	}

	for _, tC := range tests {
		t.Run(tC.battery, func(t *testing.T) {
			got := Joltage1(tC.battery)
			want := tC.answer
			if got != want {
				t.Fatalf("want %v, got %v", want, got)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	want := 357

	input := strings.NewReader(`987654321111111
811111111111119
234234234234278
818181911112111`)

	got := Part1(input)

	if want != got {
		t.Fatalf("want %v, got %v", want, got)
	}
}