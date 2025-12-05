package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening input: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	part1, ranges, err := Part1(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error part 1: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("part 1: %v\n", part1)

	minRanges := MinRanges(ranges)
	part2 := Part2(minRanges)
	fmt.Printf("part 2: %v\n", part2)
}

type Range struct{ start, end int }

func Part1(input io.Reader) (int, []Range, error) {
	scanner := bufio.NewScanner(input)

	// read ranges `3-5\n10-16\n` into []Range struct
	ranges := []Range{}
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}

		nums := strings.Split(text, "-")
		start, err := strconv.Atoi(nums[0])
		if err != nil {
			return 0, nil, fmt.Errorf("error converting start: %w", err)
		}

		end, err := strconv.Atoi(nums[1])
		if err != nil {
			return 0, nil, fmt.Errorf("error converting end: %w", err)
		}

		ranges = append(ranges, Range{start, end})
	}

	answer := 0
	// check numbers that are in the range
	// could call MinRanges before to be more efficient
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return 0, ranges, fmt.Errorf("error converting num: %w", err)
		}

		for _, r := range ranges {
			if r.start <= num && num <= r.end {
				answer++
				break
			}
		}
	}

	return answer, ranges, nil
}

func MinRanges(ranges []Range) []Range {
	// reduces ranges to be the minimal possible

	slices.SortFunc(ranges, func(a, b Range) int {
		return a.start - b.start
	})

	result := []Range{{ranges[0].start, ranges[0].end}}
	j := 0 // j -> index for `result`

	// i -> index for `ranges`
	for i := 1; i < len(ranges); i++ {
		r := ranges[i]
		if result[j].start <= r.start && r.start <= result[j].end {
			result[j].start = min(result[j].start, r.start)
			result[j].end = max(result[j].end, r.end)
		} else {
			result = append(result, r)
			j++
		}
	}

	return result
}

func Part2(ranges []Range) int {
	result := 0
	for _, r := range ranges {
		result += r.end - r.start + 1
	}

	return result
}
