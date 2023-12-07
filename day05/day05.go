package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input
var input string

var example = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
`

func main() {
	part1()
}

func part1() {
	scanner := bufio.NewScanner(strings.NewReader(input))

	// first line
	scanner.Scan()
	line := scanner.Text()
	fields := strings.Fields(line)
	seeds := make([]int, 0, len(fields)-1)
	for _, field := range fields[1:] {
		seed, err := strconv.Atoi(field)
		if err != nil {
			panic(err)
		}
		seeds = append(seeds, seed)
	}
	// empty line
	scanner.Scan()

	var mappings [][][]int

	for scanner.Scan() {
		// discard "x-to-y map:" line
		scanner.Text()
		var mapping [][]int
		for scanner.Scan() {
			line := scanner.Text()
			fields := strings.Fields(line)
			row := make([]int, 3)
			for i, field := range fields {
				num, err := strconv.Atoi(field)
				if err != nil {
					panic(err)
				}
				row[i] = num
			}
			mapping = append(mapping, row)
			if line == "" {
				break
			}
		}
		mappings = append(mappings, mapping)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	var destinations []int
	for _, seed := range seeds {
		source := seed
		for _, mapping := range mappings {
			fmt.Printf("%d -> ", source)
			for _, row := range mapping {
				if source >= row[1] && source < row[1]+row[2] {
					source = row[0] + (source - row[1])
					break
				}
			}
		}
		fmt.Printf("%d. \n", source)
		destinations = append(destinations, source)
	}

	lowest := slices.Min[[]int](destinations)

	fmt.Printf("(Part 1) Minimum Location Number: %d \n", lowest)
}
