package main

import (
	"fmt"
	"strconv"
	"strings"

	advent "github.com/BenJuan26/advent2021"
)

func Part1() {
	daysToRun := 80

	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	stateSplit := strings.Split(lines[0], ",")
	state := [9]int{}
	for _, timer := range stateSplit {
		t, err := strconv.Atoi(timer)
		if err != nil {
			panic(err)
		}
		state[t] += 1
	}

	for d := 0; d < daysToRun; d += 1 {
		zeroes := state[0]
		for i := 1; i < 9; i++ {
			state[i-1] = state[i]
		}
		state[6] += zeroes
		state[8] = zeroes
	}

	total := 0
	for _, num := range state {
		total += num
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
