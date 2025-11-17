package main

import (
	"fmt"
	"os"

	"AOC24/internal/day03"
	"AOC24/internal/day1"
	"AOC24/internal/day2"
)

func main() {
	fmt.Println("=====DAY 1=====")
	f, err := os.Open("./internal/day1/day1_historian_hysteria_input.txt")
	if err != nil {
		panic("Open day1 historian hysteria input failed")
	}

	hList := day1.ParseFile(f)
	f.Close()
	fmt.Printf("Day1 Q1: %v\n", hList.SolveListDistance())
	fmt.Printf("Day1 Q2: %v\n", hList.SimilarityScore())

	fmt.Println("=====DAY 2=====")
	f, err = os.Open("./internal/day2/day2_red-nosed_reports_input.txt")
	if err != nil {
		panic("Open day2 Red-Nosed Reports input failed")
	}

	rReport := day2.ParseFile(f)
	f.Close()
	fmt.Printf("Day2 P1: %v\n", rReport.Part1())
	fmt.Printf("Day2 P2: %v\n", rReport.Part2())

	fmt.Println("=====DAY 3=====")
	f, err = os.Open("./internal/day03/day03_mull_it_over_input.txt")
	if err != nil {
		panic("Open day03 Mull it Over input failed")
	}

	mParse := day03.ParseFile(f)
	f.Close()
	fmt.Printf("Day03 P1: %v\n", mParse.Part1())
}
