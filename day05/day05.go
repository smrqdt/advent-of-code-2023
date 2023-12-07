package main

import (
	"bufio"
	"cmp"
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
	part2()
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

type Mapping struct {
	Dest, Source, Length int
}

func part2() {
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

	var mappings [][]Mapping

	for scanner.Scan() {
		// discard "x-to-y map:" line
		scanner.Text()
		var mapping []Mapping
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
			mappingRow := Mapping{Dest: row[0], Source: row[1], Length: row[2]}
			mapping = append(mapping, mappingRow)
			if line == "" {
				break
			}
		}
		mappings = append(mappings, mapping)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for i := range mappings {
		slices.SortFunc[[]Mapping](
			mappings[i],
			func(a, b Mapping) int {
				return cmp.Compare[int](a.Dest, b.Dest)
			})
		for _, line := range mappings[i] {
			fmt.Println(line)
		}
		fmt.Println()
	}

	solution := func() int {
		for _, row := range mappings[len(mappings)-1] {
			for location := row.Dest; location < row.Dest+row.Length; location++ {
				source := row.Source
				for i := len(mappings) - 2; i >= 0; i-- {
					for _, row := range mappings[i] {
						if source >= row.Dest && source < row.Dest+row.Length {
							source = row.Source + (source - row.Dest)
							break
						}
					}
				}
				for i := 0; i < len(seeds)-1; i += 2 {
					if source >= seeds[i] && source < seeds[i]+seeds[i+1] {
						return location
					}
				}
			}
		}
		return -1
	}()

	fmt.Printf("(Part 2) Minimum Location Number: %d \n", solution)
}
