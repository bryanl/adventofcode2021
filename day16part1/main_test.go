package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInput(t *testing.T) {
	rows := []string{
		"D2FE28",
	}

	run(rows)
}

func TestHexToBinary(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{
			in:   "1",
			want: "0001",
		},
		{
			in:   "F",
			want: "1111",
		},
		{
			in:   "F1",
			want: "11110001",
		},
		{
			in:   "1F",
			want: "00011111",
		},
	}

	for _, tc := range tests {
		t.Run(tc.in, func(t *testing.T) {
			got := HexToBinary(tc.in)
			require.Equal(t, tc.want, got)
		})
	}
}

func TestBinaryToInt(t *testing.T) {
	tests := []struct {
		bits string
		want int
	}{
		{
			bits: "110",
			want: 6,
		},
		{
			bits: "100",
			want: 4,
		},
	}

	for _, tc := range tests {
		t.Run(tc.bits, func(t *testing.T) {
			got := BinaryToInt(tc.bits)
			require.Equal(t, tc.want, got)
		})
	}
}

func TestDecodePacket(t *testing.T) {
	tests := []struct {
		name string
		bits string
		want Packet
	}{
		{
			name: "decode 1",
			bits: "110100101111111000101000",
			want: &ValuePacket{
				header: &Header{
					Version: 6,
					Type:    4,
				},
				Value: 2021,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := DecodePacket(tc.bits)
			require.Equal(t, tc.want, got)
		})
	}
}

func TestDecodeValue(t *testing.T) {
	tests := []struct {
		value string
		want  int
	}{
		{
			value: "101111111000101000",
			want:  2021,
		},
	}

	for _, tc := range tests {
		t.Run(tc.value, func(t *testing.T) {
			got := DecodeValue(tc.value)
			require.Equal(t, tc.want, got)
		})
	}
}

func TestChunkString(t *testing.T) {
	tests := []struct {
		value string
		size  int
		want  []string
	}{
		{
			value: "01011",
			size:  2,
			want:  []string{"01", "01", "1"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.value, func(t *testing.T) {
			got := ChunkString(tc.value, tc.size)
			require.Equal(t, tc.want, got)
		})
	}
}
