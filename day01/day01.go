package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

const InputFile = "day01/input"

func main() {
	part1()
	part2()
}

func part1() {
	f, err := os.Open(InputFile)
	if err != nil {
		panic(fmt.Sprintf("Could not open file %v \n", err))
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	calibrationSum := 0

	re := regexp.MustCompile(`\d`)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllString(line, -1)
		calib, err := strconv.ParseInt(matches[0]+matches[len(matches)-1], 10, 32)
		if err != nil {
			panic(err)
		}
		calibrationSum += int(calib)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("(Part 1): Calibration Sum is %d \n", calibrationSum)
}

func part2() {
	f, err := os.Open(InputFile)
	if err != nil {
		panic(fmt.Sprintf("Could not open file %v \n", err))
	}
	defer f.Close()

	// f := strings.NewReader("two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen")

	scanner := bufio.NewScanner(f)

	calibrationSum := 0

	// Two regexes to solve the problem of overlapping matches (fiveeight must be parsed as 58)
	// there might be a better solution, but it seems go regexp doesn’t support lookahead (?)
	re := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine|\d)`)
	re_last := regexp.MustCompile(`.*(one|two|three|four|five|six|seven|eight|nine|\d)`)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("%s: ", line)
		for i, word := range []string{} {
			line = strings.ReplaceAll(line, word, fmt.Sprintf("%d", i+1))
		}
		first := re.FindString(line)
		matches := re_last.FindStringSubmatch(line)
		last := matches[1]
		wordToInt := func(word string) int {
			num64, err := strconv.ParseInt(word, 10, 32)
			num := int(num64)
			if err != nil {
				words := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
				num = slices.Index(words, word) + 1
				if num == 0 {
					panic("Neither word nor digit")
				}
			}
			fmt.Printf("%s→%d, ", word, num)
			return int(num)
		}
		calib := wordToInt(first)*10 + wordToInt(last)
		fmt.Printf("(%d)\n", calib)
		calibrationSum += int(calib)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("(Part 2): Calibration Sum is %d \n", calibrationSum)
}
