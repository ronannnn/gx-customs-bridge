package sasmodels

import (
	"encoding/xml"

	"github.com/ronannnn/gx-customs-bridge/pkg/customs/commonmodels"
)

type Sas223 struct {
	commonmodels.MqDataCommonPayload
	HdeApprResult commonmodels.HdeApprResult `json:"hdeApprResult"`
	// 和SAS221共用Head
	Head Sas221Head `json:"head"`
}

type Sas223Xml struct {
	XMLName     xml.Name `xml:"Package"`
	EnvelopInfo commonmodels.EnvelopInfo
	DataInfo    struct {
		PocketInfo   commonmodels.PocketInfo
		BusinessData struct {
			Sas223 struct {
				HdeApprResult  commonmodels.HdeApprResult
				SasPassportBsc Sas221Head
			} `xml:"SAS223"`
		} `xml:"BussinessData"`
	}
}
