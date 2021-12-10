package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	advent "github.com/BenJuan26/advent2021"
)

type Point struct {
	X int
	Y int
}

type Line struct {
	Point1 Point
	Point2 Point
}

func (l *Line) IsHorizontal() bool {
	return l.Point1.Y == l.Point2.Y
}

func (l *Line) IsVertical() bool {
	return l.Point1.X == l.Point2.X
}

func (l *Line) Left() int {
	return int(math.Min(float64(l.Point1.X), float64(l.Point2.X)))
}

func (l *Line) Right() int {
	return int(math.Max(float64(l.Point1.X), float64(l.Point2.X)))
}

func (l *Line) Top() int {
	return int(math.Min(float64(l.Point1.Y), float64(l.Point2.Y)))
}

func (l *Line) Bottom() int {
	return int(math.Max(float64(l.Point1.Y), float64(l.Point2.Y)))
}

func Part1() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	grid := [990][990]int{}
	total := 0
	for _, l := range lines {
		split := strings.Split(l, " -> ")
		point1String := split[0]
		point2String := split[1]

		point1Split := strings.Split(point1String, ",")
		point2Split := strings.Split(point2String, ",")

		point1X, err := strconv.Atoi(point1Split[0])
		if err != nil {
			panic(err)
		}
		point1Y, err := strconv.Atoi(point1Split[1])
		if err != nil {
			panic(err)
		}
		point2X, err := strconv.Atoi(point2Split[0])
		if err != nil {
			panic(err)
		}
		point2Y, err := strconv.Atoi(point2Split[1])
		if err != nil {
			panic(err)
		}

		point1 := Point{X: point1X, Y: point1Y}
		point2 := Point{X: point2X, Y: point2Y}

		line := &Line{Point1: point1, Point2: point2}

		if line.IsHorizontal() {
			for i := line.Left(); i <= line.Right(); i++ {
				grid[line.Point1.Y][i] += 1
				if grid[line.Point1.Y][i] == 2 {
					total += 1
				}
			}
		} else if line.IsVertical() {
			for i := line.Top(); i <= line.Bottom(); i++ {
				grid[i][line.Point1.X] += 1
				if grid[i][line.Point1.X] == 2 {
					total += 1
				}
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
