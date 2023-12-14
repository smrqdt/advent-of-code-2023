package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input
var input string

var example = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

func main() {
	_ = example
	scanner := bufio.NewScanner(strings.NewReader(input))
	var report [][]int

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		history := make([]int, 0, len(fields))
		for _, field := range fields {
			value, err := strconv.Atoi(field)
			if err != nil {
				panic(err)
			}
			history = append(history, value)
		}
		report = append(report, history)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	var sumFirst, sumLast int
	for _, values := range report {
		fmt.Println(values)
		predFirst, predLast := predict(values)
		sumFirst += predFirst
		sumLast += predLast
		fmt.Printf("Prediction: %d/%d Sum: %d/%d\n", predFirst, predLast, sumFirst, sumLast)
	}

	fmt.Printf("(Part 1) Sum of the extrapolated values: %d\n", sumLast)
	fmt.Printf("(Part 2) Sum of the extrapolated values: %d\n", sumFirst)
}

func predict(input []int) (first, last int) {
	zero := true
	for _, val := range input {
		if val != 0 {
			zero = false
			break
		}
	}
	if zero {
		return 0, 0
	}
	diff := make([]int, len(input)-1)
	for i := 0; i < len(input)-1; i++ {
		diff[i] = input[i+1] - input[i]
		fmt.Printf("%d ", diff[i])
	}
	fmt.Println()
	belowFirst, belowLast := predict(diff)
	resultLast := input[len(input)-1] + belowLast
	resultFirst := input[0] - belowFirst
	return resultFirst, resultLast
}
