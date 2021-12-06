package main

import (
	"strings"
	"testing"

	"github.com/bryan/adventofcode2021/internal/support"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewGameBuilder(t *testing.T) {
	data, err := support.ReadData("sample.txt")
	require.NoError(t, err)

	builder := NewGameBuilder(data)

	game, err := builder.Build()
	require.NoError(t, err)

	assert.Len(t, game.Draws, 27)
	assert.Len(t, game.Cards, 3)
}

func TestCard_IsWin(t *testing.T) {

	tests := []struct {
		name  string
		draws string
		isWin bool
	}{
		{
			name:  "row win 1",
			draws: "22,13,17,11,0",
			isWin: true,
		},
		{
			name:  "row win 2",
			draws: "21,9,14,16,7",
			isWin: true,
		},
		{
			name:  "col win 1",
			draws: "22,8,21,6,1",
			isWin: true,
		},
		{
			name:  "col win 2",
			draws: "17,23,14,3,20",
			isWin: true,
		},
		{
			name:  "diag win 1",
			draws: "22,2,14,18,19",
			isWin: true,
		},
		{
			name:  "diag win 2",
			draws: "1,10,14,4,0",
			isWin: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			card := createCard(t)

			draws := strings.Split(tc.draws, ",")
			for _, draw := range draws {
				card.Select(draw)
			}

			require.Equal(t, tc.isWin, card.IsWin())
			card.IsWin()
		})
	}
}

func TestCard_UnmarkedSum(t *testing.T) {

}

func createCard(t *testing.T) *Card {
	in := []string{
		"22", "13", "17", "11", "0",
		"8", "2", "23", "4", "24",
		"21", "9", "14", "16", "7",
		"6", "10", "3", "18", "5",
		"1", "12", "20", "15", "19",
	}

	card, err := NewCard(in)
	if err != nil {
		t.Errorf("unable to create card: %v", err)
	}

	require.Equal(t, 5, card.Side)
	require.Len(t, card.Numbers, len(in))

	return card
}
