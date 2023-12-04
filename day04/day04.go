package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input
var input string

func main() {
	scanner := bufio.NewScanner(strings.NewReader(input))
	re := regexp.MustCompile(`\d+`)

	var pointSum int
	copies := []int{1}
	var cardSum int

	for scanner.Scan() {
		line := scanner.Text()
		_, line, _ = strings.Cut(line, ": ")
		winningLine, haveLine, _ := strings.Cut(line, " | ")
		winningMatches := re.FindAllString(winningLine, -1)
		haveMatches := re.FindAllString(haveLine, -1)
		var winning []int
		haves := make(map[int]bool)
		for _, match := range winningMatches {
			num, err := strconv.Atoi(match)
			if err != nil {
				panic(err)
			}
			winning = append(winning, num)
		}
		for _, match := range haveMatches {
			num, err := strconv.Atoi(match)
			if err != nil {
				panic(err)
			}
			haves[num] = true
		}

		multiplier := copies[0]
		cardSum += multiplier
		copies = copies[1:]

		var points int
		var wins int
		for _, num := range winning {
			_, ok := haves[num]
			if ok {
				points *= 2
				if points == 0 {
					points = 1
				}
				wins++
			}
		}
		for i := 0; i < wins; i++ {
			if len(copies) < i+1 {
				copies = append(copies, multiplier+1)
			} else {
				copies[i] += multiplier
			}
		}

		if len(copies) == 0 {
			copies = append(copies, 1)
		}

		fmt.Printf("%7d | %2d | %8d | %8d \n", multiplier, wins, cardSum, copies)

		pointSum += points
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("(Part 1) Sum of Point of all Cards: %d \n", pointSum)
	fmt.Printf("(Part 2) Sum of all Cards: %d \n", cardSum)
}
