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

var example = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
`

func main() {
	leftRight, nodes := parseInput()
	part1(leftRight, nodes)
	part2(leftRight, nodes)
}

type Node struct {
	Name                string
	Left, Right         *Node
	LeftName, RightName string
}

func (n Node) String() string {
	return fmt.Sprintf("Node{%s (%s|%s)}", n.Name, n.LeftName, n.RightName)
}

func parseInput() (leftRight []byte, nodes map[string]*Node) {
	scanner := bufio.NewScanner(strings.NewReader(input))

	scanner.Scan()
	leftRight = []byte(scanner.Text())
	scanner.Scan()

	re := regexp.MustCompile(`\w{3}`)

	nodes = make(map[string]*Node)
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
	return
}

func part1(leftRight []byte, nodes map[string]*Node) {
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

func part2(leftRight []byte, nodes map[string]*Node) {
	var activeNodes []*Node
	for name, node := range nodes {
		if name[2] == 'A' {
			activeNodes = append(activeNodes, node)
		}
	}

	steps := make([]int, len(activeNodes))
	for i := range activeNodes {
		for steps[i] = 0; activeNodes[i].Name[2] != 'Z'; steps[i]++ {
			switch leftRight[steps[i]%len(leftRight)] {
			case 'L':
				activeNodes[i] = activeNodes[i].Left
			case 'R':
				activeNodes[i] = activeNodes[i].Right
			}
		}
	}
	gcd := 0
	for _, step := range steps {
		gcd = GreatestCommonDivisor(gcd, step)
	}
	result := 1
	for _, step := range steps {
		result *= step / gcd
	}
	result *= gcd

	fmt.Printf("(Part 2) Steps from '*A' to '*Z': %d \n", result)

}

func GreatestCommonDivisor(a, b int) int {
	switch {
	case a == 0:
		return b
	case b == 0:
		return a
	case a < b:
		return GreatestCommonDivisor(b, a)
	default:
		return GreatestCommonDivisor(b, a%b)
	}
}
