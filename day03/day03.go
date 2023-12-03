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
	part1()
}

func part1() {
	scanner := bufio.NewScanner(strings.NewReader(input))

	re := regexp.MustCompile(`\d+`)
	reSymbol := regexp.MustCompile(`[/#%&*+=@$-]`)

	var partNumberSum int

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for lineNo, line := range lines {
		fmt.Printf("%s: ", line)
		matches := re.FindAllStringIndex(line, -1)
		for _, match := range matches {
			var check []byte
			firstLine := lineNo == 0
			firstColumn := match[0] == 0
			lastColumn := match[1] == len(line)
			lastLine := lineNo == len(lines)-1

			if !firstLine {
				// line before
				from := match[0]
				if !firstColumn {
					from--
				}
				to := match[1]
				if !lastColumn {
					to++
				}

				check = append(check, lines[lineNo-1][from:to]...)
			}
			if !firstColumn {
				// column before
				check = append(check, line[match[0]-1])
			}
			if !lastColumn {
				// column after
				check = append(check, line[match[1]])
			}
			if !lastLine {
				// line after

				from := match[0]
				if !firstColumn {
					from--
				}
				to := match[1]
				if !lastColumn {
					to++
				}

				check = append(check, lines[lineNo+1][from:to]...)
			}
			symbols := reSymbol.Find(check)
			if symbols != nil {

				partNum, err := strconv.Atoi(line[match[0]:match[1]])
				if err != nil {
					panic(err)
				}
				partNumberSum += partNum
				fmt.Printf("%d ", partNum)
			}
		}
		fmt.Println()
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("(Part 1) Sum of all part numbers: %d\n", partNumberSum)
}
