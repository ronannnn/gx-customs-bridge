package sas_test

import (
	"io"
	"os"
	"testing"

	"github.com/ronannnn/gx-customs-bridge/internal"
	"github.com/ronannnn/gx-customs-bridge/internal/services/customs/sas"
	"github.com/stretchr/testify/require"
)

func TestGenInv101Xml(t *testing.T) {
	service := sas.ProvideSasXmlService(&internal.CustomsCfg{
		Inv101SysId:     "Z7",
		OperCusRegCode:  "1234567890",
		IcCardNo:        "11111111111111111111",
		Sas121DclErConc: "Nobody",
	})
	impexpMarkcd := "I"
	service.GenInv101Xml(sas.Inv101{
		Head: sas.Inv101Head{
			ImpexpMarkcd: &impexpMarkcd,
		},
	}, "1")
}

func TestGenSas121Xml(t *testing.T) {
	service := sas.ProvideSasXmlService(&internal.CustomsCfg{
		Inv101SysId:     "Z7",
		OperCusRegCode:  "1234567890",
		IcCardNo:        "11111111111111111111",
		Sas121DclErConc: "Nobody",
	})
	impexpMarkcd := "I"
	service.GenInv101Xml(sas.Inv101{
		Head: sas.Inv101Head{
			ImpexpMarkcd: &impexpMarkcd,
		},
	}, "1")
}

func TestParseInv201(t *testing.T) {
	service := sas.ProvideSasXmlService(&internal.CustomsCfg{
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
	var inv201 sas.Inv201
	inv201, err = service.ParseInv201Xml(xmlBytes)
	require.NoError(t, err)

	// validate
	require.Equal(t, "202400000122717119", inv201.HdeApprResult.EtpsPreentNo)
	require.Equal(t, "QD311124E000098849", inv201.Head.BondInvtNo)
}

func TestParseInv202(t *testing.T) {
	service := sas.ProvideSasXmlService(&internal.CustomsCfg{
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
	var inv202 sas.Inv202
	inv202, err = service.ParseInv202Xml(xmlBytes)
	require.NoError(t, err)

	// validate
	require.Equal(t, "202400000122534496", inv202.InvApprResult.InvPreentNo)
	require.Equal(t, "QD311124E000097307", inv202.InvApprResult.BusinessId)
}

func TestParseInv211(t *testing.T) {
	service := sas.ProvideSasXmlService(&internal.CustomsCfg{
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
	var inv211 sas.Inv211
	inv211, err = service.ParseInv211Xml(xmlBytes)
	require.NoError(t, err)

	// validate
	require.Equal(t, "QD311124I000067486", inv211.Head.BondInvtNo)
	require.Equal(t, "T3111W000076", inv211.List[0].BwsNo)
	require.Equal(t, "2255", inv211.List[0].GdsSeqno)
}

func TestParseInv221(t *testing.T) {
	service := sas.ProvideSasXmlService(&internal.CustomsCfg{
		Inv101SysId:     "Z7",
		OperCusRegCode:  "1234567890",
		IcCardNo:        "11111111111111111111",
		Sas121DclErConc: "Nobody",
	})

	// read file
	file, err := os.Open("test/SAS221.xml")
	require.NoError(t, err)
	defer file.Close()

	// convert to bytes
	var xmlBytes []byte
	xmlBytes, err = io.ReadAll(file)
	require.NoError(t, err)

	// parse to sas221
	var sas221 sas.Sas221
	sas221, err = service.ParseSas221Xml(xmlBytes)
	require.NoError(t, err)

	// validate
	require.Equal(t, "202400000069440319", sas221.HdeApprResult.EtpsPreentNo)
	require.Equal(t, "Z3111I240728000000000003", sas221.Head.PassportNo)
	require.Equal(t, "9505100090", sas221.List[0].GdsMtno)
	require.Equal(t, "QD311124I000067647", sas221.Acmp[0].RltNo)
}

func TestParseInv223(t *testing.T) {
	service := sas.ProvideSasXmlService(&internal.CustomsCfg{
		Inv101SysId:     "Z7",
		OperCusRegCode:  "1234567890",
		IcCardNo:        "11111111111111111111",
		Sas121DclErConc: "Nobody",
	})

	// read file
	file, err := os.Open("test/SAS223.xml")
	require.NoError(t, err)
	defer file.Close()

	// convert to bytes
	var xmlBytes []byte
	xmlBytes, err = io.ReadAll(file)
	require.NoError(t, err)

	// parse to sas223
	var sas223 sas.Sas223
	sas223, err = service.ParseSas223Xml(xmlBytes)
	require.NoError(t, err)

	// validate
	require.Equal(t, "202400000069405902", sas223.HdeApprResult.EtpsPreentNo)
	require.Equal(t, "Z3111E240726000000000456", sas223.Head.PassportNo)
}

func TestParseInv224(t *testing.T) {
	service := sas.ProvideSasXmlService(&internal.CustomsCfg{
		Inv101SysId:     "Z7",
		OperCusRegCode:  "1234567890",
		IcCardNo:        "11111111111111111111",
		Sas121DclErConc: "Nobody",
	})

	// read file
	file, err := os.Open("test/SAS224.xml")
	require.NoError(t, err)
	defer file.Close()

	// convert to bytes
	var xmlBytes []byte
	xmlBytes, err = io.ReadAll(file)
	require.NoError(t, err)

	// parse to sas224
	var sas224 sas.Sas224
	sas224, err = service.ParseSas224Xml(xmlBytes)
	require.NoError(t, err)

	// validate
	require.Equal(t, "202400000069449716", sas224.HdeApprResult.EtpsPreentNo)
	require.Equal(t, "Z3111I240728000000000111", sas224.Head.PassportNo)
}
