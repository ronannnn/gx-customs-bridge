package dec

import (
	"encoding/xml"

	"github.com/ronannnn/gx-customs-bridge/internal"
	"github.com/ronannnn/gx-customs-bridge/pkg/customs/decmodels"
)

type DecXmlService interface {
	GenDecTmpXml(decTmp decmodels.Dec, operType string) ([]byte, error)
}

func ProvideDecXmlService(
	customsCfg *internal.CustomsCfg,
) DecXmlService {
	return &DecXmlServiceImpl{
		customsCfg: customsCfg,
	}
}

type DecXmlServiceImpl struct {
	customsCfg *internal.CustomsCfg
}

func (s *DecXmlServiceImpl) GenDecTmpXml(decModel decmodels.Dec, operType string) (xmlBytes []byte, err error) {
	// validation
	if err = decmodels.CheckIfDecOperTypeValid(operType); err != nil {
		return
	}
	decModel.DecHead.InputerName = &s.customsCfg.DclErConc
	decModel.DecHead.DeclareName = &s.customsCfg.DclErConc
	decModel.DecHead.TypistNo = &s.customsCfg.IcCardNo
	decModel.DecSign.ICCode = &s.customsCfg.IcCardNo
	decModel.DecSign.OperName = &s.customsCfg.DclErConc
	convertedOperType := decmodels.DecOperType(operType)
	decModel.DecSign.OperType = &convertedOperType
	// 生成decTmpXml
	decTmpXml := decmodels.DecXml{}
	decTmpXml.Version = "4.8"
	decTmpXml.Xmlns = "http://www.chinaport.gov.cn/dec"
	decTmpXml.DecHead = decModel.DecHead
	decTmpXml.DecLists.DecLists = decModel.DecLists
	decTmpXml.DecContainers.DecContainers = decModel.DecContainers
	decTmpXml.DecFreeTxt = decModel.DecFreeTxt
	decTmpXml.DecSign = decModel.DecSign
	// 保存xml
	var xmlBodyBytes []byte
	if xmlBodyBytes, err = xml.Marshal(decTmpXml); err != nil {
		return
	}
	xmlBytes = []byte(xml.Header + string(xmlBodyBytes))
	return
}
