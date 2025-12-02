package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	// Read file
	file, err := os.Open("day01/input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
		os.Exit(1)
	}

	// Part 1
	result, err := Part1(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error calculating part 1: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("part 1: %v\n", result)

	// Part 2
	file.Seek(0, 0)
	result, err = Part2(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error calculating part 2: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("part 2: %v\n", result)
}

func Part1(file io.Reader) (int, error) {
	answer := 0
	currentAmount := 50

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dial, err := NewDial(scanner.Text())
		if err != nil {
			return 0, fmt.Errorf("error creating dial: %w", err)
		}

		switch dial.Direction {
		case LEFT:
			currentAmount -= dial.Amount
			currentAmount %= 100
			if currentAmount < 0 {
				currentAmount += 100
			}
		case RIGHT:
			currentAmount += dial.Amount
			currentAmount %= 100
		}

		if currentAmount == 0 {
			answer++
		}
	}

	return answer, nil
}

func Part2(file io.Reader) (int, error) {
	answer := 0
	currentAmount := 50

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dial, err := NewDial(scanner.Text())
		if err != nil {
			return 0, fmt.Errorf("error creating dial: %w", err)
		}

		prev := currentAmount

		switch dial.Direction {
		case RIGHT:
			currentAmount += dial.Amount
			answer += currentAmount / 100
			currentAmount %= 100
		case LEFT:
			currentAmount -= dial.Amount
			for currentAmount < 0 {
				currentAmount += 100
				answer++
			}
			if currentAmount == 0 {
				answer++
			}
			if prev == 0 {
				answer--
			}
		}
	}

	return answer, nil
}

////////////////////////////
// DIAL
////////////////////////////

type Direction rune

const (
	LEFT  Direction = 'L'
	RIGHT Direction = 'R'
)

type Dial struct {
	Direction Direction
	Amount    int
}

func NewDial(text string) (Dial, error) {
	if len(text) < 2 {
		return Dial{}, fmt.Errorf("dial text must be 2 or larger, got: %v", text)
	}

	directionStr := text[0]
	var direction Direction
	switch directionStr {
	case 'L':
		direction = LEFT
	case 'R':
		direction = RIGHT
	default:
		return Dial{}, fmt.Errorf("direction was not 'L' or 'R', got %v", directionStr)
	}

	amountStr := text[1:]
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		return Dial{}, fmt.Errorf("error converting amount: %w", err)
	}

	return Dial{
		Direction: direction,
		Amount:    amount,
	}, nil
}

func (d Dial) String() string {
	return fmt.Sprintf("%c%d", d.Direction, d.Amount)
}
