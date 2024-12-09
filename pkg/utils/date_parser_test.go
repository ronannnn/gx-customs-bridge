package utils_test

import (
	"testing"
	"time"

	"github.com/ronannnn/gx-customs-bridge/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestCorrectParseHdeTime(t *testing.T) {
	time.Local = time.UTC
	testDate := "2000-08-01 14:53:02"
	parsedTime := utils.ParseHdeApprResultTime(testDate)
	require.Equal(t, 2000, parsedTime.Year())
	require.Equal(t, 8, int(parsedTime.Month()))
	require.Equal(t, 1, parsedTime.Day())
	require.Equal(t, 6, parsedTime.Hour())
	require.Equal(t, 53, parsedTime.Minute())
	require.Equal(t, 2, parsedTime.Second())
}

func TestWrongParseHdeTime(t *testing.T) {
	testDate := "2000-08-01-14:53:02"
	parsedTime := utils.ParseHdeApprResultTime(testDate)
	require.Equal(t, time.Now().Year(), parsedTime.Year())
}

func TestCorrectParseClientGeneratedTime(t *testing.T) {
	time.Local = time.UTC
	testDate := "2000101908250905786339"
	parsedTime := utils.ParseClientGeneratedTime(testDate)
	require.Equal(t, 2000, parsedTime.Year())
	require.Equal(t, 10, int(parsedTime.Month()))
	require.Equal(t, 19, parsedTime.Day())
	require.Equal(t, 0, parsedTime.Hour())
	require.Equal(t, 25, parsedTime.Minute())
	require.Equal(t, 9, parsedTime.Second())
}

func TestWrongParseClientGeneratedTime(t *testing.T) {
	testDate := "20001019082"
	parsedTime := utils.ParseClientGeneratedTime(testDate)
	require.Equal(t, time.Now().Year(), parsedTime.Year())
}
