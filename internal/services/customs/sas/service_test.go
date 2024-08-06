package sas_test

import (
	"testing"

	"github.com/ronannnn/gx-customs-bridge/internal"
	"github.com/ronannnn/gx-customs-bridge/internal/services/customs/inv"
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
			ImpexpMarkcd: &impexpMarkcd,
		},
	}, "1")
}
