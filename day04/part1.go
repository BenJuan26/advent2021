package main

import (
	"fmt"
	"strconv"
	"strings"

	advent "github.com/BenJuan26/advent2021"
)

func Part1() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	boards := []*Board{}

	i := 2
	for i < len(lines) {
		spaces := [][]*Space{}
		lastRow := i + 4
		for i <= lastRow {
			row := []*Space{}
			line := lines[i]
			split := strings.Fields(line)
			for _, col := range split {
				num, err := strconv.Atoi(col)
				if err != nil {
					panic(err)
				}
				row = append(row, &Space{Value: num})
			}
			spaces = append(spaces, row)
			i += 1
		}

		boards = append(boards, &Board{spaces})
		i += 1
	}

	numbersLine := lines[0]
	numbers := strings.Split(numbersLine, ",")

	for _, numString := range numbers {
		num, err := strconv.Atoi(numString)
		if err != nil {
			panic(err)
		}

		changedBoards := []*Board{}
		for _, board := range boards {
			changed := board.Mark(num)
			if changed {
				changedBoards = append(changedBoards, board)
			}
		}

		for _, board := range changedBoards {
			if board.IsWinner() {
				fmt.Println(board.Score() * num)
				return
			}
		}
	}

	fmt.Println("no winner!")
}

func main() {
	if advent.ShouldRunPart1() {
		Part1()
	}

	if advent.ShouldRunPart2() {
		Part2()
	}
}
