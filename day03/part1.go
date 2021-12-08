package main

import (
	"fmt"
	"math"

	advent "github.com/BenJuan26/advent2021"
)

func countBits(lines []string) (cols int, zeroes, ones []int) {
	cols = len(lines[0])

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

	return cols, zeroes, ones
}

func getGamma(cols int, zeroes, ones []int) (int, string) {
	gammaI := 0
	gammaS := ""

	for i := 0; i < cols; i++ {
		if ones[i] >= zeroes[i] {
			gammaI += int(math.Pow(2, float64(cols-i-1)))
			gammaS += "1"
		} else {
			gammaS += "0"
		}
	}

	return gammaI, gammaS
}

func getEpsilon(cols int, zeroes, ones []int) (int, string) {
	epI := 0
	epS := ""

	for i := 0; i < cols; i++ {
		if ones[i] < zeroes[i] {
			epI += int(math.Pow(2, float64(cols-i-1)))
			epS += "1"
		} else {
			epS += "0"
		}
	}

	return epI, epS
}

func Part1() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	cols, zeroes, ones := countBits(lines)
	gamma, _ := getGamma(cols, zeroes, ones)

	// epsilon := int(math.Pow(2, float64(cols))) - gamma - 1
	epsilon, _ := getEpsilon(cols, zeroes, ones)
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
