# AI Elves - Advent of Code 2025

Advent of Code solutions in Go, December 2025.

---

## Quick Start

```bash
# Build
go build -o aoc-runner ./cmd/

# Run all solutions
./aoc-runner

# Run specific day
./aoc-runner -day=1

# Run specific part
./aoc-runner -day=1 -part=2
```

## Current Progress

**December 2nd, 2025**

```
🎄 AI Elves - Advent of Code Runner
======================================================================

✅ Day 1 Part 1: 106635100 (256.208µs)
✅ Day 1 Part 2: 213270200 (18.667µs)

⏱️  Total time: 331.041µs

The story about AI obsolescence is now a solvable problem with integer answers.
Read the full narrative: Day1Part1.md
======================================================================
```

| Day | Part 1 | Part 2 | Notes |
|-----|--------|--------|-------|
| 1 | ⭐ | ⭐ | [The Claude Elves](Day1Part1.md) |
| 2 | ⬜ | ⬜ | |
| 3 | ⬜ | ⬜ | |
| 4 | ⬜ | ⬜ | |
| 5 | ⬜ | ⬜ | |
| ... | | | |

---

## Project Structure

```
ai_elves/
├── aoc/
│   ├── day1/              # Days 1-25 solutions
│   │   ├── part1.go
│   │   └── part2.go
│   ├── day2/              # (upcoming)
│   └── ...
├── cmd/
│   └── main.go            # Runner
├── inputs/
│   ├── day1_input.txt
│   └── ...
├── Day1Part1.md           # Day 1 problem statement
└── go.mod
```

---

## Architecture

**Library-First Design**: Each solution is a library function returning `(int, error)`. The unified runner in `cmd/main.go` imports and executes all registered solutions.

**Pattern**:
- Solutions live in `aoc/dayN/part{1,2}.go`
- Input files in `inputs/dayN_input.txt`
- Register in `cmd/main.go` solvers slice
- No main() functions, no package conflicts

Designed to scale.

---

## Day 1: The Claude Elves

Problem statement: [Day1Part1.md](Day1Part1.md)

- **Part 1**: 106,635,100
- **Part 2**: 213,270,200

Input: Days 1-25.

---

## Next Steps

Days 2-25 TBD. The architecture is ready to scale.

Stay tuned.

---

## Requirements

- Go 1.24.5+
- No external dependencies

---

## Additional Documentation

- **[Day1Part1.md](Day1Part1.md)** - Day 1 problem statement (narrative format)

---

*Advent of Code 2025*
