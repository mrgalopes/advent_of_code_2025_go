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
	answer := 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		answer += Joltage1(scanner.Text())
	}

	return answer
}

func Joltage1(battery string) int {
	firstDigit, idx := int(battery[0]-'0'), 0
	for i := 1; i < len(battery)-1; i++ {
		d := int(battery[i] - '0')
		if d > firstDigit {
			firstDigit = d
			idx = i
		}
	}

	secondDigit := int(battery[idx+1] - '0')
	for i := idx + 2; i < len(battery); i++ {
		d := int(battery[i] - '0')
		if d > secondDigit {
			secondDigit = d
		}
	}

	return 10*firstDigit + secondDigit
}

func Part2(input io.Reader) int {
	answer := 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		answer += Joltage2(scanner.Text())
	}

	return answer
}

func Joltage2(battery string) int {
	idxs := make([]int, 12)
	for i := range 12 {
		idxs[i] = i + len(battery) - 12
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
