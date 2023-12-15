package day07

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Star1(input string) error {
	if input == "" {
		return fmt.Errorf("input file expected")
	}

	hands, err := parseInput(input, false)
	if err != nil {
		return fmt.Errorf("Parse error: %v", err)
	}

	slices.SortFunc(hands, func(a, b HandWithBid) int {
		return compareHands(a.hand, b.hand, strengthOfCardsWithoutJokers)
	})

	output := 0
	for idx, hand := range hands {
		rank := idx + 1
		output += rank * hand.bid
	}

	fmt.Println(output)
	return nil
}

func Star2(input string) error {
	if input == "" {
		return fmt.Errorf("input file expected")
	}

	hands, err := parseInput(input, true)
	if err != nil {
		return fmt.Errorf("Parse error: %v", err)
	}

	slices.SortFunc(hands, func(a, b HandWithBid) int {
		return compareHands(a.hand, b.hand, strengthOfCardsWithJokers)
	})

	output := 0
	for idx, hand := range hands {
		rank := idx + 1
		output += rank * hand.bid
	}

	fmt.Println(output)
	return nil
}

// Return -1 when a < b, 1 when a > b, 0 when a == b
func compareHands(a Hand, b Hand, strengthFunc func(h Hand) int) int {
	aStrength := strengthFunc(a)
	bStrength := strengthFunc(b)

	if aStrength < bStrength {
		return -1
	} else if aStrength > bStrength {
		return 1
	} else {
		return slices.Compare(a, b)
	}
}

func strengthOfCardsWithoutJokers(h Hand) int {
	countsByCard := map[Card]int{}
	for _, card := range h {
		countsByCard[card] += 1
	}

	counts := []int{}
	for _, count := range countsByCard {
		counts = append(counts, count)
	}

	slices.SortFunc(counts, func(a, b int) int {
		return cmp.Compare(a, b) * -1
	})

	strength := counts[0] * 2
	if (counts[0] == 3 || counts[0] == 2) && counts[1] == 2 {
		strength += 1
	}

	return strength
}

func strengthOfCardsWithJokers(h Hand) int {
	numJokers := 0
	countsByCard := map[Card]int{}
	for _, card := range h {
		// don't count jokers at first
		if card != JOKER {
			countsByCard[card] += 1
		} else {
			numJokers += 1
		}
	}

	counts := []int{}
	for _, count := range countsByCard {
		counts = append(counts, count)
	}

	slices.SortFunc(counts, func(a, b int) int {
		return cmp.Compare(a, b) * -1
	})

	// Now, improve counts by making jokers whatever gives best
	// hand, ie whatever card already is seen the most.
	if len(counts) > 0 {
		counts[0] += numJokers
	} else {
		counts = append(counts, numJokers)
	}

	strength := counts[0] * 2
	if (counts[0] == 3 || counts[0] == 2) && counts[1] == 2 {
		strength += 1
	}

	return strength
}

type HandWithBid struct {
	hand Hand
	bid  int
}

type Hand []Card
type Card int

const (
	JOKER Card = 1
	TWO        = 2
	THREE      = 3
	FOUR       = 4
	FIVE       = 5
	SIX        = 6
	SEVEN      = 7
	EIGHT      = 8
	NINE       = 9
	TEN        = 10
	JACK       = 11
	QUEEN      = 12
	KING       = 13
	ACE        = 14
)

func parseCard(card rune, useJoker bool) (Card, error) {
	switch card {
	case '2':
		return TWO, nil
	case '3':
		return THREE, nil
	case '4':
		return FOUR, nil
	case '5':
		return FIVE, nil
	case '6':
		return SIX, nil
	case '7':
		return SEVEN, nil
	case '8':
		return EIGHT, nil
	case '9':
		return NINE, nil
	case 'T':
		return TEN, nil
	case 'J':
		if useJoker {
			return JOKER, nil
		} else {
			return JACK, nil
		}
	case 'Q':
		return QUEEN, nil
	case 'K':
		return KING, nil
	case 'A':
		return ACE, nil
	default:
		return Card(0), fmt.Errorf("Cannot parse %v into a card", card)
	}
}

func parseHand(handStr string, useJoker bool) (Hand, error) {
	hand := Hand{}
	for _, c := range handStr {
		card, err := parseCard(c, useJoker)
		if err != nil {
			return hand, err
		}
		hand = append(hand, card)
	}
	return hand, nil
}

func parseInput(input string, useJoker bool) ([]HandWithBid, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	output := []HandWithBid{}
	for _, line := range lines {
		bits := strings.Split(line, " ")
		handString := bits[0]
		bidString := bits[1]

		hand, err := parseHand(handString, useJoker)
		if err != nil {
			return output, fmt.Errorf("Error parsing hand (%v) in %v: %v", line, handString, err)
		}

		bid, err := strconv.Atoi(bidString)
		if err != nil {
			return output, fmt.Errorf("Error parsing bid (%v) in %v: %v", bid, handString, err)
		}

		output = append(output, HandWithBid{hand, bid})
	}
	return output, nil
}
