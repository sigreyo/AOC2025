package main

import (
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
	want1 := 123123
	got1 := checkDuplicateOccurences(123123)
	require.Equal(t, want1, got1)

	want2 := 0
	got2 := checkDuplicateOccurences(123124)
	require.Equal(t, want2, got2)
}
