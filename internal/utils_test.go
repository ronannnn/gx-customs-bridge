package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetFilenamePrefix(t *testing.T) {
	filename := "test.txt"
	want := "test"
	got := GetFilenamePrefix(filename)
	require.Equal(t, want, got)
}
