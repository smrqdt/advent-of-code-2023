package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input
var input string

func main() {
	part1()
	part2()
}

func part1() {

	timeLine, distanceLine, _ := strings.Cut(input, "\n")
	times := strings.Fields(timeLine)[1:]
	distances := strings.Fields(distanceLine)[1:]
	winProduct := 1

	for i := range times {
		time, err := strconv.Atoi(times[i])
		if err != nil {
			panic(err)
		}
		distance, err := strconv.Atoi(distances[i])
		if err != nil {
			panic(err)
		}
		var wins int
		for holdTime := 1; holdTime < time; holdTime++ {
			result := holdTime * (time - holdTime)
			if result > distance {
				wins++
			}
		}
		winProduct *= wins

	}

	fmt.Printf("(Part 1) Result: %d \n", winProduct)
}

func part2() {

	timeLine, distanceLine, _ := strings.Cut(input, "\n")
	time, err := strconv.Atoi(strings.Join(strings.Fields(timeLine)[1:], ""))
	if err != nil {
		panic(err)
	}
	distance, err := strconv.Atoi(strings.Join(strings.Fields(distanceLine)[1:], ""))
	if err != nil {
		panic(err)
	}

	var wins int
	for holdTime := 1; holdTime < time; holdTime++ {
		result := holdTime * (time - holdTime)
		if result > distance {
			wins++
		}
	}

	fmt.Printf("(Part 2) Result: %d \n", wins)
}
