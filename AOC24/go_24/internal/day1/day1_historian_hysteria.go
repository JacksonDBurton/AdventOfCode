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

// SolveListDistance returns the solution to Q1 of D1 for this HistorianList
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

// SimilarityScore returns the solution to Q2 of D1 for this HistorianList.
// Similarity is determined by multiplying the value of each element of list x by how
// often that element can be found in list y.
// Take the individual similarity score and add it to totalSimilarity
func (hList HistorianList) SimilarityScore() (totalSimilarity int) {
	yMap := make(map[int]int)

	for _, v := range hList.Y {
		yMap[v] += 1
	}

	for _, v := range hList.X {
		if yCount, ok := yMap[v]; ok {
			totalSimilarity += yCount * v
		}
	}

	return totalSimilarity
}
