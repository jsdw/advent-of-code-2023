package main

import (
	"flag"
	"fmt"
	"os"
	"unicode/utf8"

	"github.com/jsdw/advent-of-code-2023/internal/day01"
	"github.com/jsdw/advent-of-code-2023/internal/day02"
)

func main() {
	dayPtr := flag.Int("day", 0, "1-25 depending on how far I bother getting")
	starPtr := flag.Int("star", 0, "1 or 2 to run the solution for the first or second star")
	inputPathPtr := flag.String("input", "", "file with input data in. Can be omitted if no input for the day")

	flag.Parse()

	day := *dayPtr
	star := *starPtr
	inputPath := *inputPathPtr

	if day == 0 {
		fmt.Println("Please specify a day to run eg -day 1")
		os.Exit(1)
	}
	if star != 1 && star != 2 {
		fmt.Println("Please specify a star to run eg -star 1 or -star 2")
		os.Exit(1)
	}

	input := ""
	if inputPath != "" {
		inputBytes, err := os.ReadFile(inputPath)
		if err != nil {
			fmt.Printf("Could not read file at path %v: %v\n", inputPath, err)
			os.Exit(1)
		}
		if !utf8.Valid(inputBytes) {
			fmt.Printf("File %v is not valid UTF8\n", inputPath)
			os.Exit(1)
		}
		input = string(inputBytes)
	}

	// Our map of solutions that can be selected; nil for any missing ones.
	solutions := map[int][2](func(string) error){
		1: {day01.Star1, day01.Star2},
		2: {day02.Star1, day02.Star2},
	}

	daySolutions, foundDay := solutions[day]
	if !foundDay {
		fmt.Printf("No solutions exists for day %v\n", day)
		os.Exit(1)
	}

	daySolution := daySolutions[star-1]
	if daySolution == nil {
		fmt.Printf("No solution exists for day %v star %v\n", day, star)
		os.Exit(1)
	}

	err := daySolution(input)
	if err != nil {
		fmt.Printf("Error running day %v star %v: %v\n", day, star, err)
		os.Exit(1)
	}
}
