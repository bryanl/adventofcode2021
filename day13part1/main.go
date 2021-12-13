package main

import (
	"fmt"
	"regexp"

	"github.com/bryan/adventofcode2021/internal/support"
)

func main() {
	support.SetupInput(run)
}

func run(rows []string) error {
	input := Parse(rows)

	fmt.Println(input.CountVisible())
	input.Fold()
	fmt.Println(input.CountVisible())
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
	FoldHorizontal FoldDirection = iota
	FoldVertical
)

type Fold struct {
	Direction FoldDirection
	At        int
}

func NewFold(typeStr, at string) Fold {
	dir := FoldHorizontal
	if typeStr == "x" {
		dir = FoldVertical
	}

	return Fold{
		Direction: dir,
		At:        support.ParseInt(at),
	}
}

func (f *Fold) Do(points Points) {
	if f.Direction == FoldHorizontal {
		f.doHorizontal(points)
		return
	}

	f.doVertical(points)
}

func (f *Fold) doHorizontal(points Points) {
	for oldPoint := range points {
		if oldPoint.Y > f.At {
			newY := f.At - (oldPoint.Y - f.At)
			newPoint := Point{X: oldPoint.X, Y: newY}
			points[oldPoint] -= 1
			points[newPoint] += 1
		}
	}
}

func (f *Fold) doVertical(points Points) {
	for oldPoint := range points {
		if oldPoint.X > f.At {
			newX := f.At - (oldPoint.X - f.At)
			newPoint := Point{X: newX, Y: oldPoint.Y}
			points[oldPoint] -= 1
			points[newPoint] += 1
		}
	}
}

type Input struct {
	Points map[Point]int
	Folds  []Fold
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

func (input *Input) AddFold(f Fold) {
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
	var fold Fold
	fold, input.Folds = input.Folds[0], input.Folds[1:]

	fold.Do(input.Points)
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
