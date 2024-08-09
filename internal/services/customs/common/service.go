package common

import (
	"encoding/xml"

	"github.com/ronannnn/gx-customs-bridge/internal"
	"github.com/ronannnn/gx-customs-bridge/pkg/customs/commonmodels"
)

type CustomsCommonXmlService interface {
	ParseCommonResponseMessageXml([]byte) (commonmodels.CommonResponeMessage, error)
	ParseReceiptMessageHeader([]byte) (commonmodels.ReceiptMessageHeader, error)
}

func ProvideCustomsCommonXmlService(
	customsCfg *internal.CustomsCfg,
) CustomsCommonXmlService {
	return &CustomsCommonXmlServiceImpl{
		customsCfg: customsCfg,
	}
}

type CustomsCommonXmlServiceImpl struct {
	customsCfg *internal.CustomsCfg
}

func (s *CustomsCommonXmlServiceImpl) ParseCommonResponseMessageXml(xmlBytes []byte) (crm commonmodels.CommonResponeMessage, err error) {
	err = xml.Unmarshal(xmlBytes, &crm)
	return
}

func (s *CustomsCommonXmlServiceImpl) ParseReceiptMessageHeader(xmlBytes []byte) (rmh commonmodels.ReceiptMessageHeader, err error) {
	err = xml.Unmarshal(xmlBytes, &rmh)
	return
}
