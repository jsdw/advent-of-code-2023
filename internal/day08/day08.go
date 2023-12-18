package day08

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/jsdw/advent-of-code-2023/internal/utils/mathutils"
	"github.com/jsdw/advent-of-code-2023/internal/utils/sliceutils"
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

	// This takes way too long to do "manually"; it turns out that
	// each input takes a fixed number of steps to get to a finish,
	// and loops this same number each time. So we can get all of these
	// periods (erroring if this assumption is broken) and then find
	// the lowest common multiple of them all to find out when they all
	// converge
	periods, err := findPeriodForStarts(input)
	if err != nil {
		return fmt.Errorf("Cannot find periods: %v", err)
	}

	lcm := mathutils.LCMSlice(periods)
	fmt.Println(lcm)
	return nil
}

func findPeriodForStarts(input Input) ([]int, error) {
	m := map[string]Mapping{}
	for _, mapping := range input.mappings {
		m[mapping.src] = mapping
	}

	currents := sliceutils.KeepIf(input.mappings, func(m Mapping) bool {
		return strings.HasSuffix(m.src, "A")
	})

	type CurrentData struct {
		start              string
		end                string
		current            Mapping
		stepsToFinish      int
		stepsToFinishAgain int
	}

	currentsWithData := sliceutils.Map(currents, func(m Mapping) *CurrentData {
		return &CurrentData{
			start:              m.src,
			end:                "",
			current:            m,
			stepsToFinish:      0,
			stepsToFinishAgain: 0,
		}
	})

	for _, c := range currentsWithData {
		steps := 0

	outer:
		for {
			for _, dir := range input.directions {
				steps += 1
				if dir == LEFT {
					c.current = m[c.current.left]
				} else {
					c.current = m[c.current.right]
				}

				if strings.HasSuffix(c.current.src, "Z") {
					if c.end == "" {
						c.stepsToFinish = steps
						c.end = c.current.src
					} else {
						if c.current.src != c.end {
							return []int{}, fmt.Errorf("Value %v did not loop around to itself", c.start)
						}
						c.stepsToFinishAgain = steps - c.stepsToFinish
						if c.stepsToFinish != c.stepsToFinishAgain {
							return []int{}, fmt.Errorf("Value %v did not loop to end again in same number of steps", c.start)
						}
						break outer
					}
				}
			}
		}
	}

	mapStepsToFinish := func(c *CurrentData) int { return c.stepsToFinish }
	return sliceutils.Map(currentsWithData, mapStepsToFinish), nil
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
