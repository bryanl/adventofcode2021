package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	text := []string{
		"on x=10..12,y=10..12,z=10..12",
		"on x=11..13,y=11..13,z=11..13",
		"off x=9..11,y=9..11,z=9..11",
		"on x=10..10,y=10..10,z=10..10",
	}

	got := Parse(text)

	want := []Step{
		{Status: "on", X: Range{Min: 10, Max: 12}, Y: Range{Min: 10, Max: 12}, Z: Range{Min: 10, Max: 12}},
		{Status: "on", X: Range{Min: 11, Max: 13}, Y: Range{Min: 11, Max: 13}, Z: Range{Min: 11, Max: 13}},
		{Status: "off", X: Range{Min: 9, Max: 11}, Y: Range{Min: 9, Max: 11}, Z: Range{Min: 9, Max: 11}},
		{Status: "on", X: Range{Min: 10, Max: 10}, Y: Range{Min: 10, Max: 10}, Z: Range{Min: 10, Max: 10}},
	}

	require.Equal(t, want, got)
}

func TestParseRow(t *testing.T) {
	row := "on x=10..12,y=10..12,z=10..12"

	want := Step{
		Status: "on",
		X:      Range{Min: 10, Max: 12},
		Y:      Range{Min: 10, Max: 12},
		Z:      Range{Min: 10, Max: 12},
	}

	got := ParseRow(row)

	require.Equal(t, want, got)
}

func TestSet_Add(t *testing.T) {
	tests := []struct {
		name string
		a    Range
		b    Range
		want Range
	}{
		{
			name: "simple",
			a:    Range{Min: 10, Max: 12},
			b:    Range{Min: 11, Max: 13},
			want: Range{Min: 10, Max: 13},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := NewSet(tc.a)
			s.Add(tc.b)
			got := s.Range()
			require.Equal(t, tc.want, got)
		})
	}
}

func TestSet_Sub(t *testing.T) {
	tests := []struct {
		name string
		a    Range
		b    Range
		want Range
	}{
		{
			name: "simple",
			a:    Range{10, 13},
			b:    Range{9, 11},
			want: Range{12, 13},
		},
		{
			name: "simple reverse",
			a:    Range{9, 11},
			b:    Range{10, 13},
			want: Range{9, 9},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := NewSet(tc.a)
			s.Sub(tc.b)
			got := s.Range()
			require.Equal(t, tc.want, got)
		})
	}
}
