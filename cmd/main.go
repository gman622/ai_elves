package main

import (
	"ai_elves/aoc/day1"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type DayPart struct {
	day   int
	part  int
	solve func(inputPath string) (int, error)
}

var solvers = []DayPart{
	{1, 1, day1.Part1},
	{1, 2, day1.Part2},
}

func main() {
	dayFlag := flag.Int("day", 0, "Day to run (0 for all)")
	partFlag := flag.Int("part", 0, "Part to run (0 for all parts of the day)")
	flag.Parse()

	// Filter solvers based on flags
	var toRun []DayPart
	if *dayFlag == 0 {
		toRun = solvers
	} else {
		for _, s := range solvers {
			if s.day == *dayFlag {
				if *partFlag == 0 || s.part == *partFlag {
					toRun = append(toRun, s)
				}
			}
		}
		if len(toRun) == 0 {
			log.Fatalf("No solutions found for day %d part %d", *dayFlag, *partFlag)
		}
	}

	// Run solvers
	fmt.Println("🎄 AI Elves - Advent of Code Runner")
	fmt.Println(strings.Repeat("=", 70))
	fmt.Println()

	totalStart := time.Now()
	for _, solver := range toRun {
		// Try part-specific input file first, fall back to day-wide file
		partSpecificPath := filepath.Join(".", "inputs", fmt.Sprintf("day%d_part%d_input.txt", solver.day, solver.part))
		dayWidePath := filepath.Join(".", "inputs", fmt.Sprintf("day%d_input.txt", solver.day))

		inputPath := dayWidePath
		if _, err := os.Stat(partSpecificPath); err == nil {
			inputPath = partSpecificPath
		}

		if _, err := os.Stat(inputPath); err != nil {
			fmt.Printf("❌ Day %d Part %d: Input file not found at %s\n", solver.day, solver.part, inputPath)
			continue
		}

		start := time.Now()
		result, err := solver.solve(inputPath)
		elapsed := time.Since(start)

		if err != nil {
			fmt.Printf("❌ Day %d Part %d: %v\n", solver.day, solver.part, err)
		} else {
			fmt.Printf("✅ Day %d Part %d: %d (%v)\n", solver.day, solver.part, result, elapsed)
		}
	}

	totalElapsed := time.Since(totalStart)
	fmt.Println()
	fmt.Printf("⏱️  Total time: %v\n", totalElapsed)
	fmt.Println()
	fmt.Println("The story about AI obsolescence is now a solvable problem with integer answers.")
	fmt.Println("Read the full narrative: Day1Part1.md")
	fmt.Println(strings.Repeat("=", 70))
}
