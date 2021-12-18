package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/bryan/adventofcode2021/internal/support"
	"github.com/stretchr/testify/require"
)

func TestCreatePair(t *testing.T) {
	tests := []struct {
		input   string
		wantErr bool
		want    Pair
	}{
		{
			input:   "[2,9]",
			wantErr: false,
			want: &ComplexPair{
				IntLeft:  NewNumber(2, 1),
				IntRight: NewNumber(9, 1)},
		},
		{
			input:   "[5,[[1,2],[3,7]]]",
			wantErr: false,
			want: &ComplexPair{
				IntLeft: NewNumber(5, 1),
				IntRight: &ComplexPair{
					IntLeft: &ComplexPair{
						IntLeft:  NewNumber(1, 3),
						IntRight: NewNumber(2, 3),
					},
					IntRight: &ComplexPair{
						IntLeft:  NewNumber(3, 3),
						IntRight: NewNumber(7, 3),
					},
				},
			},
		},
		{
			input:   "[[1,2],3]",
			wantErr: false,
			want: &ComplexPair{
				IntLeft: &ComplexPair{
					IntLeft:  NewNumber(1, 2),
					IntRight: NewNumber(2, 2),
				},
				IntRight: NewNumber(3, 1),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			got, err := CreatePair(tc.input)
			support.CheckError(t, tc.wantErr, err, func() {
				require.Equal(t, tc.want, got)
			})
		})
	}
}

func TestComplexPair_String(t *testing.T) {
	tests := []struct {
		input string
	}{
		{input: "[2,9]"},
		{input: "[5,[[1,2],[3,7]]]"},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			pair, err := CreatePair(tc.input)
			require.NoError(t, err)

			got := pair.String()
			require.Equal(t, tc.input, got)
		})
	}
}

func TestComplexPair_Add(t *testing.T) {
	tests := []struct {
		in    string
		other string
		want  string
	}{
		{
			in:    "[1,2]",
			other: "[[3,4],5]",
			want:  "[[1,2],[[3,4],5]]",
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s -> %s", tc.in, tc.want), func(t *testing.T) {
			a, err := CreatePair(tc.in)
			require.NoError(t, err)

			b, err := CreatePair(tc.other)
			require.NoError(t, err)

			c := a.Add(b)
			require.Equal(t, tc.want, c.String())
		})

	}
}

func TestComplexPair_Reduce(t *testing.T) {
	in := "[[[[[9,8],1],2],3],4]"
	pair, err := CreatePair(in)
	require.NoError(t, err)

	pair.Reduce()

	data, err := json.MarshalIndent(pair, "", "  ")
	require.NoError(t, err)

	fmt.Println(string(data))
}
