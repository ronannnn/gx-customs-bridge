package sas_test

import (
	"testing"

	"github.com/ronannnn/gx-customs-bridge/internal/services/customs/sas"
	"github.com/stretchr/testify/require"
)

func TestSasFilenameParts(t *testing.T) {
	t.Run("NewFilenamePrefix", func(t *testing.T) {
		impexpMarkcdI := "I"
		etpsInnerInvtNo := "123456"
		filenameParts, err := sas.NewSasFilenameParts(
			sas.UploadTypeInv101,
			&impexpMarkcdI,
			&etpsInnerInvtNo,
		)
		require.NoError(t, err)
		require.Equal(t, "INV101_I_123456_t-1", filenameParts.GenOutBoxFilenamePrefix())
		require.Equal(t, "INV101_I_123456_t-1.xml", filenameParts.GenOutBoxFilename("xml"))
	})

	t.Run("ParseSasFilenameParts", func(t *testing.T) {
		filename := "Successed_SAS121_I_浙B5Z502-1_t-2_202408191442314383412"
		sasFilenameParts, err := sas.ParseSasFilename(filename)
		require.NoError(t, err)
		require.Equal(t, sas.SuccessedOrFailedTypeSuccessed, sasFilenameParts.SuccessedOrFailed)
		require.Equal(t, sas.UploadTypeSas121, sasFilenameParts.UploadType)
		require.Equal(t, sas.ImpexpMarkcdI, sasFilenameParts.ImpexpMarkcd)
		require.Equal(t, "浙B5Z502-1", sasFilenameParts.EtpsInnerInvtNo)
		require.Equal(t, 2, sasFilenameParts.RetryTimes)
		require.Equal(t, "202408191442314383412", sasFilenameParts.Timestamp)
	})
}
