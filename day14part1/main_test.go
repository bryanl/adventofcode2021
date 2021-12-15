package main

import (
	"testing"

	"github.com/bryan/adventofcode2021/internal/support"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	rows := []string{"NNCB", "CH -> B"}

	data := Parse(rows)

	want := &Data{
		Template: "NNCB",
		Rules: map[string]string{
			"CH": "B",
		},
	}

	require.Equal(t, want, data)
}

func TestData_Pairs(t *testing.T) {
	rows := support.ReadFromDisk(t, "sample.txt")
	data := Parse(rows)
	got := data.Pairs()

	want := []string{"NN", "NC", "CB"}
	require.Equal(t, want, got)
}

func TestData_Perform(t *testing.T) {
	rows := support.ReadFromDisk(t, "sample.txt")

	tests := []struct {
		template string
		want     string
	}{
		{
			template: "NNCB",
			want:     "NCNBCHB",
		},
		{
			template: "NCNBCHB",
			want:     "NBCCNBBBCBHCB",
		},
		{
			template: "NBCCNBBBCBHCB",
			want:     "NBBBCNCCNBBNBNBBCHBHHBCHB",
		},
		{
			template: "NBBBCNCCNBBNBNBBCHBHHBCHB",
			want:     "NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB",
		},
	}

	for _, tc := range tests {
		t.Run(tc.template, func(t *testing.T) {
			data := Parse(rows)
			data.Template = tc.template
			data.Perform()

			got := data.Template
			require.Equal(t, tc.want, got)
		})
	}

}
