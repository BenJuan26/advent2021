package main

import (
	"fmt"
	"strings"

	advent "github.com/BenJuan26/advent2021"
)

func Part1() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	ones := 0
	fours := 0
	sevens := 0
	eights := 0
	for _, line := range lines {
		lineSplit := strings.Split(line, " | ")
		output := lineSplit[1]
		digits := strings.Split(output, " ")
		for _, digit := range digits {
			switch len(digit) {
			case 2:
				ones += 1
			case 3:
				sevens += 1
			case 4:
				fours += 1
			case 7:
				eights += 1
			}
		}
	}

	fmt.Println(ones + fours + sevens + eights)
}

func main() {
	if advent.ShouldRunPart1() {
		Part1()
	}

	if advent.ShouldRunPart2() {
		Part2()
	}
}
