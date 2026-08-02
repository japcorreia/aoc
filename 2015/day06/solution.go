package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const gridSize = 1000

type point struct {
	x int
	y int
}

type Grid [gridSize][gridSize]bool
type BGrid [gridSize][gridSize]int

type operation uint8

const (
	off operation = iota
	on
	toggle
)

func nextState(state bool, op operation) bool {
	switch op {
	case off:
		return false
	case on:
		return true
	case toggle:
		return !state
	default:
		return state
	}
}

func nextBrightness(brightness int, op operation) int {
	switch op {
	case off:
		if brightness > 0 {
			return brightness - 1
		}
		return 0
	case on:
		return brightness + 1
	case toggle:
		return brightness + 2
	default:
		return brightness
	}
}

func createPoint(value string) (point, error) {
	coords := strings.Split(value, ",")
	if len(coords) != 2 {
		return point{}, fmt.Errorf("invalid point %q", value)
	}

	x, err := strconv.Atoi(coords[0])
	if err != nil {
		return point{}, fmt.Errorf("invalid x coordinate %q: %w", coords[0], err)
	}

	y, err := strconv.Atoi(coords[1])
	if err != nil {
		return point{}, fmt.Errorf("invalid y coordinate %q: %w", coords[1], err)
	}

	if x < 0 || x >= gridSize || y < 0 || y >= gridSize {
		return point{}, fmt.Errorf("point outside grid: %d,%d", x, y)
	}

	return point{x: x, y: y}, nil
}

func applyState(grid *Grid, start, end point, op operation) {
	for y := start.y; y <= end.y; y++ {
		for x := start.x; x <= end.x; x++ {
			grid[y][x] = nextState(grid[y][x], op)
		}
	}
}

func applyBrightness(grid *BGrid, start, end point, op operation) {
	for y := start.y; y <= end.y; y++ {
		for x := start.x; x <= end.x; x++ {
			grid[y][x] = nextBrightness(grid[y][x], op)
		}
	}
}

func solve1(data string) (int, error) {
	result := 0
	commands := strings.Split(data, "\n")
	var grid Grid
	for _, com := range commands {
		args := strings.Fields(com)
		switch args[0] {
		case "turn":
			switch args[1] {
			case "off":
				start, err := createPoint(args[2])
				if err != nil {
					return -1, err
				}
				end, err := createPoint(args[4])
				if err != nil {
					return -1, err
				}
				applyState(&grid, start, end, off)
			case "on":
				start, err := createPoint(args[2])
				if err != nil {
					return -1, err
				}
				end, err := createPoint(args[4])
				if err != nil {
					return -1, err
				}
				applyState(&grid, start, end, on)
			}
		case "toggle":
			start, err := createPoint(args[1])
			if err != nil {
				return -1, err
			}
			end, err := createPoint(args[3])
			if err != nil {
				return -1, err
			}
			applyState(&grid, start, end, toggle)

		}
	}

	for y := 0; y < gridSize; y++ {
		for x := 0; x < gridSize; x++ {
			if grid[y][x] {
				result++
			}
		}
	}

	return result, nil
}

func solve2(data string) (int, error) {
	result := 0
	commands := strings.Split(data, "\n")
	var grid BGrid
	for _, com := range commands {
		args := strings.Fields(com)
		switch args[0] {
		case "turn":
			switch args[1] {
			case "off":
				start, err := createPoint(args[2])
				if err != nil {
					return -1, err
				}
				end, err := createPoint(args[4])
				if err != nil {
					return -1, err
				}
				applyBrightness(&grid, start, end, off)
			case "on":
				start, err := createPoint(args[2])
				if err != nil {
					return -1, err
				}
				end, err := createPoint(args[4])
				if err != nil {
					return -1, err
				}
				applyBrightness(&grid, start, end, on)
			}
		case "toggle":
			start, err := createPoint(args[1])
			if err != nil {
				return -1, err
			}
			end, err := createPoint(args[3])
			if err != nil {
				return -1, err
			}
			applyBrightness(&grid, start, end, toggle)

		}
	}

	for y := 0; y < gridSize; y++ {
		for x := 0; x < gridSize; x++ {
			result += grid[y][x]
		}
	}

	return result, nil
}

func run(inputPath string) error {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("read input file %q: %w", inputPath, err)
	}

	trimmedData := strings.TrimSpace(string(data))

	part1, err := solve1(trimmedData)
	if err != nil {
		return fmt.Errorf("solve part 1: %w", err)
	}

	part2, err := solve2(trimmedData)
	if err != nil {
		return fmt.Errorf("solve part 2: %w", err)
	}

	fmt.Printf("Solution 1: %d\n", part1)
	fmt.Printf("Solution 2: %d\n", part2)

	return nil
}

func main() {
	inputPath := flag.String("input", "input1", "path to the puzzle input file")
	flag.Parse()

	if err := run(*inputPath); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
