package part1

import (
	"bufio"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Hand struct {
	Hand []Card
	Bid  int
}

type HandType int

const (
	HIGH_CARD HandType = iota // five distinct cards
	ONE_PAIR
	TWO_PAIR
	THREE_OF_A_KIND
	FULL_HOUSE // Three of a Kind + Pair
	FOUR_OF_A_KIND
	FIVE_OF_A_KIND
)

type Card byte

var CARD_ORDER = []Card{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}

func (c Card) Value() int {
	return slices.Index(CARD_ORDER, c)
}

func (c Card) Cmp(other Card) int {
	return c.Value() - other.Value()
}

func (h Hand) getPoints() int {
	hand := make([]Card, len(h.Hand))
	copy(hand, h.Hand)
	slices.Sort(hand)
	compact := make([]Card, len(h.Hand))
	copy(compact, hand)
	compact = slices.Compact(compact)

	switch len(compact) {
	case 1:
		return int(FIVE_OF_A_KIND)
	case 2:
		if hand[1] == hand[3] {
			return int(FOUR_OF_A_KIND)
		}
		return int(FULL_HOUSE)
	case 3:
		if hand[2] == hand[0] || hand[2] == hand[4] || hand[1] == hand[3] {
			return int(THREE_OF_A_KIND)
		}
		return int(TWO_PAIR)
	case 4:
		return int(ONE_PAIR)
	}
	return int(HIGH_CARD)
}

func (h Hand) Cmp(other Hand) int {
	cmp := h.getPoints() - other.getPoints()
	for i := 0; cmp == 0 && i < 5; i++ {
		cmp = h.Hand[i].Cmp(other.Hand[i])
	}
	return cmp
}

func Part1(input string) {
	var hands []Hand

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		bid, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(err)
		}
		hand := Hand{Hand: []Card(fields[0]), Bid: bid}
		hands = append(hands, hand)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	slices.SortFunc(hands, func(a, b Hand) int { return a.Cmp(b) })

	var winnings int
	for i, hand := range hands {
		fmt.Printf("% 3d) %s: %d (%d)\n", i, hand.Hand, hand.getPoints(), hand.Bid)
		winnings += (i + 1) * hand.Bid
	}

	fmt.Printf("(Part 1) Total Winnings: %d\n", winnings)
}
