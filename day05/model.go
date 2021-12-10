package main

import "math"

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

func (l *Line) IsAscending() bool {
	if l.Point1.X == l.Left() {
		return l.Point1.Y == l.Bottom()
	}

	return l.Point1.Y == l.Top()
}
