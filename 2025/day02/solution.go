package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func i_min(a, b int) int {
	if a < b {
		return a
	}else {
		return b
	}
}

func computePresentRibbon(a, b, c int) int {
	maxSide := a
	if b > maxSide {
		maxSide = b
	}
	if c > maxSide {
		maxSide = c
	}
	return 2 * (a + b + c - maxSide)
}

func solution1(dimsList []string) {
	totalPaper := 0
	for i := 0; i < len(dimsList)-1; i++{
		dims := strings.Split(dimsList[i], "x")
		
		l, err := strconv.Atoi(dims[0])
		check(err)
		w, err := strconv.Atoi(dims[1])
		check(err)
		h, err := strconv.Atoi(dims[2])
		check(err)
		l_w := l * w
		w_h := w * h
		h_l := h * l
		slack := i_min(i_min(l_w, w_h), h_l)
		totalPaper += slack + (2 * (l_w + w_h + h_l))
		
	}
	fmt.Printf("Solution 1: %d\n", totalPaper)
}

func solution2(dimsList []string) {
	totalRibbon := 0
	for i := 0; i < len(dimsList)-1; i++{
		dims := strings.Split(dimsList[i], "x")
		
		l, err := strconv.Atoi(dims[0])
		check(err)
		w, err := strconv.Atoi(dims[1])
		check(err)
		h, err := strconv.Atoi(dims[2])
		check(err)
		presentRibbon := computePresentRibbon(l, w, h)
		bowRibbon := l*w*h
		totalRibbon += presentRibbon + bowRibbon
		
	}
	fmt.Printf("Solution 2: %d\n", totalRibbon)
}

func main(){
	dat, err := os.ReadFile("./input1")
	check(err)
	dimsList := strings.Split(string(dat), "\n")
	solution1(dimsList)
	solution2(dimsList)
}
