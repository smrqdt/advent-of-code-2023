package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input
var input string

func main() {
	parse()
	part1()
}

type Galaxy struct {
	id   int
	x, y int
}

func (g *Galaxy) String() string {
	return fmt.Sprintf("Galaxy(%d [%d:%d])", g.id, g.x, g.y)
}

var Galaxies []*Galaxy
var Universe [][]*Galaxy
var colHasGalaxy []bool
var rowHasGalaxy []bool

func parse() {
	scanner := bufio.NewScanner(strings.NewReader(input))

	colHasGalaxy = make([]bool, strings.IndexRune(input, '\n'))
	var galID int

	for scanner.Scan() {
		line := scanner.Text()
		Universe = append(Universe, make([]*Galaxy, len(line)))
		y := len(Universe) - 1

		var galaxyFound bool
		for x, pixel := range line {
			if pixel == '#' {
				galaxyFound = true
				colHasGalaxy[x] = true
				newGalaxy := Galaxy{galID, x, y}
				galID++
				Universe[y][x] = &newGalaxy
				Galaxies = append(Galaxies, &newGalaxy)
			}
		}
		rowHasGalaxy = append(rowHasGalaxy, galaxyFound)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func part1() {
	var resultSum int
	for i, startGalaxy := range Galaxies {
		for _, targetGalaxy := range Galaxies[i+1:] {
			path := findPath(startGalaxy, targetGalaxy)
			// fmt.Println(startGalaxy, targetGalaxy, path)
			resultSum += path
		}
	}
	fmt.Printf("(Part 1) Sum of lengths: %d \n", resultSum)
}

func findPath(start, target *Galaxy) int {
	var patheLen int
	patheLen += abs(target.x - start.x)
	patheLen += countEmpty(start.x, target.x, colHasGalaxy)
	patheLen += abs(target.y - start.y)
	patheLen += countEmpty(start.y, target.y, rowHasGalaxy)
	return patheLen
}

func countEmpty(start, end int, hasGalaxy []bool) int {
	if start > end {
		start, end = end, start
	}
	var count int
	for _, has := range hasGalaxy[start:end] {
		if !has {
			count++
		}
	}
	return count
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
