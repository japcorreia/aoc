package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func solve1(data string) (int, error) {
	codeChars := 0
	memoryChars := 0
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		codeChars += len(line)
		for i := 0; i < len(line); i++ {
			if line[i] == '"' {
				continue
			} else if line[i] == '\\' {
				switch line[i+1] {
				case '\\', '"':
					i++
				case 'x':
					i += 3
				}
			}
			memoryChars++
		}
	}
	return codeChars - memoryChars, nil
}

func solve2(data string) (int, error) {
	codeChars := 0
	encodedChars := 0
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		codeChars += len(line)
		encodedChars += 2 // The addition of 2 " per line
		for i := 0; i < len(line); i++ {
			if line[i] == '"' || line[i] == '\\' {
				encodedChars += 2
				continue
			}
			encodedChars++
		}
	}
	return encodedChars - codeChars, nil
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
