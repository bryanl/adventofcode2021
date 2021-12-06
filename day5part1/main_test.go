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
		want  []Point
	}{
		{
			name:  "horizontal line",
			start: Point{X: 0, Y: 9},
			end:   Point{X: 5, Y: 9},
			want:  []Point{{X: 0, Y: 9}, {X: 1, Y: 9}, {X: 2, Y: 9}, {X: 3, Y: 9}, {X: 4, Y: 9}, {X: 5, Y: 9}},
		},
		{
			name:  "vertical line",
			start: Point{X: 1, Y: 0},
			end:   Point{X: 1, Y: 2},
			want:  []Point{{X: 1}, {X: 1, Y: 1}, {X: 1, Y: 2}},
		},
		{
			name:  "diagonal line",
			start: Point{X: 1, Y: 1},
			end:   Point{X: 3, Y: 3},
			want:  []Point{{X: 1, Y: 1}, {X: 2, Y: 2}, {X: 3, Y: 3}},
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
