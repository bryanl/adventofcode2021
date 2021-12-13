package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_visit(t *testing.T) {
	al := adjacencyList{
		"c":     []string{"A"},
		"d":     []string{"b"},
		"end":   []string{"A", "b"},
		"start": []string{"A", "b"},
		"A":     []string{"start", "c", "b", "end"},
		"b":     []string{"start", "A", "d", "end"},
	}

	previous := []string{"start", "A", "c", "A", "b"}

	got := visit(al, previous, 1)

	want := [][]string{
		{"start", "A", "c", "A", "b", "A", "end"},
		{"start", "A", "c", "A", "b", "end"},
	}

	require.Equal(t, want, got)
}
