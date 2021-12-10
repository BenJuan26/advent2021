package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	advent "github.com/BenJuan26/advent2021"
)

func Part1() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	posStrings := strings.Split(lines[0], ",")
	positions := []int{}
	for _, posString := range posStrings {
		pos, err := strconv.Atoi(posString)
		if err != nil {
			panic(err)
		}
		positions = append(positions, pos)
	}

	sort.Ints(positions)

	middle := int(len(positions) / 2)
	median := positions[middle]
	total := 0

	for _, pos := range positions {
		total += int(math.Abs(float64(pos - median)))
	}

	fmt.Println(total)
}

func main() {
	if advent.ShouldRunPart1() {
		Part1()
	}

	if advent.ShouldRunPart2() {
		Part2()
	}
}
