package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSplitAndCheckForRepeats(t *testing.T) {
	tests := []struct {
		fullWord string
		want     int
	}{
		{
			fullWord: "987654321111111",
			want:     98,
		},
		{
			fullWord: "811111111111119",
			want:     89,
		},
		{
			fullWord: "234234234234278",
			want:     78,
		},
		{
			fullWord: "818181911112111",
			want:     92,
		},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			require.Equal(t, tc.want, collectJolts(tc.fullWord))

		})
	}
}
