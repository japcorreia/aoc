package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const size = 1000

type point struct {
	x int
	y int
}

type Grid [size][size]bool
type BGrid [size][size]int

/* Operations:
* 0 - turn off
* 1 - turn on
* 2 - toogle
 */

func nextState(state bool, op int) bool {
	switch op {
	case 0:
		return false
	case 1:
		return true
	case 2:
		return !state
	default:
		return false
	}
}

func nextBrightness(brightness int, op int) int {
	switch op {
	case 0:
		if brightness > 0 {
			return brightness - 1
		}
		return 0
	case 1:
		return brightness + 1
	case 2:
		return brightness + 2
	default:
		return brightness
	}
}
func createPoint(p string) (point, error) {
	coords := strings.Split(p, ",")
	px, err := strconv.Atoi(coords[0])
	if err != nil {
		return point{}, err
	}
	py, err := strconv.Atoi(coords[1])
	if err != nil {
		return point{}, err
	}
	return point{x: px, y: py}, nil
}

func applyGridState[T any](grid *T, start point, end point, op int) {
	for i := start.x; i <= end.x; i++ {
		for j := start.y; j <= end.y; j++ {
			switch g := any(grid).(type) {
			case *Grid:
				g[i][j] = nextState(g[i][j], op)
			case *BGrid:
				g[i][j] = nextBrightness(g[i][j], op)
			}
		}
	}
}

func solve1(data string) (int, error) {
	result := 0
	commands := strings.Split(data, "\n")
	var grid Grid
	for _, com := range commands {
		args := strings.Split(com, " ")
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
				applyGridState(&grid, start, end, 0)
			case "on":
				start, err := createPoint(args[2])
				if err != nil {
					return -1, err
				}
				end, err := createPoint(args[4])
				if err != nil {
					return -1, err
				}
				applyGridState(&grid, start, end, 1)
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
			applyGridState(&grid, start, end, 2)

		}
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if grid[j][i] {
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
		args := strings.Split(com, " ")
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
				applyGridState(&grid, start, end, 0)
			case "on":
				start, err := createPoint(args[2])
				if err != nil {
					return -1, err
				}
				end, err := createPoint(args[4])
				if err != nil {
					return -1, err
				}
				applyGridState(&grid, start, end, 1)
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
			applyGridState(&grid, start, end, 2)

		}
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			result += grid[j][i]
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
