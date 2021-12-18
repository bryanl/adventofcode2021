package main

import (
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
				left:  NewNumber(2, 1),
				right: NewNumber(9, 1)},
		},
		{
			input:   "[5,[[1,2],[3,7]]]",
			wantErr: false,
			want: &ComplexPair{
				left: NewNumber(5, 1),
				right: &ComplexPair{
					left: &ComplexPair{
						left:  NewNumber(1, 3),
						right: NewNumber(2, 3),
					},
					right: &ComplexPair{
						left:  NewNumber(3, 3),
						right: NewNumber(7, 3),
					},
				},
			},
		},
		{
			input:   "[[1,2],3]",
			wantErr: false,
			want: &ComplexPair{
				left: &ComplexPair{
					left:  NewNumber(1, 2),
					right: NewNumber(2, 2),
				},
				right: NewNumber(3, 1),
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
