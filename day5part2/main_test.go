package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewPoint(t *testing.T) {
	in := "0,9"

	want := Point{X: 0, Y: 9}
	actual := NewPoint(in)
	require.Equal(t, want, actual)
}

func TestNewLine(t *testing.T) {
	in := "0,9 -> 5,9"
	actual := NewLine(in)
	want := Line{Start: Point{Y: 9}, End: Point{X: 5, Y: 9}}
	require.Equal(t, want, actual)
}

func TestLine_Range(t *testing.T) {
	tests := []struct {
		name  string
		start Point
		end   Point
		want  Points
	}{
		{
			name:  "horizontal line",
			start: Point{X: 0, Y: 9},
			end:   Point{X: 5, Y: 9},
			want:  Points{{X: 0, Y: 9}, {X: 1, Y: 9}, {X: 2, Y: 9}, {X: 3, Y: 9}, {X: 4, Y: 9}, {X: 5, Y: 9}},
		},
		{
			name:  "vertical line",
			start: Point{X: 1, Y: 0},
			end:   Point{X: 1, Y: 2},
			want:  Points{{X: 1}, {X: 1, Y: 1}, {X: 1, Y: 2}},
		},
		{
			name:  "diagonal line 1",
			start: Point{X: 8, Y: 0},
			end:   Point{X: 0, Y: 8},
			want:  Points{{Y: 8}, {X: 1, Y: 7}, {X: 2, Y: 6}, {X: 3, Y: 5}, {X: 4, Y: 4}, {X: 5, Y: 3}, {X: 6, Y: 2}, {X: 7, Y: 1}, {X: 8}},
		},
		{
			name:  "diagonal line 2",
			start: Point{X: 0, Y: 8},
			end:   Point{X: 8, Y: 0},
			want:  Points{{Y: 8}, {X: 1, Y: 7}, {X: 2, Y: 6}, {X: 3, Y: 5}, {X: 4, Y: 4}, {X: 5, Y: 3}, {X: 6, Y: 2}, {X: 7, Y: 1}, {X: 8}},
		},
		{
			name:  "diagonal line 3",
			start: Point{},
			end:   Point{X: 8, Y: 8},
			want:  Points{{}, {X: 1, Y: 1}, {X: 2, Y: 2}, {X: 3, Y: 3}, {X: 4, Y: 4}, {X: 5, Y: 5}, {X: 6, Y: 6}, {X: 7, Y: 7}, {X: 8, Y: 8}},
		},
		{
			name:  "diagonal line 4",
			start: Point{X: 1, Y: 1},
			end:   Point{X: 3, Y: 3},
			want:  Points{{X: 1, Y: 1}, {X: 2, Y: 2}, {X: 3, Y: 3}},
		},
		{
			name:  "diagonal line 5",
			start: Point{X: 9, Y: 7},
			end:   Point{X: 7, Y: 9},
			want:  Points{{X: 7, Y: 9}, {X: 8, Y: 8}, {X: 9, Y: 7}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			line := Line{Start: tc.start, End: tc.end}
			actual := line.Range()
			require.Equal(t, tc.want, actual)
		})
	}

}
