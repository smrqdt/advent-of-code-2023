package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed input
var input string

func main() {
	part1()
}

type Direction int

const (
	UP Direction = iota
	RIGHT
	DOWN
	LEFT
	STAY
	INVALID
)

var DIRECTIONS = [...]string{"UP", "RIGHT", "DOWN", "LEFT", "STAY", "INVALID"}

func (d Direction) String() string {
	return DIRECTIONS[d]
}

type GridField byte

const (
	VERTICAL   GridField = '|' // │
	HORIZONTAL GridField = '-' // ─
	BEND_NE    GridField = 'L' // └
	BEND_NW    GridField = 'J' // ┘
	BEND_SW    GridField = '7' // ┐
	BEND_SE    GridField = 'F' // ┌
	GROUND     GridField = '.'
	START      GridField = 'S'
)

type Coordinates struct {
	Row, Col int
}

func (c Coordinates) check(max Coordinates) bool {
	return c.Row >= 0 && c.Row < max.Row && c.Col >= 0 && c.Col < max.Col
}

func (c Coordinates) String() string {
	return fmt.Sprintf("(%d|%d)", c.Row, c.Col)
}

func (c Coordinates) dir(dir Direction) (nextCoord Coordinates) {
	switch dir {
	case UP:
		return c.up()
	case DOWN:
		return c.down()
	case LEFT:
		return c.left()
	case RIGHT:
		return c.right()
	}
	panic("invalid direction")
}

func (c Coordinates) up() (nextCoord Coordinates) {
	c.Row--
	return c

}

func (c Coordinates) down() (nextCoord Coordinates) {
	c.Row++
	return c
}

func (c Coordinates) left() (nextCoord Coordinates) {
	c.Col--
	return c
}

func (c Coordinates) right() (nextCoord Coordinates) {
	c.Col++
	return c
}

func (f GridField) next(c Coordinates, in Direction) (nextCoord Coordinates, out Direction) {
	switch f {
	case VERTICAL:
		switch in {
		case DOWN:
			return c.down(), DOWN
		case UP:
			return c.up(), UP
		}
	case HORIZONTAL:
		switch in {
		case LEFT:
			return c.left(), LEFT
		case RIGHT:
			return c.right(), RIGHT
		}
	case BEND_NE:
		switch in {
		case DOWN:
			return c.right(), RIGHT
		case LEFT:
			return c.up(), UP
		}
	case BEND_NW:
		switch in {
		case DOWN:
			return c.left(), LEFT
		case RIGHT:
			return c.up(), UP
		}
	case BEND_SE:
		switch in {
		case UP:
			return c.right(), RIGHT
		case LEFT:
			return c.down(), DOWN
		}
	case BEND_SW:
		switch in {
		case UP:
			return c.left(), LEFT
		case RIGHT:
			return c.down(), DOWN
		}
	case START:
		return c, STAY
	}
	return c, INVALID
}

func part1() {
	var grid [][]GridField

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []GridField(line))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Find Start
	var startCoord Coordinates
	for row, line := range grid {
		col := slices.Index(line, START)
		if col != -1 {
			startCoord = Coordinates{Row: row, Col: col}
			break
		}
	}
	fmt.Printf("%v (Start): %q\n", startCoord, grid[startCoord.Row][startCoord.Col])

	// Find First Pipe
	max := Coordinates{Row: len(grid), Col: len(grid[0])}
	var dir Direction
	cur := startCoord

	for _, dir = range []Direction{UP, RIGHT, DOWN, LEFT} {
		cur = startCoord.dir(dir)
		if cur.check(max) {
			_, newDir := grid[cur.Row][cur.Col].next(cur, dir)
			if newDir != INVALID {
				fmt.Printf("%v (%v): %q (First) \n", cur, dir, grid[cur.Row][cur.Col])
				break
			}
		}
	}

	steps := 0
	for dir != STAY {
		cur, dir = grid[cur.Row][cur.Col].next(cur, dir)
		steps++
		fmt.Printf("%v (%v): %q\n", cur, dir, grid[cur.Row][cur.Col])
		if dir == INVALID {
			panic("invalid direction")
		}
	}
	fmt.Printf("%d Steps \n", steps)

	fmt.Printf("(Part 1) Farthest Point: %d \n", steps/2)
}
