package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		rows    []string
		want    []int
		wantErr bool
	}{
		{
			name:    "in general",
			rows:    []string{"012", "345", "678"},
			want:    []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
			wantErr: false,
		},
		{
			name:    "not square",
			rows:    []string{"012"},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			got, err := Parse(tc.rows)
			if tc.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tc.want, got)
		})
	}
}

func TestNeighbors(t *testing.T) {
	grid := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	tests := []struct {
		name string
		grid []int
		pos  int
		want []int
	}{
		{
			name: "top left",
			grid: grid,
			pos:  0,
			want: []int{1, 3},
		},
		{
			name: "center",
			grid: grid,
			pos:  4,
			want: []int{1, 3, 5, 7},
		},
		{
			name: "bottom right",
			grid: grid,
			pos:  8,
			want: []int{5, 7},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Neighbors(tc.grid, tc.pos)
			require.Equal(t, tc.want, got)
		})
	}
}

func TestHorizontalSum(t *testing.T) {
	grid := []int{
		0, 1, 1,
		0, 4, 1,
		0, 0, 1,
	}

	tests := []struct {
		name    string
		grid    []int
		start   int
		end     int
		want    int
		wantErr bool
	}{

		{
			name:    "0 -> 2",
			grid:    grid,
			start:   0,
			end:     2,
			want:    2,
			wantErr: false,
		},
		{
			name:    "5 -> 4",
			grid:    grid,
			start:   5,
			end:     4,
			want:    4,
			wantErr: false,
		},
		{
			name:    "0 -> 3",
			grid:    grid,
			start:   0,
			end:     3,
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := HorizontalSum(tc.grid, tc.start, tc.end)

			if tc.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tc.want, got)
		})
	}
}

func TestVerticalSum(t *testing.T) {
	grid := []int{
		0, 1, 1,
		0, 4, 1,
		0, 0, 1,
	}

	tests := []struct {
		name    string
		grid    []int
		start   int
		end     int
		want    int
		wantErr bool
	}{
		{
			name:    "2 -> 8",
			grid:    grid,
			start:   2,
			end:     8,
			want:    2,
			wantErr: false,
		},
		{
			name:    "7 -> 1",
			grid:    grid,
			start:   7,
			end:     1,
			want:    5,
			wantErr: false,
		},
		{
			name:    "0 -> 1",
			grid:    grid,
			start:   0,
			end:     1,
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := VerticalSum(tc.grid, tc.start, tc.end)

			if tc.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tc.want, got)
		})
	}
}
