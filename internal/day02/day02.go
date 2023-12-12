package day02

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

	games, err := parseGames(input)
	if err != nil {
		return fmt.Errorf("Error parsing games: %v", err)
	}

	const MAX_RED = 12
	const MAX_GREEN = 13
	const MAX_BLUE = 14

	sumOfOkIds := 0
	for _, game := range games {
		if game.MaxOfColour(RED) <= MAX_RED && game.MaxOfColour(BLUE) <= MAX_BLUE && game.MaxOfColour(GREEN) <= MAX_GREEN {
			sumOfOkIds += game.id
		}
	}

	fmt.Println(sumOfOkIds)
	return nil
}

func Star2(input string) error {
	if input == "" {
		return fmt.Errorf("input file expected")
	}

	games, err := parseGames(input)
	if err != nil {
		return fmt.Errorf("Error parsing games: %v", err)
	}

	sumOfPowers := 0
	for _, game := range games {
		maxReds := game.MaxOfColour(RED)
		maxBlues := game.MaxOfColour(BLUE)
		maxGreens := game.MaxOfColour(GREEN)

		sumOfPowers += maxReds * maxBlues * maxGreens
	}

	fmt.Println(sumOfPowers)
	return nil
}

func parseGames(s string) ([]Game, error) {
	gameRe := regexp.MustCompile("[gG]ame ([0-9]+)")
	colourRe := regexp.MustCompile("([0-9]+) (red|green|blue)")
	lines := strings.Split(strings.TrimSpace(s), "\n")

	games := []Game{}
	for _, line := range lines {
		gameMatches := gameRe.FindStringSubmatch(line)
		if len(gameMatches) != 2 {
			return games, fmt.Errorf("Could not find game ID in '%v'", line)
		}

		gameId, err := strconv.Atoi(gameMatches[1])
		if err != nil {
			return games, fmt.Errorf("Could not parse %v into an int", gameMatches[1])
		}

		// I didn't even notice that ;'s separated different sets of colours that
		// were drawn. But it turns out that it's irrelevant anyway.
		colourCounts := []ColourCount{}
		allColourMatches := colourRe.FindAllStringSubmatch(line, -1)
		for _, matches := range allColourMatches {
			count, err := strconv.Atoi(matches[1])
			if err != nil {
				return games, fmt.Errorf("Could not parse colour count %v into an int", matches[1])
			}

			col := GREEN
			switch matches[2] {
			case "red":
				col = RED
			case "blue":
				col = BLUE
			}

			colourCounts = append(colourCounts, ColourCount{col, count})
		}

		games = append(games, Game{gameId, colourCounts})
	}

	return games, nil
}

type Game struct {
	id      int
	colours []ColourCount
}

func (game *Game) MaxOfColour(colour Colour) int {
	max := 0
	for _, col := range game.colours {
		if col.colour == colour && col.count > max {
			max = col.count
		}
	}
	return max
}

type ColourCount struct {
	colour Colour
	count  int
}

type Colour int

const (
	GREEN Colour = iota
	RED
	BLUE
)
