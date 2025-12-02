package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read file
	file, err := os.ReadFile("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
		os.Exit(1)
	}
	input := strings.Trim(string(file), " \n")

	// Part 1
	answer, err := Part1(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error part 1: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Part 1: %v\n", answer)

	// Part 2
	answer, err = Part2(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error part 2: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Part 2: %v\n", answer)
}

func Part1(input string) (int64, error) {
	return CalculateAnswer(input, InvalidIDsBetween)
}

func Part2(input string) (int64, error) {
	return CalculateAnswer(input, InvalidIDsBetween2)
}

func CalculateAnswer(input string, intervalFunc func(int64, int64) []int64) (int64, error) {
	var answer int64 = 0

	intervals := strings.Split(input, ",")
	for _, interval := range intervals {
		i := strings.Split(interval, "-")
		if len(i) != 2 {
			return 0, fmt.Errorf("error reading interval, got %v", i)
		}
		startStr, endStr := i[0], i[1]

		start, err := strconv.ParseInt(startStr, 10, 0)
		if err != nil {
			return 0, fmt.Errorf("error converting start interval %v: %w", startStr, err)
		}

		end, err := strconv.ParseInt(endStr, 10, 0)
		if err != nil {
			return 0, fmt.Errorf("error converting end interval %v: %w", endStr, err)
		}

		invalidIDs := intervalFunc(start, end)
		for _, id := range invalidIDs {
			answer += id
		}
	}

	return answer, nil
}

func InvalidIDsBetween(start, end int64) []int64 {
	result := []int64{}

	for i := start; i <= end; i++ {
		s := strconv.FormatInt(i, 10)
		if len(s)%2 != 0 {
			continue
		}
		firstHalf, secondHalf := s[:len(s)/2], s[len(s)/2:]

		if firstHalf == secondHalf {
			result = append(result, i)
		}
	}

	return result
}

func InvalidIDsBetween2(start, end int64) []int64 {
	result := []int64{}

	for n := start; n <= end; n++ {
		s := strconv.FormatInt(n, 10)
		isInvalid := false

		for i := 1; i <= len(s)/2; i++ {
			if len(s)%i != 0 {
				continue
			}
			s1 := s[:i]
			b := true
			for j := i; j+i <= len(s); j += i {
				s2 := s[j : i+j]
				if s1 != s2 {
					b = false
					break
				}
			}
			if b {
				isInvalid = true
				break
			}
		}

		if isInvalid {
			result = append(result, n)
		}
	}

	return result
}
