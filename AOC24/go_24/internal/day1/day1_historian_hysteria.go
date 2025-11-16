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

type HistorianList struct {
	X []int
	Y []int
}

func ParseFile(questionFile *os.File) (hList HistorianList) {
	scanner := csv.NewReader(questionFile)
	scanner.Comma = ' '
	scanner.TrimLeadingSpace = true
	questionList, err := scanner.ReadAll()
	if err != nil {
		panic("Failed to read question list")
	}

	for _, v := range questionList {
		if i, err := strconv.Atoi(v[0]); err == nil {
			hList.X = append(hList.X, i)
		} else {
			panic(fmt.Sprintf("Failed to convert Atoi: %v", v[0]))
		}
		if i, err := strconv.Atoi(v[1]); err == nil {
			hList.Y = append(hList.Y, i)
		} else {
			panic(fmt.Sprintf("Failed to convert Atoi: %v", v[0]))
		}
	}

	return hList
}

func (hList HistorianList) SolveListDistance() (totalDistance int) {
	slices.Sort(hList.X)
	slices.Sort(hList.Y)

	for i := range hList.X {
		product := hList.X[i] - hList.Y[i]
		if product < 0 {
			product *= -1
		}
		totalDistance += product
	}
	return totalDistance
}
