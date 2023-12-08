package part2

import (
	"bufio"
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Hand struct {
	Hand   []Card
	Bid    int
	Points int
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

var CARD_ORDER = []Card{'J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A'}

func (c Card) Value() int {
	return slices.Index(CARD_ORDER, c)
}

func (c Card) Cmp(other Card) int {
	return c.Value() - other.Value()
}

func (h *Hand) calcPoints() {
	cardMap := make(map[Card]int)

	for _, card := range h.Hand {
		cardMap[card]++
	}

	type CardNum struct {
		Card Card
		Num  int
	}
	var cards []CardNum
	for card, num := range cardMap {
		cards = append(cards, CardNum{card, num})
	}
	slices.SortStableFunc(cards, func(a, b CardNum) int { return cmp.Compare(b.Num, a.Num) })

	var three, pair bool

	maxCards := 5 - cardMap['J']

	for _, c := range cards {
		if c.Card == 'J' {
			continue
		}
		switch c.Num {
		case maxCards:
			h.Points = int(FIVE_OF_A_KIND)
			return
		case maxCards - 1:
			h.Points = int(FOUR_OF_A_KIND)
			return
		case maxCards - 2:
			if pair || three {
				h.Points = int(FULL_HOUSE)
				return
			}
			three = true
			maxCards = 5
		case maxCards - 3:
			if three {
				h.Points = int(FULL_HOUSE)
				return
			}
			if pair {
				h.Points = int(TWO_PAIR)
				return
			}
			pair = true
			maxCards = 5
		}
	}
	if three {
		h.Points = int(THREE_OF_A_KIND)
		return
	}
	if pair {
		h.Points = int(ONE_PAIR)
		return
	}
	if cardMap['J'] == 5 {
		h.Points = int(FIVE_OF_A_KIND)
		return
	}
	h.Points = int(HIGH_CARD)
}

func (h Hand) Cmp(other Hand) int {
	cmp := h.Points - other.Points
	for i := 0; cmp == 0 && i < 5; i++ {
		cmp = h.Hand[i].Cmp(other.Hand[i])
	}
	return cmp
}

func Part2(input string) {
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
		hand.calcPoints()
		hands = append(hands, hand)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	slices.SortStableFunc(hands, func(a, b Hand) int { return a.Cmp(b) })

	var winnings int
	for i, hand := range hands {
		fmt.Printf("%s: %d (%d)\n", hand.Hand, hand.Points, hand.Bid)
		winnings += (i + 1) * hand.Bid
	}

	fmt.Printf("(Part 2) Total Winnings: %d\n", winnings)
}
