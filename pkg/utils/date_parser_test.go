package utils_test

import (
	"testing"
	"time"

	"github.com/ronannnn/gx-customs-bridge/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestCorrectParseHdeDate(t *testing.T) {
	testDate := "2000-08-01 14:53:02"
	date := utils.ParseHdeApprResultDate(testDate)
	require.Equal(t, 2000, date.Year())
}

func TestWrongParseHdeDate(t *testing.T) {
	testDate := "2000-08-01-14:53:02"
	date := utils.ParseHdeApprResultDate(testDate)
	require.Equal(t, time.Now().Year(), date.Year())
}
