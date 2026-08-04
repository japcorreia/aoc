package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const noRoute = -1

type edge struct {
	u string
	v string
	w int
}

func getAdjacencyMatrix(edges []edge, nodes map[string]int) [][]int {
	nodes_size := len(nodes)
	matrix := make([][]int, nodes_size)
	for i := range matrix {
		matrix[i] = make([]int, nodes_size)
		for j := range matrix[i] {
			matrix[i][j] = noRoute
		}
		matrix[i][i] = 0
	}

	for _, e := range edges {
		matrix[nodes[e.u]][nodes[e.v]] = e.w
		matrix[nodes[e.v]][nodes[e.u]] = e.w
	}
	return matrix
}

func parseRoutes(routesList []string) ([]edge, map[string]int, error) {
	edges := make([]edge, len(routesList))
	nodes := make(map[string]int)

	for _, route := range routesList {
		var e edge
		args := strings.Fields(route)
		if len(args) != 5 {
			return nil, nil, fmt.Errorf("bad route: %q", route)
		}
		weigth, err := strconv.Atoi(args[4])
		if err != nil {
			return nil, nil, err
		}
		appendNode := func(n string) {
			if _, exists := nodes[n]; !exists {
				nodes[n] = len(nodes)
			}
		}
		appendNode(args[0])
		appendNode(args[2])

		e = edge{u: args[0], v: args[2], w: weigth}
		edges = append(edges, e)
	}
	return edges, nodes, nil
}

func computePaths(matrix [][]int) (int, int, error) {
	size := len(matrix)

	if size == 0 {
		return 0, 0, fmt.Errorf("no cities found")
	}

	statesSize := (1 << size)
	fullMask := statesSize - 1

	minDp := make([][]int, statesSize)
	maxDp := make([][]int, statesSize)

	for mask := 0; mask < statesSize; mask++ {
		minDp[mask] = make([]int, size)
		maxDp[mask] = make([]int, size)

		for node := 0; node < size; node++ {
			minDp[mask][node] = math.MaxInt
			maxDp[mask][node] = math.MinInt
		}
	}

	// The route can start at any city, so visinting the city has a distance of zero
	for node := 0; node < size; node++ {
		minDp[1<<node][node] = 0
		maxDp[1<<node][node] = 0
	}

	for mask := 1; mask < statesSize; mask++ {
		for node := 0; node < size; node++ {
			// If the city is not in the mask skip
			if mask&(1<<node) == 0 {
				continue
			}

			shortestDistance := minDp[mask][node]
			longestDistance := maxDp[mask][node]

			if shortestDistance == math.MaxInt && longestDistance == math.MinInt {
				continue
			}

			for nextNode := 0; nextNode < size; nextNode++ {
				// If the city is already in the mask skip
				if mask&(1<<nextNode) != 0 {
					continue
				}

				edgeDistance := matrix[node][nextNode]
				if edgeDistance == noRoute {
					continue
				}

				nextMask := mask | (1 << nextNode)

				if shortestDistance != math.MaxInt {
					candidate := shortestDistance + edgeDistance

					if candidate < minDp[nextMask][nextNode] {
						minDp[nextMask][nextNode] = candidate
					}
				}

				if longestDistance != math.MinInt {
					candidate := longestDistance + edgeDistance

					if candidate > maxDp[nextMask][nextNode] {
						maxDp[nextMask][nextNode] = candidate
					}
				}
			}
		}
	}

	shortest := math.MaxInt
	longest := math.MinInt

	for node := 0; node < size; node++ {
		if minDp[fullMask][node] < shortest {
			shortest = minDp[fullMask][node]
		}
		if maxDp[fullMask][node] > longest {
			longest = maxDp[fullMask][node]
		}
	}

	if shortest == math.MaxInt || longest == math.MinInt {
		return 0, 0, fmt.Errorf("no route visists every city")
	}

	return shortest, longest, nil
}

func solve1(data string) (int, error) {
	routesList := strings.Split(data, "\n")
	edges, nodes, err := parseRoutes(routesList)
	if err != nil {
		return 0, err
	}
	matrix := getAdjacencyMatrix(edges, nodes)
	shortest, _, err := computePaths(matrix)
	if err != nil {
		return 0, err
	}
	return shortest, nil
}

func solve2(data string) (int, error) {
	routesList := strings.Split(data, "\n")
	edges, nodes, err := parseRoutes(routesList)
	if err != nil {
		return 0, err
	}
	matrix := getAdjacencyMatrix(edges, nodes)
	_, longest, err := computePaths(matrix)
	if err != nil {
		return 0, err
	}
	return longest, nil
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
