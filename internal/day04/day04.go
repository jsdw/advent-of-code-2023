package day04

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func Star1(input string) error {
	if input == "" {
		return fmt.Errorf("input file expected")
	}

	cards := parseCards(input)
	score := 0
	for _, card := range cards {
		winning := card.left
		have := card.right
		points := 0

		for _, n := range have {
			if slices.Contains(winning, n) {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}

		score += points
	}

	fmt.Println(score)
	return nil
}

func Star2(input string) error {
	if input == "" {
		return fmt.Errorf("input file expected")
	}

	cards := parseCards(input)

	type CountedCard struct {
		card  Card
		count int
	}

	countedCards := []CountedCard{}
	for _, card := range cards {
		countedCards = append(countedCards, CountedCard{card, 1})
	}

	for idx, card := range countedCards {
		nMatches := 0
		for _, n := range card.card.right {
			if slices.Contains(card.card.left, n) {
				nMatches += 1
			}
		}

		for i := idx + 1; i <= idx+nMatches; i++ {
			countedCards[i].count += card.count
		}
	}

	numCards := 0
	for _, card := range countedCards {
		numCards += card.count
	}

	fmt.Println(numCards)
	return nil
}

type Card struct {
	id    int
	left  []int
	right []int
}

func parseCards(input string) []Card {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	cards := []Card{}
	lineRe := regexp.MustCompile("Card +([0-9]+): ([0-9 ]+)\\|([0-9 ]+)")
	digitRe := regexp.MustCompile("[0-9]+")

	getNumbers := func(s string) []int {
		ns := []int{}
		for _, nStr := range digitRe.FindAllString(s, -1) {
			n, _ := strconv.Atoi(nStr)
			ns = append(ns, n)
		}
		return ns
	}

	for _, line := range lines {
		parts := lineRe.FindStringSubmatch(line)
		if len(parts) == 0 {
			continue
		}

		id, _ := strconv.Atoi(parts[1])
		left := getNumbers(parts[2])
		right := getNumbers(parts[3])
		cards = append(cards, Card{id, left, right})
	}

	return cards
}
