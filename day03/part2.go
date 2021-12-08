package main

import (
	"fmt"
	"math"

	advent "github.com/BenJuan26/advent2021"
)

func Part2() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	cols := len(lines[0])
	col := 0
	for len(lines) > 1 {
		_, zeroes, ones := countBits(lines)
		_, gamma := getGamma(cols, zeroes, ones)

		newLines := []string{}
		for _, line := range lines {
			if line[col] == gamma[col] {
				newLines = append(newLines, line)
			}
		}

		lines = newLines
		col += 1
	}

	o2S := lines[0]
	o2I := 0
	for i := 0; i < cols; i++ {
		if o2S[i] == "1"[0] {
			o2I += int(math.Pow(2, float64(cols-i-1)))
		}
	}

	lines, err = advent.ReadInput()
	if err != nil {
		panic(err)
	}

	col = 0
	for len(lines) > 1 {
		_, zeroes, ones := countBits(lines)
		_, ep := getEpsilon(cols, zeroes, ones)

		newLines := []string{}
		for _, line := range lines {
			if line[col] == ep[col] {
				newLines = append(newLines, line)
			}
		}

		lines = newLines
		col += 1
	}

	co2S := lines[0]
	co2I := 0
	for i := 0; i < cols; i++ {
		if co2S[i] == "1"[0] {
			co2I += int(math.Pow(2, float64(cols-i-1)))
		}
	}

	fmt.Println(co2I * o2I)
}
