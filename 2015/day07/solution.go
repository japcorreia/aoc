package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	op    string
	left  string
	right string
}

func isBinaryOp(op string) bool {
	return op == "AND" || op == "OR" || op == "RSHIFT" || op == "LSHIFT"
}

func parseInstructions(booklet []string) (map[string]instruction, error) {
	circuit := make(map[string]instruction)
	for lineNumber, line := range booklet {
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		var output string
		var inst instruction

		switch {
		// 123 -> x
		case len(args) == 3 && args[1] == "->":
			output = args[2]
			inst = instruction{
				op:   "ASSIGN",
				left: args[0],
			}

		// NOT e -> f
		case len(args) == 4 && args[0] == "NOT" && args[2] == "->":
			output = args[3]
			inst = instruction{
				op:   args[0],
				left: args[1],
			}

		// x AND y -> z
		// x OR y -> z
		// x RSHIFT y -> z
		// x LSHIFT y -> z
		case len(args) == 5 && args[3] == "->" && isBinaryOp(args[1]):
			output = args[4]
			inst = instruction{
				op:    args[1],
				left:  args[0],
				right: args[2],
			}
		default:
			return nil, fmt.Errorf(
				"line %d: malformed instruction %q",
				lineNumber+1,
				line,
			)
		}

		if _, exists := circuit[output]; exists {
			return nil, fmt.Errorf(
				"line %d: wire %q is assigned more than once",
				lineNumber+1,
				output,
			)

		}

		circuit[output] = inst

	}
	return circuit, nil
}

func parseLiteral(s string) (uint16, bool) {
	value, err := strconv.ParseUint(s, 10, 16)
	if err != nil {
		return 0, false
	}
	return uint16(value), true
}

func resolve(value string, circuit map[string]instruction, wireCache map[string]uint16) (uint16, error) {
	if literal, ok := parseLiteral(value); ok {
		return literal, nil
	}
	return computeWire(value, circuit, wireCache)
}

func computeWire(wire string, circuit map[string]instruction, wireCache map[string]uint16) (uint16, error) {
	if value, exists := wireCache[wire]; exists {
		return value, nil
	}

	inst, exists := circuit[wire]
	if !exists {
		return 0, fmt.Errorf("wire %q has no instruction", wire)
	}

	var result uint16

	switch inst.op {
	case "ASSIGN":
		value, err := resolve(inst.left, circuit, wireCache)
		if err != nil {
			return 0, err
		}

		result = value

	case "NOT":
		value, err := resolve(inst.left, circuit, wireCache)
		if err != nil {
			return 0, err
		}

		result = ^value

	case "AND":
		left, err := resolve(inst.left, circuit, wireCache)
		if err != nil {
			return 0, err
		}

		right, err := resolve(inst.right, circuit, wireCache)
		if err != nil {
			return 0, err
		}

		result = left & right

	case "OR":
		left, err := resolve(inst.left, circuit, wireCache)
		if err != nil {
			return 0, err
		}

		right, err := resolve(inst.right, circuit, wireCache)
		if err != nil {
			return 0, err
		}

		result = left | right

	case "LSHIFT":
		left, err := resolve(inst.left, circuit, wireCache)
		if err != nil {
			return 0, err
		}

		right, err := resolve(inst.right, circuit, wireCache)
		if err != nil {
			return 0, err
		}

		result = left << right

	case "RSHIFT":
		left, err := resolve(inst.left, circuit, wireCache)
		if err != nil {
			return 0, err
		}

		right, err := resolve(inst.right, circuit, wireCache)
		if err != nil {
			return 0, err
		}

		result = left >> right

	default:
		return 0, fmt.Errorf(
			"wire %q has unknown operation %q",
			wire,
			inst.op,
		)
	}

	wireCache[wire] = result
	return result, nil
}

func solve1(data string, targetWire string) (uint16, error) {
	booklet := strings.Split(data, "\n")

	circuit, err := parseInstructions(booklet)
	if err != nil {
		return 0, err
	}

	cache := make(map[string]uint16)
	return computeWire(targetWire, circuit, cache)
}

func solve2(data string, targetWire string, overrideWire string, overload uint16) (uint16, error) {
	booklet := strings.Split(data, "\n")
	circuit, err := parseInstructions(booklet)
	if err != nil {
		return 0, err
	}

	circuit[overrideWire] = instruction{
		op:   "ASSIGN",
		left: strconv.FormatUint(uint64(overload), 10),
	}

	cache := make(map[string]uint16)
	return computeWire(targetWire, circuit, cache)
}

func run(inputPath string) error {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("read input file %q: %w", inputPath, err)
	}

	trimmedData := strings.TrimSpace(string(data))

	part1, err := solve1(trimmedData, "a")
	if err != nil {
		return fmt.Errorf("solve part 1: %w", err)
	}

	part2, err := solve2(trimmedData, "a", "b", part1)
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
