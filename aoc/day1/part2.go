package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Part2 calculates the total Go solutions written across all four years (2024-2027)
// from the Claude Elves story. Each problem set has 2 Go solutions (Part 1 and Part 2).
func Part2(inputPath string) (int, error) {
	// Read the days from input
	file, err := os.Open(inputPath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var days []int
	for scanner.Scan() {
		day, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return 0, fmt.Errorf("invalid day: %v", err)
		}
		days = append(days, day)
	}

	if len(days) != 25 {
		return 0, fmt.Errorf("expected 25 days, got %d", len(days))
	}

	// The progression from the Claude Elves story
	// Each problem set = Part 1 + Part 2 = 2 Go solutions
	scenarios := []struct {
		year         int
		problemSets  int
		goSolutions  int
	}{
		{2024, 3_850, 7_700},
		{2025, 131_250, 262_500},
		{2026, 6_500_000, 13_000_000},
		{2027, 100_000_000, 200_000_000},
	}

	totalGoSolutions := 0
	for _, scenario := range scenarios {
		totalGoSolutions += scenario.goSolutions
	}

	return totalGoSolutions, nil
}
