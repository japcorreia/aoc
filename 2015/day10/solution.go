package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func lookAndSay(in string) (string, error) {
	if len(in) == 0 {
		return in, fmt.Errorf("empty input")
	}
	output := make([]byte, 0)
	for i := 0; i < len(in); i++ {
		currentChar := in[i]
		count := 1
		for i < len(in)-1 && in[i] == in[i+1] {
			count++
			i++
		}
		c := []byte(strconv.Itoa(count))
		output = append(output, c...)
		output = append(output, currentChar)
	}
	return string(output), nil
}

func solve1(data string, processTimes int) (int, error) {
	var err error
	var sequence string = data
	for i := 0; i < processTimes; i++ {
		sequence, err = lookAndSay(sequence)
		if err != nil {
			return 0, err
		}
	}
	return len(sequence), nil
}

func solve2(data string, processTimes int) (int, error) {
	return solve1(data, processTimes)
}

func run(inputPath string) error {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("read input file %q: %w", inputPath, err)
	}
	trimmedData := strings.TrimSpace(string(data))
	processTimes := 40

	part1, err := solve1(trimmedData, processTimes)
	if err != nil {
		return fmt.Errorf("solve part 1: %w", err)
	}
	processTimes = 50
	part2, err := solve2(trimmedData, processTimes)
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
