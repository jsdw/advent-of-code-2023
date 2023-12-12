package day01

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func Star1(input string) error {
	if input == "" {
		return fmt.Errorf("input file expected")
	}

	isDigit := func(r rune) bool {
		return r >= '0' && r <= '9'
	}
	toInt := func(s string, fst int, last int) int {
		fstRune, _ := utf8.DecodeRuneInString(s[fst:])
		lastRune, _ := utf8.DecodeRuneInString(s[last:])

		return int((fstRune-'0')*10 + (lastRune - '0'))
	}

	lines := strings.Split(strings.TrimSpace(input), "\n")
	sum := 0
	for _, line := range lines {
		firstIdx := strings.IndexFunc(line, isDigit)
		lastIdx := strings.LastIndexFunc(line, isDigit)

		if firstIdx == -1 || lastIdx == -1 {
			return fmt.Errorf("Cannot find digits in line '%v'", line)
		}

		sum += toInt(line, firstIdx, lastIdx)
	}

	fmt.Println(sum)
	return nil
}

func Star2(input string) error {
	if input == "" {
		return fmt.Errorf("input file expected")
	}

	type digit struct {
		s string
		n int
	}
	mappings := []digit{
		{"one", 1},
		{"1", 1},
		{"two", 2},
		{"2", 2},
		{"three", 3},
		{"3", 3},
		{"four", 4},
		{"4", 4},
		{"five", 5},
		{"5", 5},
		{"six", 6},
		{"6", 6},
		{"seven", 7},
		{"7", 7},
		{"eight", 8},
		{"8", 8},
		{"nine", 9},
		{"9", 9},
	}

	lines := strings.Split(strings.TrimSpace(input), "\n")
	sum := 0

	for _, line := range lines {
		foundFst := false
		fst := 0
		last := 0

		for pos := range line {
			for _, mapping := range mappings {
				if strings.HasPrefix(line[pos:], mapping.s) {
					if !foundFst {
						fst = mapping.n
						foundFst = true
					}
					last = mapping.n
				}
			}
		}

		sum += fst*10 + last
	}

	fmt.Println(sum)
	return nil
}
