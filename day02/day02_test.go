package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"
)

const EXAMPLE_INPUT = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`

func TestInvalidIDs(t *testing.T) {
	testCases := []struct {
		start, end int64
		answer     []int64
	}{
		{
			start:  11,
			end:    22,
			answer: []int64{11, 22},
		},
		{
			start:  95,
			end:    115,
			answer: []int64{99},
		},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("%d-%d", tC.start, tC.end), func(t *testing.T) {
			got := InvalidIDsBetween(tC.start, tC.end)

			if !reflect.DeepEqual(got, tC.answer) {
				t.Fatalf("want %+v, got %+v", tC.answer, got)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	const want int64 = 1227775554
	got, err := Part1(EXAMPLE_INPUT)

	if err != nil {
		t.Fatalf("want err nil, got %v", err)
	}

	if want != got {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestInvalidIDs2(t *testing.T) {
	testCases := []struct {
		start, end int64
		answer     []int64
	}{
		{
			start:  11,
			end:    22,
			answer: []int64{11, 22},
		},
		{
			start:  95,
			end:    115,
			answer: []int64{99, 111},
		},
		{
			start:  1188511880,
			end:    1188511890,
			answer: []int64{1188511885},
		},
		{
			start:  222220,
			end:    222224,
			answer: []int64{222222},
		},
		{
			start:  1698522,
			end:    1698528,
			answer: []int64{},
		},
		{
			start:  2121212118,
			end:    2121212124,
			answer: []int64{2121212121},
		},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("%d-%d", tC.start, tC.end), func(t *testing.T) {
			got := InvalidIDsBetween2(tC.start, tC.end)

			if !reflect.DeepEqual(got, tC.answer) {
				t.Fatalf("want %+v, got %+v", tC.answer, got)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	const want int64 = 4174379265
	got, err := Part2(EXAMPLE_INPUT)

	if err != nil {
		t.Fatalf("want err nil, got %v", err)
	}

	if want != got {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestAnswers(t *testing.T) {
	const part1 int64 = 18595663903
	const part2 int64 = 19058204438

	file, _ := os.ReadFile("input")
	input := strings.Trim(string(file), " \n")

	// Part 1
	answer, err := Part1(input)
	if err != nil {
		t.Fatalf("want err nil, got %v", err)
	}
	if part1 != answer {
		t.Fatalf("want part1 %v, got %v", part1, answer)
	}

	// Part 2
	answer, err = Part2(input)
	if err != nil {
		t.Fatalf("want err nil, got %v", err)
	}
	if part2 != answer {
		t.Fatalf("want part2 %v, got %v", part1, answer)
	}
}
