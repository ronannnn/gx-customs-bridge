package sas_test

import (
	"io"
	"os"
	"testing"

	"github.com/ronannnn/gx-customs-bridge/internal"
	"github.com/ronannnn/gx-customs-bridge/internal/services/customs/inv"
	"github.com/ronannnn/gx-customs-bridge/internal/services/customs/sas"
	"github.com/stretchr/testify/require"
)

func TestGenXml(t *testing.T) {
	service := inv.ProvideInvService(internal.CustomsCfg{
		Inv101SysId:     "Z7",
		OperCusRegCode:  "1234567890",
		IcCardNo:        "11111111111111111111",
		Sas121DclErConc: "Nobody",
	})
	impexpMarkcd := "I"
	service.GenXml(inv.Inv101{
		Head: inv.Inv101Head{
			ImpexpMarkcd: impexpMarkcd,
		},
	}, "1")
}

func TestParseInv221(t *testing.T) {
	service := sas.ProvideSasService(internal.CustomsCfg{
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
	service := sas.ProvideSasService(internal.CustomsCfg{
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
	service := sas.ProvideSasService(internal.CustomsCfg{
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
