package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"regexp"
	"strings"
)

//go:embed input
var input string

func main() {
	part1()
}

type Node struct {
	Name                string
	Left, Right         *Node
	LeftName, RightName string
}

func part1() {
	scanner := bufio.NewScanner(strings.NewReader(input))

	scanner.Scan()
	leftRight := []byte(scanner.Text())
	scanner.Scan()

	re := regexp.MustCompile(`[A-Z]{3}`)

	nodes := make(map[string]*Node)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllString(line, -1)
		nodes[matches[0]] = &Node{Name: matches[0], LeftName: matches[1], RightName: matches[2]}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for _, node := range nodes {
		if leftNode, ok := nodes[node.LeftName]; ok {
			node.Left = leftNode
		}
		if rightNode, ok := nodes[node.RightName]; ok {
			node.Right = rightNode
		}
	}

	node := nodes["AAA"]
	var steps int

	for steps = 0; node != nodes["ZZZ"]; steps++ {
		switch leftRight[steps%len(leftRight)] {
		case 'L':
			node = node.Left
		case 'R':
			node = node.Right
		}
	}

	fmt.Printf("(Part 1) Steps from 'AAA' to 'ZZZ': %d \n", steps)
}
