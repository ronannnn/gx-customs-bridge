package sas

import (
	"encoding/xml"

	"github.com/ronannnn/gx-customs-bridge/internal/services/customs/common"
)

type Sas224 struct {
	HdeApprResult common.HdeApprResult `json:"hdeApprResult"`
	// 和SAS221共用Head
	Head Sas221Head `json:"head"`
}

type Sas224Xml struct {
	XMLName     xml.Name `xml:"Package"`
	EnvelopInfo common.EnvelopInfo
	DataInfo    struct {
		PocketInfo   common.PocketInfo
		BusinessData struct {
			Sas224 struct {
				HdeApprResult  common.HdeApprResult
				SasPassportBsc Sas221Head
			} `xml:"SAS224"`
		} `xml:"BussinessData"`
	}
}
