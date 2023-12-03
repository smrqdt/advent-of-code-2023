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

type Cubes struct {
	Red, Green, Blue int
}

func (c Cubes) Power() int {
	return c.Red * c.Green * c.Blue
}

func main() {
	scanner := bufio.NewScanner(strings.NewReader(input))
	reGame := regexp.MustCompile(`Game (\d+): (.+)`)
	reRound := regexp.MustCompile(`(\d+) (red|green|blue)`)
	// Part 1 Result
	validGameSum := 0
	// Part 2 Resul
	powerSum := 0
	for gameID := 1; scanner.Scan(); gameID++ {
		line := scanner.Text()
		matches := reGame.FindStringSubmatch(line)
		rounds := strings.Split(matches[2], ";")

		// Part 1
		maxCubes := Cubes{Red: 12, Green: 13, Blue: 14}
		gameValid := true
		// Part 2
		minimumCubes := Cubes{}

		for _, round := range rounds {
			matches := reRound.FindAllStringSubmatch(round, 3)
			for _, match := range matches {
				num, err := strconv.Atoi(match[1])
				if err != nil {
					panic(err)
				}
				switch match[2] {
				case "red":
					gameValid = gameValid && num <= maxCubes.Red
					minimumCubes.Red = max(minimumCubes.Red, num)
				case "green":
					gameValid = gameValid && num <= maxCubes.Green
					minimumCubes.Green = max(minimumCubes.Green, num)
				case "blue":
					gameValid = gameValid && num <= maxCubes.Blue
					minimumCubes.Blue = max(minimumCubes.Blue, num)
				}
			}
		}
		if gameValid {
			validGameSum += gameID
		}
		powerSum += minimumCubes.Power()
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("(Part 1) Sum of IDs of valid Games: %d \n", validGameSum)
	fmt.Printf("(Part 2) Sum of Power of minimum Cubes: %d \n", powerSum)
}
