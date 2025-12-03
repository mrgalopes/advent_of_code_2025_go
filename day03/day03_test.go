package main

import (
	"os"
	"strings"
	"testing"
)

func TestJoltage1(t *testing.T) {
	tests := []struct {
		battery string
		answer  int
	}{
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

func TestJoltage2(t *testing.T) {
	tests := []struct {
		battery string
		answer  int
	}{
		{
			battery: "987654321111111",
			answer:  987654321111,
		},
		{
			battery: "811111111111119",
			answer:  811111111119,
		},
		{
			battery: "234234234234278",
			answer:  434234234278,
		},
		{
			battery: "818181911112111",
			answer:  888911112111,
		},
	}

	for _, tC := range tests {
		t.Run(tC.battery, func(t *testing.T) {
			got := Joltage2(tC.battery)
			want := tC.answer
			if got != want {
				t.Fatalf("want %v, got %v", want, got)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	want := 3121910778619

	input := strings.NewReader(`987654321111111
811111111111119
234234234234278
818181911112111`)

	got := Part2(input)

	if want != got {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestAnswers(t *testing.T) {
	part1 := 17535
	part2 := 173577199527257

	file, _ := os.Open("./input")

	answer1 := Part1(file)
	if answer1 != part1 {
		t.Fatalf("want part1 %v, got %v", part1, answer1)
	}

	file.Seek(0, 0)
	answer2 := Part2(file)
	if answer2 != part2 {
		t.Fatalf("want part2 %v, got %v", part2, answer2)
	}
}
