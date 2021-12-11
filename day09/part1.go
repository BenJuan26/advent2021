package main

import (
	"fmt"
	"strconv"

	advent "github.com/BenJuan26/advent2021"
)

func IsLowest(grid [][]int, x, y int) bool {
	w := len(grid[0])
	h := len(grid)

	val := grid[y][x]

	// above
	if y > 0 && grid[y-1][x] <= val {
		return false
	}

	// below
	if y < h-1 && grid[y+1][x] <= val {
		return false
	}

	// left
	if x > 0 && grid[y][x-1] <= val {
		return false
	}

	// right
	if x < w-1 && grid[y][x+1] <= val {
		return false
	}

	return true
}

func Part1() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	grid := [][]int{}
	for _, line := range lines {
		row := []int{}
		for _, char := range line {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				panic(err)
			}
			row = append(row, num)
		}
		grid = append(grid, row)
	}

	total := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if IsLowest(grid, x, y) {
				total += grid[y][x] + 1
			}
		}
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
