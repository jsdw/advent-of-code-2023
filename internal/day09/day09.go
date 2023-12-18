package day09

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/jsdw/advent-of-code-2023/internal/utils/sliceutils"
)

func Star1(inputString string) error {
	if inputString == "" {
		return fmt.Errorf("input file expected")
	}

	inputs := parseInputs(inputString)
	sum := 0
	for _, input := range inputs {
		sum += findNext(input)
	}

	fmt.Println(sum)
	return nil
}

func Star2(inputString string) error {
	if inputString == "" {
		return fmt.Errorf("input file expected")
	}

	inputs := parseInputs(inputString)
	sum := 0
	for _, input := range inputs {
		sum += findPrev(input)
	}

	fmt.Println(sum)
	return nil
}

func findNext(inputs []int) int {
	allDiffs := findDiffs(inputs)
	lasts := sliceutils.Map(allDiffs, func(s []int) int { return sliceutils.Last(s) })
	sum := sliceutils.Fold(lasts, 0, func(s int, n int) int { return s + n })

	return sliceutils.Last(inputs) + sum
}

func findPrev(inputs []int) int {
	allDiffs := findDiffs(inputs)
	firsts := sliceutils.Map(allDiffs, func(s []int) int { return s[0] })
	slices.Reverse(firsts)
	sum := sliceutils.Fold(firsts, 0, func(s int, n int) int { return n - s })

	return inputs[0] - sum
}

func findDiffs(inputs []int) [][]int {
	diffs := func(inputs []int) []int {
		out := []int{}
		for i := 0; i < len(inputs)-1; i++ {
			out = append(out, inputs[i+1]-inputs[i])
		}
		return out
	}

	// 1   3   6  10  15  21
	//   2   3   4   5   6
	//     1   1   1   1
	allDiffs := [][]int{diffs(inputs)}
	for {
		d := diffs(sliceutils.Last(allDiffs))
		if sliceutils.All(d, func(n int) bool { return n == 0 }) {
			break
		}
		allDiffs = append(allDiffs, d)
	}

	return allDiffs
}

func parseInputs(inputs string) [][]int {
	lines := strings.Split(strings.TrimSpace(inputs), "\n")
	re := regexp.MustCompile("-?[0-9]+")

	output := [][]int{}
	for _, line := range lines {
		ns := re.FindAllString(line, -1)
		ints := sliceutils.Map(ns, func(n string) int {
			i, err := strconv.Atoi(n)
			if err != nil {
				panic("Unable to parse string to number")
			}
			return i
		})
		output = append(output, ints)
	}

	return output
}
