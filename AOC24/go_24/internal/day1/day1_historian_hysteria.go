// Package day1 solves AOC24 day1 problems.
// -Historian Hysteria
package day1

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func SolveQuestion(questionFile *os.File) (answer string) {
	answerList := []int{}
	scanner := bufio.NewScanner(questionFile)
	for scanner.Scan() {
		line := scanner.Text()
		pair := strings.Split(line, " ")
		if len(pair) != 2 {
			panic("Pair length not 2")
		}

		xInt, err := strconv.Atoi(pair[0])
		if err != nil {
			panic(fmt.Sprintf("Failed to convert x: %v", pair[0]))
		}
		yInt, err := strconv.Atoi(pair[1])
		if err != nil {
			panic(fmt.Sprintf("Failed to convert y: %v", pair[1]))
		}
		x := splitDigits(xInt)
		y := splitDigits(yInt)
		answerList = append(answerList, SolveListDistance(x, y))
	}

	return strings.Join(strings.Split(fmt.Sprint(answerList), ", "), "")
}

func splitDigits(n int) (digits []int) {
	for n > 0 {
		dig := n % 10
		n /= 10
		digits = append(digits, dig)
	}
	return
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
