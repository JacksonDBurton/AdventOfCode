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

func (rReport RedNosedReports) Part1() (s string) {
	validReports := 0
	for _, report := range rReport {
		if len(report) < 2 {
			continue
		}
		// increase holds if the report should be increasing or decreasing
		// grad holds if the report is increasing or decreasing gradualy
		increase, grad := false, true
		if report[1] > report[0] {
			increase = true
		}
		for i := 1; i < len(report); i++ {
			diff := report[i] - report[i-1]
			if !increase {
				diff *= -1
			}

			if diff > 3 || diff < 1 {
				grad = false
				break
			}
		}
		if grad {
			validReports++
		}
	}
	return strconv.Itoa(validReports)
}
