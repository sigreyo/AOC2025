package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuildLine(t *testing.T) {

	tests := []struct {
		start      int
		spins      int
		isLeftspin bool
		want       []int
	}{
		{
			start:      50,
			spins:      5,
			isLeftspin: true,
			want:       []int{49, 48, 47, 46, 45},
		},
		{
			start:      98,
			spins:      3,
			isLeftspin: false,
			want:       []int{99, 0, 1},
		},
		{
			start:      -10,
			spins:      3,
			isLeftspin: false,
			want:       []int{-9, -8, -7},
		},
	}
	for _, tc := range tests {
		got := buildLine(tc.start, tc.spins, tc.isLeftspin)
		require.Equal(t, tc.want, got)
	}
}

func TestCheckZeroes(t *testing.T) {
	want := 2
	got := checkZeroes([]int{99, 0, 100})
	require.Equal(t, want, got)
}
