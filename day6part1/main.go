package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bryan/adventofcode2021/internal/support"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	support.SetupInput(run)
}

func run(input []string) error {
	initialStr := input[0]
	var ages []int
	for _, s := range strings.Split(initialStr, ",") {
		ages = append(ages, support.ParseInt(s))
	}

	w := NewWorld(ages)
	w.Simulate(256)

	spew.Dump(w.Summary())

	return nil
}

type World struct {
	Population []*Fish
}

func NewWorld(ages []int) *World {
	w := &World{}

	for _, age := range ages {
		w.Create(age)
	}

	return w
}

func (w *World) Create(age int) {
	f := &Fish{Timer: age}
	w.Population = append(w.Population, f)
}

func (w *World) Simulate(days int) {
	for i := 0; i < days; i++ {
		start := time.Now()
		w.Tick()
		elapsed := time.Since(start)
		fmt.Printf("%d %d - %s\n", i, len(w.Population), elapsed)
	}
}

func (w *World) Tick() {
	for i := range w.Population {
		w.Population[i].Age(w)
	}
}

func (w *World) Summary() int {
	return len(w.State())
}

func (w *World) State() []int {
	var out []int
	for _, f := range w.Population {
		out = append(out, f.Timer)
	}
	return out
}

type Fish struct {
	Timer int
}

func (f *Fish) Age(w *World) {
	if f.Timer == 0 {
		f.Timer = 6
		f.Spawn(w)
	} else {
		f.Timer -= 1
	}
}

func (f *Fish) Spawn(w *World) {
	w.Create(8)
}
