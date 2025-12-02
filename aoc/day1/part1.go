package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Part1 calculates the total problem sets created across all four years (2024-2027)
// from the Claude Elves story. The input is the days 1-25.
func Part1(inputPath string) (int, error) {
	// Read the days from input (just to follow AoC convention)
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
	// Each year shows: elves, sets per elf per day, total problem sets
	scenarios := []struct {
		year         int
		problemSets  int
		description  string
	}{
		{2024, 3_850, "No AI - 200 elves × 0.77 sets/elf/day × 25 days"},
		{2025, 131_250, "Claude Code - ramping from 200 to 350 elves × 15 sets/elf/day"},
		{2026, 6_500_000, "Claude 2.0 (AGI) - 350 elves reviewing"},
		{2027, 100_000_000, "Claude 3.0 (ASI) - 0 elves, full automation"},
	}

	totalProblemSets := 0
	for _, scenario := range scenarios {
		totalProblemSets += scenario.problemSets
	}

	return totalProblemSets, nil
}
