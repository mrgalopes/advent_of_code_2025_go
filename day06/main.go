package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("./input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not read input: %v\n", err)
		os.Exit(1)
	}

	input := strings.Trim(string(file), "\n")
	nums, ops, err := ParseInput(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing input: %v\n", err)
		os.Exit(1)
	}

	part1 := Part1(nums, ops)
	fmt.Printf("part 1: %v\n", part1)

	part2 := Part2(input)
	fmt.Printf("part 2: %v\n", part2)
}

func ParseInput(input string) ([][]int, []string, error) {
	lines := strings.Split(input, "\n")

	lastLine := lines[len(lines)-1]
	ops := strings.Fields(lastLine)

	nums := [][]int{}
	for i := 0; i < len(lines)-1; i++ {
		splitLine := strings.Fields(lines[i])

		row := []int{}
		for _, c := range splitLine {
			num, err := strconv.Atoi(c)
			if err != nil {
				return nil, nil, fmt.Errorf("error converting num in line %v: %w", lines[i], err)
			}

			row = append(row, num)
		}
		nums = append(nums, row)
	}

	return nums, ops, nil
}

func Part1(nums [][]int, ops []string) int {
	answer := 0
	ROWS, COLS := len(nums), len(nums[0])

	for i := 0; i < COLS; i++ {
		op := ops[i]

		total := 0
		if op == "*" {
			total = 1
		}

		for j := 0; j < ROWS; j++ {
			switch op {
			case "+":
				total += nums[j][i]
			case "*":
				total *= nums[j][i]
			default:
				fmt.Fprintf(os.Stderr, "unexpected op: %v", op)
			}
		}

		answer += total
	}

	return answer
}

func Part2(input string) int {
	answer := 0
	lines := strings.Split(input, "\n")

	nums := lines[:len(lines)-1]
	ops := lines[len(lines)-1]

	currNums := []int{}
	for i := len(lines[0]) - 1; i >= 0; i-- {
		curr := 0
		mult := 1
		for j := len(nums) - 1; j >= 0; j-- {
			alg := int(nums[j][i] - '0')
			if !(0 <= alg && alg <= 9) {
				continue
			}
			curr += alg * mult
			mult *= 10
		}
		currNums = append(currNums, curr)

		switch ops[i] {
		case '+':
			for _, num := range currNums {
				answer += num
			}
			currNums = []int{}
			i--
		case '*':
			total := 1
			for _, num := range currNums {
				total *= num
			}
			answer += total
			currNums = []int{}
			i--
		}
	}

	return answer
}
