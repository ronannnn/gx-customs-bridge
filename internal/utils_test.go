package internal

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetFilenamePrefix(t *testing.T) {
	filename := "test.txt"
	want := "test"
	got := GetFilenamePrefix(filename)
	require.Equal(t, want, got)
}

func TestGetFilenameSuffix(t *testing.T) {
	filename := "test.txt"
	want := "txt"
	got := GetFilenameSuffix(filename)
	require.Equal(t, want, got)
}

func TestGetFilenameWithoutDot(t *testing.T) {
	filename := "test"
	wantPrefix := "test"
	gotPrefix := GetFilenameSuffix(filename)
	require.Equal(t, wantPrefix, gotPrefix)

	wantSuffix := "test"
	gotSuffix := GetFilenameSuffix(filename)
	require.Equal(t, wantSuffix, gotSuffix)
}

func TestGetFilenameFromPath(t *testing.T) {
	filePath := filepath.Join("ok", "test", "test.txt")
	want := "test.txt"
	got := GetFilenameFromPath(filePath)
	require.Equal(t, want, got)
}

func TestGetFilenameFromPathWithoutSlash(t *testing.T) {
	filePath := "test.txt"
	want := "test.txt"
	got := GetFilenameFromPath(filePath)
	require.Equal(t, want, got)
}
