package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTarget_IsHit(t *testing.T) {
	tests := []struct {
		name  string
		start Point
		end   Point
		point Point
		want  bool
	}{
		{
			name:  "hit",
			start: Point{20, -5},
			end:   Point{30, -10},
			point: Point{21, -7},
			want:  true,
		},
		{
			name:  "x out of bounds",
			start: Point{20, -5},
			end:   Point{30, -10},
			point: Point{17, -7},
			want:  false,
		},
		{
			name:  "y out of bounds",
			start: Point{20, -5},
			end:   Point{30, -10},
			point: Point{21, 7},
			want:  false,
		},
		{
			name:  "x and y out of bounds",
			start: Point{20, -5},
			end:   Point{30, -10},
			point: Point{-21, 7},
			want:  false,
		},
		{
			name:  "negative",
			start: Point{-3, -5},
			end:   Point{-5, -3},
			point: Point{-4, 4},
			want:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			target := NewTarget(tc.start, tc.end)
			got := target.IsHit(tc.point)
			require.Equal(t, tc.want, got)
		})
	}
}

func TestTarget_IsInGrid(t *testing.T) {
	tests := []struct {
		name       string
		start      Point
		end        Point
		point      Point
		wantInGrid bool
		wantHint   bool
	}{
		{
			name:       "top left hit",
			start:      Point{-5, 5},
			end:        Point{-3, 3},
			point:      Point{-5, 3},
			wantInGrid: true,
			wantHint:   true,
		},
		{
			name:       "top left hit",
			end:        Point{-5, -3},
			start:      Point{-3, -5},
			point:      Point{-4, 4},
			wantInGrid: true,
			wantHint:   true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			target := NewTarget(tc.start, tc.end)
			gotInGrid, gotHit := target.IsInGrid(tc.point)
			require.Equal(t, tc.wantInGrid, gotInGrid, "in grid")
			require.Equal(t, tc.wantHint, gotHit, "hit")
		})
	}
}

func TestTrajectory_Step(t *testing.T) {
	tests := []struct {
		velocity Point
		steps    int
		want     Point
	}{
		{
			velocity: Point{X: 7, Y: 2},
			steps:    1,
			want:     Point{7, 2},
		},
		{
			velocity: Point{X: 7, Y: 2},
			steps:    2,
			want:     Point{13, 3},
		},
		{
			velocity: Point{X: 7, Y: 2},
			steps:    3,
			want:     Point{18, 3},
		},
		{
			velocity: Point{X: 7, Y: 2},
			steps:    4,
			want:     Point{22, 2},
		},
		{
			velocity: Point{X: 7, Y: 2},
			steps:    5,
			want:     Point{25, 0},
		},
		{
			velocity: Point{X: 7, Y: 2},
			steps:    6,
			want:     Point{27, -3},
		},
		{
			velocity: Point{X: 7, Y: 2},
			steps:    7,
			want:     Point{28, -7},
		},
	}

	for _, tc := range tests {
		name := fmt.Sprintf("%s -> %d", tc.velocity, tc.steps)
		t.Run(name, func(t *testing.T) {
			tr := NewTrajectory(tc.velocity)

			for i := 0; i < tc.steps; i++ {
				tr.Step()
			}
			got := tr.Current
			require.Equal(t, tc.want, got)
		})
	}
}
