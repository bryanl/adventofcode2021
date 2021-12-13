package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	rows := []string{"6,10", "", "fold along y=7", "fold along x=5"}
	got := Parse(rows)
	want := &Input{
		Points: map[Point]int{
			NewPoint("6", "10"): 1,
		},
		Folds: []Fold{
			NewFold("y", "7"),
			NewFold("x", "5"),
		},
	}

	require.Equal(t, want, got)
}
