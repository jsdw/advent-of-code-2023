package day06

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Star1(input string) error {
	if input == "" {
		return fmt.Errorf("input file expected")
	}

	data, err := parseInput(input, star1Parser)
	if err != nil {
		return fmt.Errorf("error parsing input: %v", err)
	}

	output := 1
	for idx, time := range data.times {
		distance := data.distances[idx]
		numWinning := 0
		for d := range calcPossibleDistances(time) {
			if d > distance {
				numWinning += 1
			}
		}
		output *= numWinning
	}

	fmt.Println(output)
	return nil
}

func Star2(input string) error {
	if input == "" {
		return fmt.Errorf("input file expected")
	}

	data, err := parseInput(input, star2Parser)
	if err != nil {
		return fmt.Errorf("error parsing input: %v", err)
	}

	output := 0
	time := data.times[0]
	distance := data.distances[0]
	for d := range calcPossibleDistances(time) {
		if d > distance {
			output += 1
		}
	}

	fmt.Println(output)
	return nil
}

// Return a makeshift (and not particularly efficient)
// iterator over the possible distances.
func calcPossibleDistances(time int) <-chan int {
	c := make(chan int, 10)

	go (func() {
		for i := 0; i < time; i++ {
			speed := i
			distance := speed * (time - i)
			c <- distance
		}
		close(c)
	})()

	return c
}

type Data struct {
	times     []int
	distances []int
}

func star1Parser(ns []string) []int {
	out := []int{}
	for _, n := range ns {
		i, _ := strconv.Atoi(n)
		out = append(out, i)
	}
	return out
}

func star2Parser(ns []string) []int {
	nStr := ""
	for _, n := range ns {
		nStr += n
	}
	out, _ := strconv.Atoi(nStr)
	return []int{out}
}

func parseInput(input string, intHandler func(ns []string) []int) (Data, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	digitRe := *regexp.MustCompile("[0-9]+")

	if len(lines) != 2 {
		return Data{}, fmt.Errorf("expected input to contain 2 lines")
	}

	lineToInts := func(line string) []int {
		ns := digitRe.FindAllString(line, -1)
		return intHandler(ns)
	}

	times := lineToInts(lines[0])
	distances := lineToInts(lines[1])

	if len(times) != len(distances) {
		return Data{}, fmt.Errorf("expected same number of times and distances")
	}
	return Data{times, distances}, nil
}
