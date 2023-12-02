package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const InputFile = "day01/input"

func main() {
	part1()
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
