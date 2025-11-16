package main

import (
	"fmt"
	"os"

	"AOC24/internal/day1"
)

func main() {
	f, err := os.Open("./internal/day1/day1_historian_hysteria_input.txt")
	if err != nil {
		panic("Open day1 historian hysteria input failed")
	}

	fmt.Printf("Day1 Q1: %v", day1.SolveQuestion(f))
}
