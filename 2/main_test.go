package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetRange(t *testing.T) {
	want := []int{3, 4, 5, 6}
	got := getRange(3, 6)
	require.Equal(t, want, got)
}

func TestSplitToStartAndEnd(t *testing.T) {
	want1, want2 := 10, 13
	got1, got2 := splitToStartAndEnd("10-13")
	require.Equal(t, want1, got1)
	require.Equal(t, want2, got2)
}

func TestCheckDuplicateOccurences(t *testing.T) {
	tests := []struct {
		input    int
		want     bool
		intRange []int
	}{
		{
			input: 11,
			want:  true,
		},
		{
			input: 111,
			want:  true,
		},
		{
			input: 123123123,
			want:  true,
		},
		{
			input: 446446,
			want:  true,
		},
		{
			input: 38593859,
			want:  true,
		},
		{
			input: 2121212121,
			want:  true,
		},
		{
			input: 1188511885,
			want:  true,
		},
		{
			input: 1188511885,
			want:  true,
		},
		{
			input:    0,
			want:     false,
			intRange: getRange(1698522, 1698528),
		},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			if tc.input > 0 {
				require.Equal(t, tc.want, hasDuplicateOccurences(tc.input))
			}

			if tc.intRange != nil {
				for _, v := range tc.intRange {
					require.Equal(t, tc.want, hasDuplicateOccurences(v))
				}
			}
		})
	}
}

func TestSplitAndCheckForRepeats(t *testing.T) {
	tests := []struct {
		fullWord string
		want     bool
	}{
		{
			fullWord: "824824824",
			want:     true,
		},
		{
			fullWord: "2121212121",
			want:     true,
		},
		{
			fullWord: "565657",
			want:     false,
		},
		{
			fullWord: "446444",
			want:     false,
		},
		{
			fullWord: "1",
			want:     false,
		},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			require.Equal(t, tc.want, splitAndCheckForRepeats(tc.fullWord))
		})
	}
}
