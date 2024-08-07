package inv_test

import (
	"io"
	"os"
	"testing"

	"github.com/ronannnn/gx-customs-bridge/internal"
	"github.com/ronannnn/gx-customs-bridge/internal/services/customs/inv"
	"github.com/stretchr/testify/require"
)

func TestGenXml(t *testing.T) {
	service := inv.ProvideInvXmlService(&internal.CustomsCfg{
		Inv101SysId:     "Z7",
		OperCusRegCode:  "1234567890",
		IcCardNo:        "11111111111111111111",
		Sas121DclErConc: "Nobody",
	})
	impexpMarkcd := "I"
	service.GenInv101Xml(inv.Inv101{
		Head: inv.Inv101Head{
			ImpexpMarkcd: impexpMarkcd,
		},
	}, "1")
}

func TestParseInv201(t *testing.T) {
	service := inv.ProvideInvXmlService(&internal.CustomsCfg{
		Inv101SysId:     "Z7",
		OperCusRegCode:  "1234567890",
		IcCardNo:        "11111111111111111111",
		Sas121DclErConc: "Nobody",
	})

	// read file
	file, err := os.Open("test/INV201.xml")
	require.NoError(t, err)
	defer file.Close()

	// convert to bytes
	var xmlBytes []byte
	xmlBytes, err = io.ReadAll(file)
	require.NoError(t, err)

	// parse to inv201
	var inv201 inv.Inv201
	inv201, err = service.ParseInv201Xml(xmlBytes)
	require.NoError(t, err)

	// validate
	require.Equal(t, "202400000122717119", inv201.HdeApprResult.EtpsPreentNo)
	require.Equal(t, "QD311124E000098849", inv201.Head.BondInvtNo)
}

func TestParseInv202(t *testing.T) {
	service := inv.ProvideInvXmlService(&internal.CustomsCfg{
		Inv101SysId:     "Z7",
		OperCusRegCode:  "1234567890",
		IcCardNo:        "11111111111111111111",
		Sas121DclErConc: "Nobody",
	})

	// read file
	file, err := os.Open("test/INV202.xml")
	require.NoError(t, err)
	defer file.Close()

	// convert to bytes
	var xmlBytes []byte
	xmlBytes, err = io.ReadAll(file)
	require.NoError(t, err)

	// parse to inv202
	var inv202 inv.Inv202
	inv202, err = service.ParseInv202Xml(xmlBytes)
	require.NoError(t, err)

	// validate
	require.Equal(t, "202400000122534496", inv202.InvApprResult.InvPreentNo)
	require.Equal(t, "QD311124E000097307", inv202.InvApprResult.BusinessId)
}

func TestParseInv211(t *testing.T) {
	service := inv.ProvideInvXmlService(&internal.CustomsCfg{
		Inv101SysId:     "Z7",
		OperCusRegCode:  "1234567890",
		IcCardNo:        "11111111111111111111",
		Sas121DclErConc: "Nobody",
	})

	// read file
	file, err := os.Open("test/INV211.xml")
	require.NoError(t, err)
	defer file.Close()

	// convert to bytes
	var xmlBytes []byte
	xmlBytes, err = io.ReadAll(file)
	require.NoError(t, err)

	// parse to inv211
	var inv211 inv.Inv211
	inv211, err = service.ParseInv211Xml(xmlBytes)
	require.NoError(t, err)

	// validate
	require.Equal(t, "QD311124I000067486", inv211.Head.BondInvtNo)
	require.Equal(t, "T3111W000076", inv211.List[0].BwsNo)
	require.Equal(t, "2255", inv211.List[0].GdsSeqno)
}
