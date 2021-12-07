package main

import (
	"fmt"
	"math"

	advent "github.com/BenJuan26/advent2021"
)

func Part1() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	cols := len(lines[0])

	zeroes := []int{}
	ones := []int{}
	for i := 0; i < cols; i++ {
		zeroes = append(zeroes, 0)
		ones = append(ones, 0)
	}

	for _, line := range lines {
		for i := 0; i < cols; i++ {
			if line[i] == "0"[0] {
				zeroes[i] += 1
			} else {
				ones[i] += 1
			}
		}
	}

	gamma := 0
	for i := 0; i < cols; i++ {
		if ones[i] > zeroes[i] {
			gamma += int(math.Pow(2, float64(cols-i-1)))
		}
	}

	epsilon := int(math.Pow(2, float64(cols))) - gamma - 1
	fmt.Println(gamma * epsilon)
}

func main() {
	if advent.ShouldRunPart1() {
		Part1()
	}

	if advent.ShouldRunPart2() {
		Part2()
	}
}
