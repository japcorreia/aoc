package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const passwordLength = 8

func isForbidden(c byte) bool {
	return c == 'i' || c == 'o' || c == 'l'
}

func isValidPassword(password []byte) bool {
	size := len(password)
	hasStraight := false
	pairs := make(map[byte]struct{})
	for i := 0; i < size; i++ {
		if isForbidden(password[i]) {
			return false
		}
		if i+2 < size && password[i] == password[i+1]-1 && password[i] == password[i+2]-2 {
			hasStraight = true
		}
		if i+1 < size && password[i] == password[i+1] {
			pairs[password[i]] = struct{}{}
		}
	}

	return hasStraight && len(pairs) > 1
}

func skipForbidden(password []byte) {
	for i, c := range password {
		if !isForbidden(c) {
			continue
		}

		// i->j, l->m, o->p
		password[i]++

		for j := i + 1; j < len(password); j++ {
			password[j] = 'a'
		}
		return
	}
}

func incrementPassword(password []byte) {
	size := len(password)
	for i := size - 1; i >= 0; i-- {
		if password[i] == 'z' {
			password[i] = 'a'
			continue
		}

		password[i]++
		return
	}
}

func solve1(data string) (string, error) {
	if len(data) != passwordLength {
		return "", fmt.Errorf("password requires %d letters, got %s", passwordLength, data)
	}

	password := []byte(data)

	for _, c := range password {
		if c < 'a' || c > 'z' {
			return "", fmt.Errorf("password must contain only lowercase letters")
		}
	}

	for {
		incrementPassword(password)
		skipForbidden(password)

		if isValidPassword(password) {
			return string(password), nil
		}
	}
}

func solve2(data string) (string, error) {
	return solve1(data)
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

	part2, err := solve2(part1)
	if err != nil {
		return fmt.Errorf("solve part 2: %w", err)
	}

	fmt.Printf("Solution 1: %s\n", part1)
	fmt.Printf("Solution 2: %s\n", part2)

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
