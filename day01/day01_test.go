package main

import (
	"os"
	"strings"
	"testing"
)

func TestNewDial(t *testing.T) {
	testCases := []struct {
		text string
		dial Dial
	}{
		{
			text: "L20",
			dial: Dial{LEFT, 20},
		},
		{
			text: "R50",
			dial: Dial{RIGHT, 50},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.text, func(t *testing.T) {
			want := tC.dial
			got, err := NewDial(tC.text)
			if err != nil {
				t.Fatalf("want err nil, got %v", err)
			}
			if want != got {
				t.Fatalf("want %v, got %v", want, got)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	file := strings.NewReader(`L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`)
	want := 3

	got, err := Part1(file)

	if err != nil {
		t.Fatalf("want err nil, got %v", err)
	}

	if want != got {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestPart2(t *testing.T) {
	file := strings.NewReader(`L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`)
	want := 6

	got, err := Part2(file)

	if err != nil {
		t.Fatalf("want err nil, got %v", err)
	}

	if want != got {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestPart2Scenarios(t *testing.T) {
	testCases := []struct {
		file   string
		answer int
	}{
		{
			file:   `L51`,
			answer: 1,
		},
		{
			file:   `L50`,
			answer: 1,
		},
		{
			file:   `L151`,
			answer: 2,
		},
		{
			file: `L151
R2
L2`,
			answer: 4,
		},
		{
			file:   `L150`,
			answer: 2,
		},
		{
			file: `R50
L100
R100`,
			answer: 3,
		},
		{
			file: `R50
L5`,
			answer: 1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.file, func(t *testing.T) {
			file := strings.NewReader(tC.file)

			want := tC.answer
			got, err := Part2(file)

			if err != nil {
				t.Fatalf("want err nil, got %v", err)
			}

			if want != got {
				t.Fatalf("want %v, got %v", want, got)
			}
		})
	}
}

func TestAnswers(t *testing.T) {
	part1, part2 := 982, 6106

	// Read file
	file, err := os.Open("./input")
	if err != nil {
		t.Fatalf("could not read file: %v", err)
	}

	// Part 1
	answer, err := Part1(file)
	if err != nil {
		t.Fatalf("want err nil, got %v", err)
	}
	if answer != part1 {
		t.Fatalf("want answer %v, got %v", part1, answer)
	}

	// Part 2
	file.Seek(0, 0)
	answer, err = Part2(file)
	if err != nil {
		t.Fatalf("want err nil, got %v", err)
	}
	if answer != part2 {
		t.Fatalf("want answer %v, got %v", part1, answer)
	}
}
