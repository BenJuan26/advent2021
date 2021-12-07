package main

import (
	"fmt"
	"strconv"

	advent "github.com/BenJuan26/advent2021"
)

func Part1() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	first := true
	count := 0
	var prev int
	for _, line := range lines {
		i, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		if i > prev && !first {
			count += 1
		}

		prev = i
		first = false
	}

	fmt.Println(count)
}

func main() {
	if advent.ShouldRunPart1() {
		Part1()
	}

	if advent.ShouldRunPart2() {
		Part2()
	}
}
