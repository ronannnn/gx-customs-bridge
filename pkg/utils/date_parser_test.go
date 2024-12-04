package utils_test

import (
	"testing"
	"time"

	"github.com/ronannnn/gx-customs-bridge/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestCorrectParseHdeTime(t *testing.T) {
	testDate := "2000-08-01 14:53:02"
	parsedTime := utils.ParseHdeApprResultTime(testDate)
	require.Equal(t, 2000, parsedTime.Year())
}

func TestWrongParseHdeTime(t *testing.T) {
	testDate := "2000-08-01-14:53:02"
	parsedTime := utils.ParseHdeApprResultTime(testDate)
	require.Equal(t, time.Now().Year(), parsedTime.Year())
}

func TestCorrectParseClientGeneratedTime(t *testing.T) {
	testDate := "2000101908250905786339"
	parsedTime := utils.ParseClientGeneratedTime(testDate)
	require.Equal(t, 2000, parsedTime.Year())
}

func TestWrongParseClientGeneratedTime(t *testing.T) {
	testDate := "20001019082"
	parsedTime := utils.ParseClientGeneratedTime(testDate)
	require.Equal(t, time.Now().Year(), parsedTime.Year())
}
