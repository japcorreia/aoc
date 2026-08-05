package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func computeSum(value any) (int, error) {
	switch value := value.(type) {
	case json.Number:
		number, err := strconv.Atoi(value.String())
		if err != nil {
			return 0, fmt.Errorf("parse number %q: %w", value, err)
		}
		return number, nil

	case []any:
		sum := 0
		for _, element := range value {
			value, err := computeSum(element)
			if err != nil {
				return 0, err
			}
			sum += value
		}
		return sum, nil

	case map[string]any:
		sum := 0
		for _, property := range value {
			value, err := computeSum(property)
			if err != nil {
				return 0, err
			}
			sum += value
		}
		return sum, nil

	case string, bool, nil:
		return 0, nil

	default:
		return 0, fmt.Errorf("unexpected JSON type %T", value)
	}
}

func computeNoRedSum(value any) (int, error) {
	switch value := value.(type) {
	case json.Number:
		number, err := strconv.Atoi(value.String())
		if err != nil {
			return 0, fmt.Errorf("parse number %q: %w", value, err)
		}
		return number, nil

	case []any:
		var sum int
		for _, element := range value {
			value, err := computeNoRedSum(element)
			if err != nil {
				return 0, err
			}
			sum += value
		}
		return sum, nil

	case map[string]any:
		for _, property := range value {
			if text, ok := property.(string); ok && text == "red" {
				return 0, nil
			}
		}

		sum := 0

		for _, property := range value {
			value, err := computeNoRedSum(property)
			if err != nil {
				return 0, err
			}

			sum += value
		}
		return sum, nil

	case string, bool, nil:
		return 0, nil

	default:
		return 0, fmt.Errorf("unexpected JSON type %T", value)
	}
}

func solve1(data string) (int, error) {
	decoder := json.NewDecoder(strings.NewReader(data))
	decoder.UseNumber()

	var root any
	if err := decoder.Decode(&root); err != nil {
		return 0, fmt.Errorf("decode JSON: %w", err)
	}

	return computeSum(root)
}

func solve2(data string) (int, error) {
	decoder := json.NewDecoder(strings.NewReader(data))
	decoder.UseNumber()

	var root any
	if err := decoder.Decode(&root); err != nil {
		return 0, fmt.Errorf("decode JSON: %w", err)
	}

	return computeNoRedSum(root)
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
