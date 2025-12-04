package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Read file
	file, err := os.ReadFile("./input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
		os.Exit(1)
	}

	// Part 1
	lines := strings.Split(string(file), "\n")
	answer := Part1(lines)
	fmt.Printf("part 1: %v\n", answer)

	// Part 2
	var matrix [][]rune
	for _, line := range lines {
		matrix = append(matrix, []rune(line))
	}
	answer = Part2(matrix)
	fmt.Printf("part 2: %v\n", answer)
}

type Direction struct{ di, dj int }

var DIRECTIONS = []Direction{
	{1, 0},
	{1, 1},
	{0, 1},
	{-1, 1},
	{-1, 0},
	{-1, -1},
	{0, -1},
	{1, -1},
}

func Part1(lines []string) int {
	answer := 0

	for j, line := range lines {
		for i, c := range line {
			if c != '@' {
				continue
			}

			count := 0
			for _, direction := range DIRECTIONS {
				ni, nj := i+direction.di, j+direction.dj
				if !(0 <= nj && nj < len(lines) && 0 <= ni && ni < len(lines[nj])) {
					continue
				}

				if lines[nj][ni] == '@' {
					count++
				}
				if count >= 4 {
					break
				}
			}

			if count < 4 {
				answer++
			}
		}
	}

	return answer
}

func Part2(lines [][]rune) int {
	answer := 0

	for {
		changed := 0

		for j, line := range lines {
			for i, c := range line {
				if c != '@' {
					continue
				}

				count := 0
				for _, direction := range DIRECTIONS {
					ni, nj := i+direction.di, j+direction.dj
					if !(0 <= nj && nj < len(lines) && 0 <= ni && ni < len(lines[nj])) {
						continue
					}

					if lines[nj][ni] == '@' || lines[nj][ni] == 'x' {
						count++
					}
					if count >= 4 {
						break
					}
				}

				if count < 4 {
					changed++
					lines[j][i] = 'x'
				}
			}
		}

		for j, line := range lines {
			for i, c := range line {
				if c == 'x' {
					lines[j][i] = '.'
				}
			}
		}

		answer += changed
		if changed == 0 {
			break
		}
	}

	return answer
}
