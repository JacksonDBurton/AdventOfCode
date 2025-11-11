// Package day1 solves AOC24 day1 problems.
// -Historian Hysteria
package day1

import (
	"slices"
)

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
