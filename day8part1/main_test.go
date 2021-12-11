package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadInput(t *testing.T) {
	text := "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf"
	got := ReadInput(text)

	want := Input{
		SignalPatterns: []string{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"},
		OutputValue:    []string{"cdfeb", "fcadb", "cdfeb", "cdbaf"},
	}

	require.Equal(t, want, got)
}

func TestCountKnown(t *testing.T) {
	input := Input{
		SignalPatterns: []string{"aa", "aaa", "a"},
		OutputValue:    []string{"aa"},
	}

	got := CountKnown(input)

	require.Equal(t, 1, got)
}

func TestGuessByLength(t *testing.T) {
	tests := []struct {
		in   string
		want int
		ok   bool
	}{
		{
			in:   "ab",
			want: 1,
			ok:   true,
		},
		{
			in:   "abcd",
			want: 4,
			ok:   true,
		},
		{
			in:   "abc",
			want: 7,
			ok:   true,
		},
		{
			in:   "abcdefg",
			want: 8,
			ok:   true,
		},
		{
			in: "abcde",
			ok: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.in, func(t *testing.T) {
			got, ok := GuessByLength(tc.in)
			if tc.ok {
				require.True(t, ok)
				require.Equal(t, tc.want, got)
			} else {
				require.False(t, ok)
			}
		})
	}
}
