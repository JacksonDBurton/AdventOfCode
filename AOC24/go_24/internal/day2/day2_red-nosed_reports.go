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

func (rReport RedNosedReports) Part2() (s string) {
	validReports := 0
	for _, report := range rReport {
		if len(report) < 2 {
			continue
		}
		// fmt.Println(report)
		// increase holds if the report should be increasing or decreasing
		// grad holds if the report is increasing or decreasing gradualy
		increase, grad, dampened := -1, true, false
		if report[1] > report[0] {
			increase = 1
		}
		for i := 1; i < len(report); i++ {
			diff := (report[i] - report[i-1]) * increase
			// fmt.Printf("%v, ", diff)

			if isNotGradual(diff) {
				if !dampened {
					if i == len(report)-1 {
						continue
					}
					if i < len(report)-1 {
						potDiff := (report[i+1] - report[i-1]) * increase
						if !isNotGradual(potDiff) {
							// fmt.Printf("%v, ", "positive damp")
							dampened = true
							report[i] = report[i-1]
							continue
						}
					}
					if i > 1 && !dampened {
						potDiff := (report[i] - report[i-2]) * increase
						if !isNotGradual(potDiff) {
							// fmt.Printf("%v, ", "negative damp")
							dampened = true
							continue
						}
					}
				}
				grad = false
				break
			}
		}
		// fmt.Println("")
		if grad {
			validReports++
		} else {
			fmt.Println(report)
		}
	}
	return strconv.Itoa(validReports)
}

func isNotGradual(i int) bool {
	return i > 3 || i < 1
}
