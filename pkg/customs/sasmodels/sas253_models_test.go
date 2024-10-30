package sasmodels_test

import (
	"encoding/xml"
	"os"
	"testing"

	"github.com/ronannnn/gx-customs-bridge/pkg/customs/sasmodels"
	"github.com/stretchr/testify/require"
)

func TestSas253Parse(t *testing.T) {
	xmlFileBytes, err := os.ReadFile("./testdata/sas253.xml")
	require.NoError(t, err)

	var sas253 sasmodels.Sas253Xml
	if err = xml.Unmarshal(xmlFileBytes, &sas253); err != nil {
		return
	}
	require.NoError(t, err)
	require.Equal(t, "2024-10-30 15:40:55", sas253.DataInfo.BusinessData.Sas253.Sas2stepPassportBsc.PassTime)
}
