// Package day03 solves AOC24 day03 problems.
// -Mull it Over
package day03

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type MulParse struct {
	rawString   string
	Part01Parse []string
	Part02Parse []string
}

func ParseFile(questionFile *os.File) (mParse MulParse) {
	scanner := bufio.NewScanner(questionFile)

	for scanner.Scan() {
		mParse.rawString += scanner.Text()
	}

	r := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	mParse.Part01Parse = r.FindAllString(mParse.rawString, 1000)

	r = regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
	mParse.Part02Parse = r.FindAllString(mParse.rawString, 1000)

	return mParse
}

func (mParse MulParse) Part1() (s string) {
	r := regexp.MustCompile(`\d{1,3}`)
	nums := []string{}
	for _, v := range mParse.Part01Parse {
		nums = append(nums, r.FindAllString(v, 2)...)
	}

	total := 0
	for i := 0; i < len(nums)-1; i += 2 {
		n1, err := strconv.Atoi(nums[i])
		if err != nil {
			panic(fmt.Sprintf("Failed to convert %v to int\n", nums[i]))
		}
		n2, err := strconv.Atoi(nums[i+1])
		if err != nil {
			panic(fmt.Sprintf("Failed to convert %v to int\n", nums[i+1]))
		}
		total += n1 * n2
	}
	return strconv.Itoa(total)
}

func (mParse MulParse) Part2() (s string) {
	r := regexp.MustCompile(`\d{1,3}`)
	do := true
	total := 0
	for _, v := range mParse.Part02Parse {
		if do && v == "don't()" {
			do = false
			continue
		} else if !do && v == "do()" {
			do = true
			continue
		} else if v == "don't()" || v == "do()" {
			continue
		}

		if !do {
			continue
		}

		nums := r.FindAllString(v, 2)
		n1, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(fmt.Sprintf("Failed to convert %v to int\n", nums[0]))
		}
		n2, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(fmt.Sprintf("Failed to convert %v to int\n", nums[1]))
		}
		total += n1 * n2
	}
	return strconv.Itoa(total)
}
