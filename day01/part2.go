package main

import (
	"fmt"
	"strconv"

	advent "github.com/BenJuan26/advent2021"
)

func Part2() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	count := 0
	var first, second, third, prevCount int
	hasPrevious := false
	for i, line := range lines {
		first = second
		second = third
		third, err = strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		currentCount := first + second + third

		if currentCount > prevCount && hasPrevious {
			count += 1
		}

		prevCount = currentCount
		if i > 2 {
			hasPrevious = true
		}
	}

	fmt.Println(count)
}
