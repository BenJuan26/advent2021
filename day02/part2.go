package main

import (
	"fmt"
	"strconv"
	"strings"

	advent "github.com/BenJuan26/advent2021"
)

func Part2() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	x := 0
	y := 0
	aim := 0
	for _, line := range lines {
		split := strings.Split(line, " ")
		if len(split) != 2 {
			panic(fmt.Sprintf("split values: want 2 got %d", len(split)))
		}

		direction := split[0]
		distance, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}

		switch direction {
		case "down":
			aim += distance
		case "up":
			aim -= distance
		case "forward":
			x += distance
			y += distance * aim
		default:
			panic("unknown direction " + direction)
		}
	}

	fmt.Println(x * y)
}
