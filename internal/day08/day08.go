package day08

import (
	"fmt"
	"regexp"
	"strings"
)

func Star1(inputString string) error {
	if inputString == "" {
		return fmt.Errorf("input file expected")
	}

	input, err := parseInput(inputString)
	if err != nil {
		return fmt.Errorf("Cannot parse input: %v", err)
	}

	m := map[string]Mapping{}
	for _, mapping := range input.mappings {
		m[mapping.src] = mapping
	}

	current := m["AAA"]
	steps := 0

outer:
	for {
		for _, dir := range input.directions {
			steps += 1
			if dir == LEFT {
				current = m[current.left]
			} else {
				current = m[current.right]
			}

			if current.src == "ZZZ" {
				break outer
			}
		}
	}

	fmt.Println(steps)
	return nil
}

func Star2(inputString string) error {
	if inputString == "" {
		return fmt.Errorf("input file expected")
	}

	input, err := parseInput(inputString)
	if err != nil {
		return fmt.Errorf("Cannot parse input: %v", err)
	}

	m := map[string]Mapping{}
	currents := []Mapping{}
	for _, mapping := range input.mappings {
		m[mapping.src] = mapping
		if strings.HasSuffix(mapping.src, "A") {
			currents = append(currents, mapping)
		}
	}

	steps := 0

	// Start at all nodes ending with A. Stop when all nodes
	// end with Z. This is slow. Maybe quicker to find steps/
	// period for each node and then some common factor or
	// something.
outer:
	for {
		for _, dir := range input.directions {
			steps += 1
			for currentIdx, current := range currents {
				if dir == LEFT {
					currents[currentIdx] = m[current.left]
				} else {
					currents[currentIdx] = m[current.right]
				}
			}

			endsWithZ := func(t Mapping) bool {
				return strings.HasSuffix(t.src, "Z")
			}

			if all(currents, endsWithZ) {
				break outer
			}
		}
	}

	fmt.Println(steps)
	return nil
}

func all[T any](input []T, f func(T) bool) bool {
	for _, val := range input {
		if !f(val) {
			return false
		}
	}
	return true
}

type Input struct {
	directions []Direction
	mappings   []Mapping
}

type Mapping struct {
	src   string
	left  string
	right string
}

type Direction int

const (
	LEFT Direction = iota
	RIGHT
)

func parseInput(input string) (Input, error) {
	parts := strings.Split(input, "\n\n")
	directionStrs := strings.TrimSpace(parts[0])
	mapStrs := strings.TrimSpace(parts[1])

	directions := []Direction{}
	for _, c := range directionStrs {
		if c == 'L' {
			directions = append(directions, LEFT)
		} else {
			directions = append(directions, RIGHT)
		}
	}

	mapRe := regexp.MustCompile("([A-Z]+) = \\(([A-Z]+), ([A-Z]+)\\)")
	mapLines := strings.Split(mapStrs, "\n")
	mappings := []Mapping{}
	for _, line := range mapLines {
		matches := mapRe.FindStringSubmatch(line)
		if matches == nil {
			return Input{}, fmt.Errorf("Line '%v' is not a valid mapping", line)
		}

		src := matches[1]
		left := matches[2]
		right := matches[3]
		mappings = append(mappings, Mapping{src, left, right})
	}

	return Input{directions, mappings}, nil
}
