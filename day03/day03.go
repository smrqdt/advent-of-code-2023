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

type PartNo struct {
	Row, From, To int
	No            int
}
type Symbol struct {
	Row, Col int
	Char     byte
}

func (sym *Symbol) String() string {
	return fmt.Sprintf("Sym{%c, [%d, %d]}", sym.Char, sym.Row, sym.Col)
}

func main() {
	scanner := bufio.NewScanner(strings.NewReader(input))

	re := regexp.MustCompile(`\d+`)
	reSymbol := regexp.MustCompile(`[/#%&*+=@$-]`)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var partNumberSum int
	var gearRatioSum int
	vertices := make(map[*Symbol][]*PartNo)
	symbols := make([][]*Symbol, len(lines))
	for i := range symbols {
		symbols[i] = make([]*Symbol, len(lines[0]))
	}

	addPartNoToSymbol := func(partNo *PartNo, row, col int, char byte) (symbol *Symbol) {
		symbol = symbols[row][col]
		if symbol == nil {
			symbol = &Symbol{Row: row, Col: col, Char: char}
			symbols[row][col] = symbol
		}
		if symbols[row][col].Char != char {
			fmt.Printf("Existing symbol %c does not match %c! ", symbols[row][col].Char, char)
		}
		vertices[symbol] = append(vertices[symbol], partNo)
		return
	}

	for lineNo, line := range lines {
		fmt.Printf("%s | ", line)
		matches := re.FindAllStringIndex(line, -1)
		for _, match := range matches {
			partNo := PartNo{
				Row:  lineNo,
				From: match[0],
				To:   match[1],
			}
			no, err := strconv.Atoi(line[partNo.From:partNo.To])
			if err != nil {
				panic(err)
			}
			partNo.No = no

			firstLine := lineNo == 0
			firstColumn := match[0] == 0
			lastColumn := match[1] == len(line)
			lastLine := lineNo == len(lines)-1

			from := match[0]
			if !firstColumn {
				from--
			}
			to := match[1]
			if !lastColumn {
				to++
			}

			hasSymbol := false

			// current line
			matches := reSymbol.FindAllStringIndex(line[from:to], -1)
			for _, match := range matches {
				symbol := addPartNoToSymbol(&partNo, lineNo, match[0]+from, line[match[0]+from])
				fmt.Printf("%d→%c ", partNo.No, symbol.Char)
				hasSymbol = true
			}

			if !firstLine {
				// line before
				matches := reSymbol.FindAllStringIndex(lines[lineNo-1][from:to], -1)
				for _, match := range matches {
					addPartNoToSymbol(&partNo, lineNo-1, match[0]+from, lines[lineNo-1][match[0]+from])
					hasSymbol = true
				}
			}

			if !lastLine {
				// line after
				matches := reSymbol.FindAllStringIndex(lines[lineNo+1][from:to], -1)
				for _, match := range matches {
					addPartNoToSymbol(&partNo, lineNo+1, match[0]+from, lines[lineNo+1][match[0]+from])
					hasSymbol = true
				}
			}

			if hasSymbol {
				partNumberSum += partNo.No
			}
		}
		fmt.Println()
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for sym, partNumbers := range vertices {
		if sym.Char == '*' && len(partNumbers) == 2 {
			gearRatioSum += partNumbers[0].No * partNumbers[1].No
		}
	}

	// for _, row := range symbols {
	// 	for _, sym := range row {
	// 		if sym != nil {
	// 			fmt.Println(sym)
	// 		}
	// 	}
	// }

	fmt.Printf("(Part 1) Sum of all part numbers: %d\n", partNumberSum)
	fmt.Printf("(Part 2) Sum of all gear ratios: %d\n", gearRatioSum)
}
