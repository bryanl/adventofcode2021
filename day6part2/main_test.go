package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSpawn(t *testing.T) {
	initial := Fish{Timer: 3, Life: 18}

	got := Spawn(initial)

	want := []Fish{
		{Timer: 8, Life: 14},
		{Timer: 8, Life: 7},
	}
	require.Equal(t, want, got)
}
