package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	advent "github.com/BenJuan26/advent2021"
)

func SeriesSum(n int) int {
	return int(n * (2 + (n - 1)) / 2)
}

func Part2() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	posStrings := strings.Split(lines[0], ",")
	positions := []int{}
	runningTotal := 0
	for _, posString := range posStrings {
		pos, err := strconv.Atoi(posString)
		if err != nil {
			panic(err)
		}
		positions = append(positions, pos)
		runningTotal += pos
	}

	avg := int(math.Round(float64(runningTotal / len(positions))))

	total := 0
	for _, pos := range positions {
		total += SeriesSum(int(math.Abs(float64(pos - avg))))
	}

	fmt.Println(total)
}
