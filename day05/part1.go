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

func (line1 *Line) IsEqual(line2 *Line) bool {
	return line1.Point1.X == line2.Point1.X && line1.Point1.Y == line2.Point1.Y && line1.Point2.X == line2.Point2.X && line1.Point2.Y == line2.Point2.Y
}

func Part1() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	horizontal := make(map[int][]*Line)
	vertical := make(map[int][]*Line)
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
			horizontal[line.Point1.Y] = append(horizontal[line.Point1.Y], line)
		} else if line.IsVertical() {
			vertical[line.Point1.X] = append(vertical[line.Point1.X], line)
		}
	}

	// Initial overlaps
	// There could be duplicates, i.e. overlap lines can overlap each other
	hOverlap := make(map[int][]*Line)
	vOverlap := make(map[int][]*Line)

	// horizontal x horizontal
	for _, atY := range horizontal {
		for i := 0; i < len(atY)-1; i++ {
			for j := i + 1; j < len(atY); j++ {
				line1 := atY[i]
				line2 := atY[j]
				if !line1.IsEqual(line2) && line1.Point1.Y == line2.Point1.Y && line2.Right() >= line1.Left() && line1.Right() >= line2.Left() {
					left := Point{X: int(math.Max(float64(line1.Left()), float64(line2.Left()))), Y: line1.Point1.Y}
					right := Point{X: int(math.Min(float64(line1.Right()), float64(line2.Right()))), Y: line1.Point1.Y}
					hOverlap[line1.Point1.Y] = append(hOverlap[line1.Point1.Y], &Line{Point1: left, Point2: right})
				}
			}
		}
	}

	// vertical x vertical
	for _, atX := range vertical {
		for i := 0; i < len(atX)-1; i++ {
			for j := i + 1; j < len(atX); j++ {
				line1 := atX[i]
				line2 := atX[j]
				if !line1.IsEqual(line2) && line1.Point1.X == line2.Point1.X && line2.Bottom() >= line1.Top() && line1.Bottom() >= line2.Top() {
					top := Point{X: line1.Point1.X, Y: int(math.Max(float64(line1.Top()), float64(line2.Top())))}
					bottom := Point{X: line1.Point1.X, Y: int(math.Min(float64(line1.Bottom()), float64(line2.Bottom())))}
					vOverlap[line1.Point1.X] = append(vOverlap[line1.Point1.X], &Line{Point1: top, Point2: bottom})
				}
			}
		}
	}

	// horizontal x vertical
	for _, atY := range horizontal {
		for _, hLine := range atY {
			for x := hLine.Left(); x <= hLine.Right(); x++ {
				for _, vLine := range vertical[x] {
					if !hLine.IsEqual(vLine) && vLine.Bottom() >= hLine.Point1.Y && vLine.Top() <= hLine.Point1.Y {
						point := Point{X: vLine.Point1.X, Y: hLine.Point1.Y}
						// TODO: is defaulting to horizontal for a point okay?
						hOverlap[hLine.Point1.Y] = append(hOverlap[hLine.Point1.Y], &Line{Point1: point, Point2: point})
					}
				}
			}
		}
	}

	// Need to reduce hOverlap and vOverlap so that every line/point doesn't overlap with any other

	total := 0
	for _, atY := range horizontal {
		for _, line := range atY {
			fmt.Printf("(%d,%d),(%d,%d)\n", line.Point1.X, line.Point1.Y, line.Point2.X, line.Point2.Y)
			total += line.Right() - line.Left() + 1
		}
	}
	for _, atX := range vertical {
		for _, line := range atX {
			fmt.Printf("(%d,%d),(%d,%d)\n", line.Point1.X, line.Point1.Y, line.Point2.X, line.Point2.Y)
			total += line.Bottom() - line.Top() + 1
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
