package main

import (
	_ "embed"

	"github.com/smrqdt/adventofcode-2023/day07/part1"
	"github.com/smrqdt/adventofcode-2023/day07/part2"
)

//go:embed input
var input string

var example = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

func main() {
	part1.Part1(input)
	part2.Part2(input)
}
