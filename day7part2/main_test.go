package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTriangleSum(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{n: 1, want: 1},
		{n: 3, want: 6},
		{n: 11, want: 66},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprint(tc.n), func(t *testing.T) {
			got := TriangleSum(tc.n)
			require.Equal(t, tc.want, got)
		})
	}
}
