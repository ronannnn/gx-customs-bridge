package dec_test

import (
	"testing"

	"github.com/ronannnn/gx-customs-bridge/internal/services/customs/dec"
	"github.com/stretchr/testify/require"
)

func TestDecFilenameParts(t *testing.T) {
	t.Run("NewFilenamePrefix", func(t *testing.T) {
		impexpMarkcdI := "I"
		clientSeqNo := "123456"
		filenameParts, err := dec.NewDecFilenameParts(
			&impexpMarkcdI,
			&clientSeqNo,
		)
		require.NoError(t, err)
		require.Equal(t, "I_123456_t-1", filenameParts.GenOutBoxFilenamePrefix())
		require.Equal(t, "I_123456_t-1.xml", filenameParts.GenOutBoxFilename("xml"))
	})

	t.Run("ParseDecFilenameParts", func(t *testing.T) {
		filename := "Successed_I_12121_t-2_202408191442314383412"
		decFilenameParts, err := dec.ParseDecFilename(filename)
		require.NoError(t, err)
		require.Equal(t, dec.SuccessedOrFailedTypeSuccessed, decFilenameParts.SuccessedOrFailed)
		require.Equal(t, dec.ImpexpMarkcdI, decFilenameParts.ImpexpMarkcd)
		require.Equal(t, "12121", decFilenameParts.ClientSeqNo)
		require.Equal(t, 2, decFilenameParts.RetryTimes)
		require.Equal(t, "202408191442314383412", decFilenameParts.Timestamp)
	})
}
