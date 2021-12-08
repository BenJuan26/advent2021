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

	boards := buildBoards(lines)

	numbersLine := lines[0]
	numbers := strings.Split(numbersLine, ",")

	for _, numString := range numbers {
		num, err := strconv.Atoi(numString)
		if err != nil {
			panic(err)
		}

		changedBoards := []*Board{}
		remainingBoards := []*Board{}
		for _, board := range boards {
			changed := board.Mark(num)
			if changed {
				changedBoards = append(changedBoards, board)
			} else {
				remainingBoards = append(remainingBoards, board)
			}
		}

		for _, board := range changedBoards {
			if !board.IsWinner() {
				remainingBoards = append(remainingBoards, board)
			} else {
				if len(boards) == 1 {
					fmt.Println(boards[0].Score() * num)
					return
				}
			}
		}

		boards = remainingBoards
	}
}
