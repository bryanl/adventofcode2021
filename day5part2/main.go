package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"sort"
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
	if err := SaveLines(lines, points); err != nil {
		return err
	}

	// PrintPoints(points)
	// filtered := FilterPoints(points)
	// fmt.Println(len(filtered))
	return nil
}

type Output struct {
	Lines         []Line                     `json:"lines"`
	Intersections map[string]OutIntersection `json:"intersections"`
}

type OutIntersection struct {
	Point Point `json:"point"`
	Count int   `json:"count"`
}

func SaveLines(lines []Line, intersections map[Point]int) error {
	f, err := os.OpenFile("output.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}

	defer f.Close()

	list := map[string]OutIntersection{}

	for k, v := range intersections {
		list[k.String()] = OutIntersection{
			Point: k,
			Count: v,
		}

	}

	o := &Output{
		Lines:         lines,
		Intersections: list,
	}

	e := json.NewEncoder(f)
	e.SetIndent("", "  ")
	return e.Encode(o)
}

func PrintPoints(points map[Point]int) {
	maxX := 0
	maxY := 0

	for p := range points {
		if p.X > maxX {
			maxX = p.X
		}

		if p.Y > maxY {
			maxY = p.Y
		}
	}

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			point := Point{X: x, Y: y}
			count := points[point]
			if count == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(count)
			}
		}
		fmt.Println()
	}

}

func CollectPoints(lines []Line) map[Point]int {
	m := map[Point]int{}

	for _, line := range lines {
		for _, point := range line.Range() {
			m[point] += 1
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
	X int `json:"x"`
	Y int `json:"y"`
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

func (p *Point) String() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)

}

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

func (points Points) Sort() {
	sort.Slice(points, func(i, j int) bool {
		return points[i].X < points[j].X
	})
}

type Line struct {
	Start Point `json:"start"`
	End   Point `json:"end"`
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
	y := l.Start.Y - l.End.Y
	x := l.Start.X - l.End.X

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

func (l *Line) Range() Points {
	var out Points

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
		start := l.Start
		end := l.End

		if start.X > end.X {
			start, end = end, start
		}

		for x := start.X; x <= end.X; x++ {
			dist := start.X - x
			y := start.Y + (dist)*-slope.Rise

			out = append(out, Point{X: x, Y: y})
		}

	}

	out.Sort()

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
