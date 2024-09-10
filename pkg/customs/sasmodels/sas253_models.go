package sasmodels

import (
	"encoding/xml"

	"github.com/ronannnn/gx-customs-bridge/pkg/customs/commonmodels"
)

// 两步申报核放单过卡回执

type Sas253 struct {
	HdeApprResult commonmodels.HdeApprResult `json:"hdeApprResult"`
	Head          Sas251Head                 `json:"head"` // 和251的结构一样
}

type Sas253Xml struct {
	XMLName     xml.Name `xml:"Package"`
	EnvelopInfo commonmodels.EnvelopInfo
	DataInfo    struct {
		PocketInfo   commonmodels.PocketInfo
		BusinessData struct {
			Sas253 struct {
				HdeApprResult       commonmodels.HdeApprResult
				Sas2stepPassportBsc Sas251Head // 和251的结构一样
			} `xml:"SAS253"`
		} `xml:"BussinessData"`
	}
}
