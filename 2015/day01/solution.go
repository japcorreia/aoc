package main

import (
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func solution_1() {

	dat, err := os.ReadFile("./input1")
	check(err)
	directions := string(dat)
	current_floor := 0

	for i := 0; i < len(directions); i++ {
		up_down := directions[i]
		if up_down == '(' {
			current_floor++
		} else if up_down == ')' {
			current_floor--
		}
	}
	fmt.Printf("Solution 1: %d\n", current_floor)
}

func solution_2() {
	dat, err := os.ReadFile("./input1")
	check(err)
	dir := string(dat)
	cur_floor := 0

	for i := 0; i < len(dir); i++ {

		switch dir[i] {
		case '(':
			cur_floor++
		case ')':
			cur_floor--
		}

		if cur_floor < 0 {
			fmt.Printf("Solution 2: %d\n", i+1) //Offset fix
			break
		}
	}
}

func main() {
	solution_1()
	solution_2()
}
