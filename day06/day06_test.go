package main

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	input := `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +`
	expectedNums := [][]int{
		{123, 328, 51, 64},
		{45, 64, 387, 23},
		{6, 98, 215, 314},
	}
	expectedOps := []string{
		"*", "+", "*", "+",
	}

	nums, ops, err := ParseInput(input)

	if err != nil {
		t.Fatalf("want err nil, got %v", err)
	}

	if !reflect.DeepEqual(nums, expectedNums) {
		t.Fatalf("want %+v, got %+v", expectedNums, nums)
	}

	if !reflect.DeepEqual(ops, expectedOps) {
		t.Fatalf("want %+v, got %+v", expectedOps, ops)
	}
}

func TestPart1(t *testing.T) {
	nums := [][]int{
		{123, 328, 51, 64},
		{45, 64, 387, 23},
		{6, 98, 215, 314},
	}
	ops := []string{
		"*", "+", "*", "+",
	}
	want := 4277556

	got := Part1(nums, ops)

	if want != got {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestPart2(t *testing.T) {
	input := `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `

	want := 3263827

	got := Part2(input)

	if want != got {
		t.Fatalf("want %v, got %v", want, got)
	}
}
