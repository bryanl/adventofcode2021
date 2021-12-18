package main

import "fmt"

func main() {

}

type Point struct {
	X int
	Y int
}

func (p Point) String() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

type Target struct {
	Start Point
	End   Point
}

func NewTarget(start, end Point) *Target {
	startY := start.Y
	endY := end.Y

	if endY < startY {
		startY, endY = endY, startY
	}

	startX := start.X
	endX := end.X

	if endX < startX {
		startX, endX = endX, startX
	}

	t := &Target{
		Start: Point{startX, startY},
		End:   Point{endX, endY},
	}

	return t
}

func (t *Target) IsInGrid(p Point) (inGrid bool, hit bool) {
	// is point between 0,0 and the target end?

	// check up left
	if p.X >= t.Start.X && p.Y <= t.Start.Y {
		return true, t.IsHit(p)
	}

	// check bottom left
	if p.X >= t.Start.X && p.Y >= t.End.Y {
		return true, t.IsHit(p)
	}

	return false, false
}

func (t *Target) IsHit(p Point) bool {
	return false
}

type Trajectory struct {
	Current  Point
	Velocity Point
	Max      Point
}

func NewTrajectory(velocity Point) *Trajectory {
	t := &Trajectory{
		Velocity: velocity,
	}

	return t
}

func (t *Trajectory) Step() {
	t.Current.X += t.Velocity.X
	t.Current.Y += t.Velocity.Y

	if t.Current.X > t.Max.X {
		t.Max.X = t.Current.X
	}

	if t.Current.Y > t.Max.Y {
		t.Max.Y = t.Current.Y
	}

	if t.Velocity.X > 0 {
		t.Velocity.X -= 1
	} else if t.Velocity.X < 0 {
		t.Velocity.X += 1
	}

	t.Velocity.Y -= 1
}
