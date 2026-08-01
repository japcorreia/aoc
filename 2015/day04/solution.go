package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"crypto/md5"
	"io"
)

var window int = 3 // Because 5 hex means 2 and half bytes

func md5Hash(text string) ([]byte, error){
	hasher := md5.New()
	_, err := io.WriteString(hasher, text)
	if err != nil{
		return nil, err
	}
	return hasher.Sum(nil), nil
}

func solve1(data string) (int, error){

	for i:=1; true; i++{
		test := fmt.Sprintf("%s%d", data, i)
		hash, err := md5Hash(test)
		if err != nil{
			return -1, err
		}
		if hash[0] == 0 && hash[1] == 0 && hash[2] >> 4 == 0 {
			return i, nil
		}
	}
	return 0, nil
}

func solve2(data string) (int, error){
	for i:=1; true; i++{
		test := fmt.Sprintf("%s%d", data, i)
		hash, err := md5Hash(test)
		if err != nil{
			return -1, err
		}
		if hash[0] == 0 && hash[1] == 0 && hash[2] == 0 {
			return i, nil
		}
	}
	return 0, nil
}

func run(inputPath string) error {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("read input file %q: %w", inputPath, err)
	}

	trimmedData := bytes.TrimSpace(data)
	str := string(trimmedData)

	part1, err := solve1(str)
	if err != nil {
		return fmt.Errorf("solve part 1: %w", err)
	}

	part2, err := solve2(str)
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
