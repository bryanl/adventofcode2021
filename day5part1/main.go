package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/bryan/adventofcode2021/internal/support"
)

func main() {
	support.SetupInput(run)
}

func run(data []string) error {
	var lines []Line
	for _, row := range data {
		lines = append(lines, NewLine(row))
	}

	points := CollectPoints(lines)
	filtered := FilterPoints(points)

	fmt.Println(len(filtered))
	return nil
}

func CollectPoints(lines []Line) map[Point]int {
	m := map[Point]int{}

	fmt.Println("line count", len(lines))

	for _, line := range lines {
		if line.IsHorizontal() || line.IsVertical() {
			for _, point := range line.Range() {
				m[point] += 1
			}
		}
	}

	return m
}

func FilterPoints(points map[Point]int) []Point {
	var out []Point

	for k, v := range points {
		if v >= 2 {
			out = append(out, k)
		}
	}

	return out
}

type Point struct {
	X int
	Y int
}

func NewPoint(in string) Point {
	parts := strings.Split(in, ",")
	if len(parts) != 2 {
		panic(fmt.Sprintf("%s is not a valid point", in))
	}

	return Point{
		X: support.ParseInt(parts[0]),
		Y: support.ParseInt(parts[1]),
	}
}

type Points []Point

func (points Points) MinX() int {
	var list []int
	for _, p := range points {
		list = append(list, p.X)
	}

	return min(list)
}

func (points Points) MaxX() int {
	var list []int
	for _, p := range points {
		list = append(list, p.X)
	}

	return max(list)
}

func (points Points) MinY() int {
	var list []int
	for _, p := range points {
		list = append(list, p.Y)
	}

	return min(list)
}

func (points Points) MaxY() int {
	var list []int
	for _, p := range points {
		list = append(list, p.Y)
	}

	return max(list)
}

type Line struct {
	Start Point
	End   Point
}

var lineRe = regexp.MustCompile(`^(?P<start>\d+,\d+) -> (?P<end>\d+,\d+)$`)

func NewLine(in string) Line {
	match := lineRe.FindStringSubmatch(in)
	result := map[string]string{}

	for i, name := range lineRe.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	start := NewPoint(result["start"])
	end := NewPoint(result["end"])

	line := Line{Start: start, End: end}
	return line
}

func (l *Line) Points() Points {
	return Points{l.Start, l.End}
}

type Slope struct {
	Rise     int
	Infinite bool
}

func (l *Line) Slope() Slope {
	minY := l.Points().MinY()
	maxY := l.Points().MaxY()
	y := maxY - minY

	minX := l.Points().MinX()
	maxX := l.Points().MaxX()
	x := maxX - minX

	if x == 0 {
		return Slope{Infinite: true}
	}

	return Slope{Rise: y / x}
}

func (l *Line) IsVertical() bool {
	return l.Start.Y == l.End.Y
}

func (l *Line) IsHorizontal() bool {
	return l.Start.X == l.End.X
}

func (l *Line) Range() []Point {
	var out []Point

	slope := l.Slope()
	if slope.Rise == 0 {
		if slope.Infinite {
			for y := l.Points().MinY(); y <= l.Points().MaxY(); y++ {
				out = append(out, Point{X: l.Start.X, Y: y})
			}
		} else {
			for x := l.Points().MinX(); x <= l.Points().MaxX(); x++ {
				out = append(out, Point{X: x, Y: l.Start.Y})
			}
		}
	} else {
		for x := l.Points().MinX(); x <= l.Points().MaxX(); x++ {
			y := slope.Rise * x
			out = append(out, Point{X: x, Y: y})
		}
	}

	return out
}

func min(in []int) int {
	var m int
	for i, x := range in {
		if i == 0 {
			m = x
		}
		if x < m {
			m = x
		}
	}

	return m
}

func max(in []int) int {
	var m int
	for i, x := range in {
		if i == 0 {
			m = x
		}

		if x > m {
			m = x
		}
	}

	return m
}
