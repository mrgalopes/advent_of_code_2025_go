package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// Read file
	file, err := os.Open("./input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not read file: %v", err)
		os.Exit(1)
	}

	// Part 1
	answer := Part1(file)
	fmt.Printf("part 1: %v\n", answer)

	// Part 2
	file.Seek(0, 0)
	answer = Part2(file)
	fmt.Printf("part 2: %v\n", answer)
}

func Part1(input io.Reader) int {
	return Solve(input, 2)
}

func Part2(input io.Reader) int {
	return Solve(input, 12)
}

func Solve(input io.Reader, quantity int) int {
	answer := 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		answer += Joltage(scanner.Text(), quantity)
	}

	return answer
}

func Joltage(battery string, quantity int) int {
	idxs := make([]int, quantity)
	for i := range quantity {
		idxs[i] = i + len(battery) - quantity
	}

	for j := 0; j < len(idxs); j++ {
		m, idx, end := 0, idxs[j], -1
		if j > 0 {
			end = idxs[j-1]
		}
		for i := idxs[j]; i > end; i-- {
			d := int(battery[i] - '0')
			if d >= m {
				m = d
				idx = i
			}
		}
		idxs[j] = idx
	}

	answer, mult := 0, 1
	for i := len(idxs) - 1; i >= 0; i-- {
		answer += int(battery[idxs[i]]-'0') * mult
		mult *= 10
	}
	return answer
}
