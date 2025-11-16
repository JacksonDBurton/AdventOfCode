// Package day2 solves AOC24 day2 problems.
// -Red-Nosed Reports
package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RedNosedReports [][]int

func ParseFile(questionFile *os.File) (rReport RedNosedReports) {
	scanner := bufio.NewScanner(questionFile)

	reportCount := 0
	for scanner.Scan() {
		rReport = append(rReport, []int{})
		splitLine := strings.SplitSeq(scanner.Text(), " ")
		for v := range splitLine {
			n, err := strconv.Atoi(v)
			if err != nil {
				panic(fmt.Sprintf("Could not convert char: %v", n))
			}
			rReport[reportCount] = append(rReport[reportCount], n)
		}
		reportCount++
	}

	return rReport
}
