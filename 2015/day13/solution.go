package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type edge struct {
	u string
	v string
	w int
}

func insertToSet(m map[string]int, value string) {
	if _, exists := m[value]; !exists {
		m[value] = len(m)
	}
}

func parseGraph(data string) ([][]int, error) {
	lines := strings.Split(data, "\n")
	edges := make([]edge, len(lines))
	nodes := make(map[string]int)
	for i, line := range lines {
		args := strings.Fields(line)
		happiness, err := strconv.Atoi(args[3])
		if err != nil {
			return nil, err
		}
		if args[2] == "lose" {
			happiness *= -1
		}
		u := args[0]
		v := strings.TrimSuffix(args[len(args)-1], ".")
		insertToSet(nodes, u)
		insertToSet(nodes, v)

		edges[i] = edge{u: u, v: v, w: happiness}

	}
	size := len(nodes)
	matrix := make([][]int, size)
	for i := 0; i < size; i++ {
		matrix[i] = make([]int, size)
	}

	for _, e := range edges {
		uIdx := nodes[e.u]
		vIdx := nodes[e.v]

		// sitting together benifits both u and v at the same time
		matrix[uIdx][vIdx] += e.w
		matrix[vIdx][uIdx] += e.w
	}
	return matrix, nil
}

func computeMaxHappiness(matrix [][]int) (int, error) {
	size := len(matrix)

	if size == 0 {
		return 0, fmt.Errorf("no people to sit")
	}

	if size == 1 {
		return 0, nil
	}

	maskCount := 1 << size
	targetMask := maskCount - 1

	dp := make([][]int, maskCount)
	for mask := 0; mask < maskCount; mask++ {
		dp[mask] = make([]int, size)
		for node := 0; node < size; node++ {
			dp[mask][node] = math.MinInt
		}
	}

	dp[1][0] = 0
	for mask := 1; mask < maskCount; mask++ {
		for node := 0; node < size; node++ {
			// Already sit at the table
			if mask&(1<<node) == 0 {
				continue
			}
			maxHappiness := dp[mask][node]
			if maxHappiness == math.MinInt {
				continue
			}

			for nextNode := 0; nextNode < size; nextNode++ {
				// Not sit at the table
				if mask&(1<<nextNode) != 0 {
					continue
				}
				happiness := matrix[node][nextNode]
				nextMask := mask | (1 << nextNode)
				candidate := maxHappiness + happiness

				if candidate > dp[nextMask][nextNode] {
					dp[nextMask][nextNode] = candidate
				}
			}
		}
	}
	totalHappiness := math.MinInt

	for node := 1; node < size; node++ {
		if dp[targetMask][node] != math.MinInt {
			circleHappiness := dp[targetMask][node] + matrix[node][0]
			if circleHappiness > totalHappiness {
				totalHappiness = circleHappiness
			}
		}
	}
	return totalHappiness, nil
}

func addMeToTheTable(matrix [][]int) [][]int {
	size := len(matrix)
	s := make([]int, size)
	matrix = append(matrix, s)
	for i := 0; i < size+1; i++ {
		matrix[i] = append(matrix[i], 0)
	}
	return matrix
}

func solve1(data string) (int, error) {
	m, err := parseGraph(data)
	if err != nil {
		return 0, err
	}
	return computeMaxHappiness(m)
}

func solve2(data string) (int, error) {
	m, err := parseGraph(data)
	if err != nil {
		return 0, err
	}
	m = addMeToTheTable(m)
	return computeMaxHappiness(m)
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
