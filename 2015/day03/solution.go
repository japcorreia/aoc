package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
)

type Vector2D struct {
	x int
	y int
}

var moves = map[byte]Vector2D{
	'<': {x: -1, y: 0},
	'>': {x: 1, y: 0},
	'^': {x: 0, y: 1},
	'v': {x: 0, y: -1},
}

func nextHouse(current, move Vector2D) Vector2D {
	return Vector2D{
		x: current.x + move.x,
		y: current.y + move.y,
	}
}

func getMove(direction byte, position int) (Vector2D, error) {
	move, ok := moves[direction]
	if !ok {
		return Vector2D{}, fmt.Errorf(
			"invalid direction %q at position %d",
			direction,
			position,
		)
	}
	return move, nil
}

func solve1(path []byte) (int, error) {
	current := Vector2D{}
	visited := map[Vector2D]struct{}{
		current: {},
	}

	for i, direction := range path {
		move, err := getMove(direction, i)
		if err != nil {
			return 0, err
		}

		current = nextHouse(current, move)
		visited[current] = struct{}{}
	}

	return len(visited), nil
}

func solve2(path []byte) (int, error) {
	positions := [2]Vector2D{}
	visited := map[Vector2D]struct{}{
		positions[0]: {},
	}

	for i, direction := range path {
		move, err := getMove(direction, i)
		if err != nil {
			return 0, err
		}
		actor := i % len(positions)
		positions[actor] = nextHouse(positions[actor], move)
		visited[positions[actor]] = struct{}{}
	}

	return len(visited), nil
}

func run(inputPath string) error {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("read input file %q: %w", inputPath, err)
	}

	path := bytes.TrimSpace(data)

	part1, err := solve1(path)
	if err != nil {
		return fmt.Errorf("solve part 1: %w", err)
	}

	part2, err := solve2(path)
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
