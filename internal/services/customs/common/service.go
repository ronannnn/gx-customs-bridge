package common

import (
	"encoding/xml"

	"github.com/ronannnn/gx-customs-bridge/internal"
)

type CustomsCommonXmlService interface {
	ParseCommonResponseMessageXml([]byte) (CommonResponeMessage, error)
	ParseReceiptMessageHeader([]byte) (ReceiptMessageHeader, error)
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

func (s *CustomsCommonXmlServiceImpl) ParseCommonResponseMessageXml(xmlBytes []byte) (crm CommonResponeMessage, err error) {
	err = xml.Unmarshal(xmlBytes, &crm)
	return
}

func (s *CustomsCommonXmlServiceImpl) ParseReceiptMessageHeader(xmlBytes []byte) (rmh ReceiptMessageHeader, err error) {
	err = xml.Unmarshal(xmlBytes, &rmh)
	return
}
