package day05

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func Star1(input string) error {
	if input == "" {
		return fmt.Errorf("input file expected")
	}

	data := parseInput(input)
	minOutput := math.MaxInt
	for _, val := range data.seeds {
		for _, mapping := range data.mappings {
			val = mapping.MapNumber(val)
		}
		if val < minOutput {
			minOutput = val
		}
	}

	fmt.Println(minOutput)
	return nil
}

func Star2(input string) error {
	if input == "" {
		return fmt.Errorf("input file expected")
	}

	data := parseInput(input)
	inputRanges := data.SeedRanges()
	for _, mapping := range data.mappings {
		outputRanges := []Range{}
		for _, inputRange := range inputRanges {
			output := mapping.MapRange(inputRange)
			for _, r := range output {
				outputRanges = append(outputRanges, r)
			}
		}
		inputRanges = outputRanges
	}

	minOutput := math.MaxInt
	for _, outputRange := range inputRanges {
		if outputRange.start < minOutput {
			minOutput = outputRange.start
		}
	}

	fmt.Println(minOutput)
	return nil
}

type Data struct {
	seeds    []int
	mappings []Mapping
}
type Mapping struct {
	name   string
	ranges [][]int
}
type Range struct {
	start int
	len   int
}

func (data *Data) SeedRanges() []Range {
	seeds := data.seeds
	seedRanges := []Range{}
	for i := 0; i < len(seeds); i += 2 {
		seedRanges = append(seedRanges, Range{seeds[i], seeds[i+1]})
	}
	return seedRanges
}

func (mapping *Mapping) MapNumber(n int) int {
	for _, r := range mapping.ranges {
		destStart := r[0]
		srcStart := r[1]
		rangeLen := r[2]

		if n >= srcStart && n < srcStart+rangeLen {
			return destStart + (n - srcStart)
		}
	}
	// If number not in any range, it stays the same.
	return n
}

func (mapping *Mapping) MapRange(r Range) []Range {
	inputRanges := []Range{r}
	outputRanges := []Range{}

	// old style loop because inputRanges can be
	// appended to in each iteration and we want to
	// work through any of those too.
	for i := 0; i < len(inputRanges); i++ {
		inputRange := inputRanges[i]
		inputRangeEnd := inputRange.start + inputRange.len

		foundMapping := false
		for _, r := range mapping.ranges {
			rangeLen := r[2]

			srcStart := r[1]
			srcEnd := srcStart + rangeLen
			destStart := r[0]

			if inputRange.start >= srcStart && inputRangeEnd <= srcEnd {
				foundMapping = true
				// inputRange fully encompassed by src
				// input      |-----|
				// source |--------------|
				outputRanges = append(outputRanges, Range{destStart + (inputRange.start - srcStart), inputRange.len})
			} else if inputRange.start < srcStart && inputRangeEnd > srcEnd {
				foundMapping = true
				// inputRange larger than src range
				// input  |-----------|
				// source    |-----|
				inputRanges = append(inputRanges, Range{inputRange.start, srcStart - inputRange.start})
				outputRanges = append(outputRanges, Range{destStart, rangeLen})
				inputRanges = append(inputRanges, Range{srcEnd, inputRangeEnd - srcEnd})
			} else if inputRange.start < srcStart && inputRangeEnd > srcStart {
				foundMapping = true
				// inputRange overlaps start of src range
				// input  |------|
				// source    |--------|
				inputRanges = append(inputRanges, Range{inputRange.start, srcStart - inputRange.start})
				outputRanges = append(outputRanges, Range{destStart, inputRangeEnd - srcStart})
			} else if inputRange.start < srcEnd && inputRangeEnd > srcEnd {
				foundMapping = true
				// inputRange overlaps end of src range
				// input       |------|
				// source |--------|
				outputRanges = append(outputRanges, Range{destStart + (inputRange.start - srcStart), srcEnd - inputRange.start})
				inputRanges = append(inputRanges, Range{srcEnd, inputRangeEnd - srcEnd})
			}
		}

		// If nothing matched our input range, it becomes an output range
		if !foundMapping {
			outputRanges = append(outputRanges, inputRange)
		}
	}

	return outputRanges
}

func parseInput(input string) Data {
	sections := strings.Split(strings.TrimSpace(input), "\n\n")
	seeds := digitsInLine(sections[0])
	mappings := []Mapping{}

	for _, section := range sections[1:] {
		lines := strings.Split(strings.TrimSpace(section), "\n")
		name := lines[0]
		ranges := [][]int{}
		for _, line := range lines[1:] {
			digits := digitsInLine(line)
			ranges = append(ranges, digits)
		}
		mappings = append(mappings, Mapping{name, ranges})
	}

	return Data{seeds, mappings}
}

var digitRe = *regexp.MustCompile("[0-9]+")

func digitsInLine(line string) []int {
	digitStrings := digitRe.FindAllString(line, -1)
	digits := []int{}
	for _, digitString := range digitStrings {
		n, _ := strconv.Atoi(digitString)
		digits = append(digits, n)
	}
	return digits
}
