// Package day1 solves AOC24 day1 problems.
// -Historian Hysteria
package day1

import (
	"encoding/csv"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func SolveQuestion(questionFile *os.File) (answer int) {
	scanner := csv.NewReader(questionFile)
	scanner.Comma = ' '
	scanner.TrimLeadingSpace = true
	questionList, err := scanner.ReadAll()
	if err != nil {
		panic("Failed to read question list")
	}

	x, y := []int{}, []int{}
	for _, v := range questionList {
		if i, err := strconv.Atoi(v[0]); err == nil {
			x = append(x, i)
		} else {
			panic(fmt.Sprintf("Failed to convert Atoi: %v", v[0]))
		}
		if i, err := strconv.Atoi(v[1]); err == nil {
			y = append(y, i)
		} else {
			panic(fmt.Sprintf("Failed to convert Atoi: %v", v[0]))
		}
	}

	return SolveListDistance(x, y)
}

func SolveListDistance(x []int, y []int) (totalDistance int) {
	slices.Sort(x)
	slices.Sort(y)

	for i := range x {
		product := x[i] - y[i]
		if product < 0 {
			product *= -1
		}
		totalDistance += product
	}
	return totalDistance
}
