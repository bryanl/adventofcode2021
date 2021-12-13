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

func run(rows []string) error {
	input := Parse(rows)

	input.Fold()
	fmt.Println(input.Print())

	return nil
}

type Point struct {
	X int
	Y int
}

type Points map[Point]int

func NewPoint(x, y string) Point {
	return Point{
		X: support.ParseInt(x),
		Y: support.ParseInt(y),
	}
}

type FoldDirection int

const (
	FoldUp FoldDirection = iota
	FoldLeft
)

type Folder struct {
	Direction FoldDirection
	At        int
}

func NewFold(typeStr, at string) Folder {
	dir := FoldUp
	if typeStr == "x" {
		dir = FoldLeft
	}

	return Folder{
		Direction: dir,
		At:        support.ParseInt(at),
	}
}

func (f *Folder) Fold(points Points) {
	if f.Direction == FoldUp {
		f.foldUp(points)
		return
	}

	f.foldLeft(points)
}

func (f *Folder) foldUp(points Points) {
	for oldPoint, v := range points {
		if oldPoint.Y > f.At && v > 0 {
			newY := f.At - (oldPoint.Y - f.At)
			newPoint := Point{X: oldPoint.X, Y: newY}
			delete(points, oldPoint)
			points[newPoint] = 1
		}
	}
}

func (f *Folder) foldLeft(points Points) {
	for oldPoint, v := range points {
		if oldPoint.X > f.At && v > 0 {
			newX := f.At - (oldPoint.X - f.At)
			newPoint := Point{X: newX, Y: oldPoint.Y}
			delete(points, oldPoint)
			points[newPoint] = 1
		}
	}
}

type Input struct {
	Points map[Point]int
	Folds  []Folder
}

func NewInput() *Input {
	input := &Input{
		Points: map[Point]int{},
	}

	return input
}

func (input *Input) AddPoint(p Point) {
	input.Points[p] += 1
}

func (input *Input) AddFold(f Folder) {
	input.Folds = append(input.Folds, f)
}

func (input *Input) CountVisible() int {
	sum := 0
	for _, v := range input.Points {
		if v > 0 {
			sum += 1
		}
	}

	return sum
}

func (input *Input) Fold() {
	for _, fold := range input.Folds {
		fold.Fold(input.Points)
	}
}

func (input *Input) Count() int {
	sum := 0
	for _, i := range input.Points {
		if i > 0 {
			sum += 1
		}
	}
	return sum
}

func (input *Input) Print() string {
	w := 0
	h := 0

	for point, v := range input.Points {
		if v < 1 {
			continue
		}

		if point.X > w {
			w = point.X
		}

		if point.Y > h {
			h = point.Y
		}
	}

	var sb strings.Builder

	for y := 0; y <= h; y++ {
		for x := 0; x <= w; x++ {
			p := Point{X: x, Y: y}
			m := " "
			if v, ok := input.Points[p]; ok && v > 0 {
				m = "*"
			}
			sb.WriteString(m)
		}

		sb.WriteString("\n")
	}

	return sb.String()
}

var (
	reCoord = regexp.MustCompile(`^(?P<x>\d+),(?P<y>\d+)$`)
	reFold  = regexp.MustCompile(`(?P<type>[xy])=(?P<at>\d+)$`)
)

func Parse(rows []string) *Input {
	i := NewInput()

	for _, row := range rows {
		switch {
		case reCoord.MatchString(row):
			res := getMatches(row, reCoord)
			i.AddPoint(NewPoint(res["x"], res["y"]))
		case reFold.MatchString(row):
			res := getMatches(row, reFold)
			i.AddFold(NewFold(res["type"], res["at"]))
		}

	}

	return i
}

func getMatches(s string, re *regexp.Regexp) map[string]string {
	match := re.FindStringSubmatch(s)
	result := map[string]string{}

	for i, name := range re.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	return result
}
