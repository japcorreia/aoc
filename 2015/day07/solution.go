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

var wireCache = make(map[string]uint16)

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

func computeLogicalGates(inst instruction) (uint16, error) {
	if inst.op == "NOT" {
		value, ok := parseLiteral(inst.left)
		if !ok {
			return 0, fmt.Errorf("error in parsing the value %q", inst.left)
		}
		return ^value, nil
	}

	left, ok := parseLiteral(inst.left)
	if !ok {
		return 0, fmt.Errorf("error in parsing the value %q", inst.left)
	}

	right, ok := parseLiteral(inst.right)
	if !ok {
		return 0, fmt.Errorf("error in parsing the value %q", inst.left)
	}
	switch inst.op {
	case "OR":
		return left | right, nil
	case "AND":
		return left & right, nil
	case "RSHIFT":
		return left >> right, nil
	case "LSHIFT":
		return left << right, nil
	}
	return 0, fmt.Errorf("error in parsing the value %q", inst.left)
}

func parseLiteral(s string) (uint16, bool) {
	value, err := strconv.ParseUint(s, 10, 16)
	if err != nil {
		return 0, false
	}
	return uint16(value), true
}

func computeWire(wire string, circuit map[string]instruction) uint16 {
	inst := circuit[wire]
	switch inst.op {
	case "ASSIGN":
		if value, ok := parseLiteral(inst.left); ok {
			wireCache[wire] = value
			return value
		} else if _, exists := wireCache[wire]; exists {
			return wireCache[wire]
		} else {
			return computeWire(inst.left, circuit)
		}

	case "NOT":
		var result uint16
		if _, ok := parseLiteral(inst.left); ok {
			value, err := computeLogicalGates(inst)
			if err != nil {
				panic(err)
			}
			result = value
		} else {
			if _, exists := wireCache[inst.left]; exists {
				left := strconv.FormatUint(uint64(wireCache[inst.left]), 10)
				value, err := computeLogicalGates(instruction{op: inst.op, left: left})
				if err != nil {
					panic(err)
				}
				result = value
			} else {
				left := strconv.FormatUint(uint64(computeWire(inst.left, circuit)), 10)
				value, err := computeLogicalGates(instruction{op: inst.op, left: left})
				if err != nil {
					panic(err)
				}
				result = value
			}
		}
		wireCache[wire] = result
		return result

	case "OR", "AND", "RSHIFT", "LSHIFT":
		var left, right uint16
		if value, ok := parseLiteral(inst.left); ok {
			left = value
		} else {
			if _, exists := wireCache[inst.left]; exists {
				left = wireCache[inst.left]
			} else {
				left = computeWire(inst.left, circuit)
			}
		}

		if value, ok := parseLiteral(inst.right); ok {
			right = value
		} else {
			if _, exists := wireCache[inst.right]; exists {
				right = wireCache[inst.right]
			} else {
				right = computeWire(inst.right, circuit)
			}
		}

		result, err := computeLogicalGates(instruction{op: inst.op, left: strconv.FormatUint(uint64(left), 10), right: strconv.FormatUint(uint64(right), 10)})
		if err != nil {
			panic(err)
		}
		wireCache[wire] = result
		return result

	}
	return 0
}

func solve1(data string) (uint16, error) {
	booklet := strings.Split(data, "\n")
	clear(wireCache)
	var result uint16
	circuit, err := parseInstructions(booklet)
	if err != nil {
		return 0, err
	}
	result = computeWire("a", circuit)
	return result, nil
}

func solve2(data string, wire string, overload uint16) (uint16, error) {
	booklet := strings.Split(data, "\n")
	clear(wireCache)
	wireCache[wire] = overload
	var result uint16
	circuit, err := parseInstructions(booklet)
	if err != nil {
		return 0, err
	}
	result = computeWire("a", circuit)
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

	part2, err := solve2(trimmedData, "b", part1)
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
