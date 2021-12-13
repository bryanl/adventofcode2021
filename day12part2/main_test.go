package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/bryan/adventofcode2021/internal/support"
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

func TestAvailable(t *testing.T) {
	input := support.ReadFromDisk(t, "sample.txt")
	al := parse(input)

	tests := []struct {
		base    []string
		current string
		want    []string
	}{
		{
			base: []string{"start"},
			want: []string{"A", "b"},
		},
		{
			base: []string{"start", "A"},
			want: []string{"b", "c", "end"},
		},
		{
			base: []string{"start", "A", "c"},
			want: []string{"A"},
		},
		{
			base: []string{"start", "A", "c", "A", "c", "A"},
			want: []string{"b", "end"},
		},
		{
			base: []string{"start", "A", "b", "A"},
			want: []string{"b", "c", "end"},
		},
	}

	for _, tc := range tests {
		t.Run(strings.Join(tc.base, ","), func(t *testing.T) {
			got := Available(al, tc.base)
			require.Equal(t, tc.want, got)
		})
	}
}

func TestMaxedCave(t *testing.T) {
	tests := []struct {
		in        []string
		want      string
		wantFound bool
	}{
		{
			in:        []string{"start", "A", "c", "b"},
			want:      "",
			wantFound: false,
		},
		{
			in:        []string{"start", "A"},
			want:      "",
			wantFound: false,
		},
		{
			in:        []string{"start", "A", "c", "A", "c"},
			want:      "c",
			wantFound: true,
		},
		{
			in:        []string{"start", "A", "c", "A", "b", "A", "b"},
			want:      "b",
			wantFound: true,
		},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			got, wasFound := MaxedCave(tc.in)
			require.Equal(t, tc.wantFound, wasFound)
			require.Equal(t, tc.want, got)

		})
	}
}

func TestCountString(t *testing.T) {
	sl := []string{"start", "A", "c", "A", "c", "A"}
	s := "c"

	got := CountString(s, sl)
	want := 2
	require.Equal(t, want, got)
}
