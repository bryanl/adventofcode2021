package main

import (
	"container/ring"
	"fmt"
)

func main() {
	fmt.Println("part1 =", part1(10, 9))
}

func part1(p1, p2 int) int {
	die := NewIncDie()
	board := NewBoard(10, p1, p2, die)
	scoreMax := 1000

	win := false
	for !win {
		for i := 1; i <= 2; i++ {
			s := board.Play(i)
			if s >= scoreMax {
				win = true
				break
			}
		}
	}

	return board.Score()
}

type Die interface {
	Next() int
	Rolls() int
}

type IncDie struct {
	roll int
}

var _ Die = &IncDie{}

func NewIncDie() *IncDie {
	d := &IncDie{
		roll: 1,
	}

	return d
}

func (d *IncDie) Next() int {
	cur := d.roll
	d.roll += 1
	return cur
}

func (d *IncDie) Rolls() int {
	return d.roll
}

type Board struct {
	ring  *ring.Ring
	pos   map[int]*ring.Ring
	score map[int]int
	die   Die
}

func NewBoard(size, p1, p2 int, die Die) *Board {
	board := &Board{
		ring:  ring.New(size),
		score: map[int]int{},
		pos:   map[int]*ring.Ring{},
		die:   die,
	}

	r := board.ring
	n := r.Len()
	for i := 1; i <= n; i++ {
		if i == p1 {
			board.pos[1] = r
		}

		if i == p2 {
			board.pos[2] = r
		}

		r.Value = i
		r = r.Next()
	}

	return board
}

func (b *Board) Play(p int) int {
	r, ok := b.pos[p]
	if !ok {
		panic(fmt.Sprintf("unknown player %d", p))
	}

	for move := 0; move < 3; move++ {
		// fmt.Printf("p%d move[%d] roll[%d]\n", p, move, b.die.Rolls())
		next := b.die.Next()
		for i := 0; i < next; i++ {
			r = r.Next()
		}
	}

	b.score[p] += spaceValue(r)
	b.pos[p] = r
	return b.score[p]
}

func (b *Board) Space(p int) int {
	r, ok := b.pos[p]
	if !ok {
		panic(fmt.Sprintf("unknown player %d", p))
	}

	return spaceValue(r)
}

func (b *Board) Score() int {
	min := b.score[1]
	if b.score[2] < min {
		min = b.score[2]
	}

	return (b.die.Rolls() - 1) * min
}

func spaceValue(r *ring.Ring) int {
	return r.Value.(int)
}
