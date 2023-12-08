package main

import (
	"bufio"
	"cmp"
	_ "embed"
	"fmt"
	"slices"
	"sort"
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

func part2() {
	seeds, maps := parseInput(example)
	fmt.Printf("Seeds: %v\n", seeds)
	fmt.Println("Maps: ")
	for i, mappings := range maps {
		fmt.Printf("  Map %d: \n", i)
		for _, mapping := range mappings {
			fmt.Printf("    %v\n", mapping)
		}
	}
	intervals := buildIntervals(seeds, maps)
	fmt.Printf("Intervals: %v\n", intervals)
	lowest := findLowest(intervals)
	fmt.Printf("(Part 2) Minimum Location Number: %d \n", lowest)
}

type Interval struct {
	Begin, End int
	Steps      []int
}

func (i Interval) String() string {
	return fmt.Sprintf("Interval{%d:%d:%v → %d}", i.Begin, i.End, i.Steps, i.Begin+i.SumSteps())
}

func (i Interval) Contains(num int) bool {
	return num >= i.Begin && num < i.End
}

func (i Interval) CmpToInt(b int) int {
	if i.Begin+i.SumSteps() > b {
		return 1
	}
	if i.End+i.SumSteps() > b {
		return -1
	}
	return 0
}

func (i Interval) SumSteps() (sum int) {
	for _, step := range i.Steps {
		sum += step
	}
	return
}

func cmpIntervalToInt(a Interval, b int) int {
	return a.CmpToInt(b)
}

func parseInput(input string) (seeds []Interval, maps [][]Interval) {
	scanner := bufio.NewScanner(strings.NewReader(input))

	// first line
	scanner.Scan()
	line := scanner.Text()
	fields := strings.Fields(line)
	for i := 1; i < len(fields); i += 2 {
		seedStart, err := strconv.Atoi(fields[i])
		if err != nil {
			panic(err)
		}
		seedRange, err := strconv.Atoi(fields[i+1])
		if err != nil {
			panic(err)
		}
		seed := Interval{Begin: seedStart, End: seedStart + seedRange}
		seeds = append(seeds, seed)
	}
	// empty line
	scanner.Scan()

	for scanner.Scan() {
		// discard "x-to-y map:" line
		scanner.Text()
		var mappings []Interval
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				break
			}
			fields := strings.Fields(line)
			row := make([]int, 3)
			for i, field := range fields {
				num, err := strconv.Atoi(field)
				if err != nil {
					panic(err)
				}
				row[i] = num
			}
			mappingRow := Interval{Begin: row[1], Steps: []int{row[0] - row[1]}, End: row[1] + row[2]}
			mappings = append(mappings, mappingRow)
		}
		maps = append(maps, mappings)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for i := range maps {
		slices.SortFunc[[]Interval](maps[i], func(a, b Interval) int { return cmp.Compare[int](a.Begin, b.Begin) })
	}

	return seeds, maps
}

func buildIntervals(seeds []Interval, maps [][]Interval) (intervals []*Interval) {
	for i := range seeds {
		intervals = append(intervals, &seeds[i])
	}

	for _, mappings := range maps {
		for _, interval := range intervals {
			mappingIdx, found := slices.BinarySearchFunc(mappings, interval.Begin, cmpIntervalToInt)
			fmt.Printf("Found (%v) %v in %v at %d\n", found, interval, mappings, mappingIdx)
			if found {
				// interval beginn found in mappings
				interval.Steps = append(interval.Steps, mappings[mappingIdx].Steps...)
			} else {
				// interval begin not included in mappings
				interval.Steps = append(interval.Steps, 0)
			}
			for ; mappingIdx > len(mappings); mappingIdx++ {
				if mappings[mappingIdx].Begin >= interval.End {
					break
				}
				// split interval
				newInterval := &Interval{Begin: mappings[mappingIdx].Begin, End: interval.End, Steps: interval.Steps}
				newInterval.Steps = append(newInterval.Steps[:len(newInterval.Steps)-1], mappings[mappingIdx].Steps...)
				interval.End = mappings[mappingIdx].Begin
				intervals = append(intervals, newInterval)
				interval = newInterval
			}
		}
		slices.SortFunc[[]*Interval](intervals, func(a, b *Interval) int { return cmp.Compare[int](a.Begin, b.Begin) })
	}
	return intervals

}

func findLowest(intervals []*Interval) (lowest int) {
	var locations []int
	for _, interval := range intervals {
		fmt.Println(interval)
		for i := interval.Begin; i < interval.End; i++ {
			fmt.Printf("%d: %v %d\n", i, interval.Steps, interval.SumSteps())
			locations = append(locations, i+interval.SumSteps())
		}
	}
	sort.Ints(locations)
	fmt.Println(locations)
	return locations[0]
}
