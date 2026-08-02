package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"strings"
)

// Due to the lack of backreferences support in go, the solution will not use regexp
func naughtyOrNice1(str string) bool {
	vowels :=0
	doubleLetters := 0
	notAllowed := [4]string {"ab", "cd", "pq", "xy"}
	
	for i:=0; i < len(str); i++ {
		switch str[i]{
		case 'a', 'e','i','o','u':
			vowels++
		}

		if i < len(str) - 1{
			if str[i] == str[i+1] && str[i] >= 'a' && str[i] <= 'z'{
				doubleLetters++;
			}

			for _, val := range notAllowed {
				if str[i:i+2] == val{
					return false
				}
			}
		}
	}
	if vowels >= 3 && doubleLetters >=1 {
		return true
	}
	return false
}


func naughtyOrNice2(str string) bool {
	pairs := make(map[string]int)
	doublePairs := false
	sameNeighbour := 0
	if len(str) > 1 {
		pairs[str[:2]]++
	}
	for i:=1; i < len(str) - 1; i++ {
		if str[i-1:i+1] != str[i:i+2] {
			pairs[str[i:i+2]]++
		}
		if str[i-1] == str[i+1] {
			sameNeighbour++
		}
	}
	for _, v := range pairs {
		if v >= 2 {
			doublePairs = true
			break
		}
	}

	return doublePairs && sameNeighbour >= 1 
}

func solve1(data string) (int, error){
	result := 0
	strs := strings.Split(data, "\n")
	for _, str := range strs {
		if naughtyOrNice1(str) {
			result++
		}
	}
	return result, nil
}

func solve2(data string) (int, error){
	result := 0
	strs := strings.Split(data, "\n")
	for _, str := range strs{
		if naughtyOrNice2(str) {
			result++
		}
	}
	return result, nil
}

func run(inputPath string) error {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("read input file %q: %w\n", inputPath, err)
	}

	trimmedData := bytes.TrimSpace(data)

	part1, err := solve1(string(trimmedData))
	if err != nil {
		return fmt.Errorf("solve part 1: %w", err)
	}

	part2, err := solve2(string(trimmedData))
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
