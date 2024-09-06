package dec

import (
	"encoding/xml"

	"github.com/ronannnn/gx-customs-bridge/internal"
	"github.com/ronannnn/gx-customs-bridge/pkg/customs/decmodels"
)

type DecXmlService interface {
	GenDecTmpXml(decTmp decmodels.DecTmp) ([]byte, error)
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

func (s *DecXmlServiceImpl) GenDecTmpXml(decTmp decmodels.DecTmp) (xmlBytes []byte, err error) {
	decTmp.DecHead.InputerName = s.customsCfg.DclErConc
	decTmp.DecHead.DeclareName = s.customsCfg.DclErConc
	decTmp.DecHead.TypistNo = s.customsCfg.IcCardNo
	decTmp.DecSign.ICCode = s.customsCfg.IcCardNo
	decTmp.DecSign.OperName = s.customsCfg.DclErConc
	// 生成decTmpXml
	decTmpXml := decmodels.DecTmpXml{}
	decTmpXml.Version = "4.8"
	decTmpXml.Xmlns = "http://www.chinaport.gov.cn/dec"
	decTmpXml.DecHead = decTmp.DecHead
	decTmpXml.DecLists.DecLists = decTmp.DecLists
	decTmpXml.DecContainers.DecContainers = decTmp.DecContainers
	decTmpXml.DecFreeTxt = decTmp.DecFreeTxt
	decTmpXml.DecSign = decTmp.DecSign
	// decTmpXml.DecLicenseDocus.DecLicenseDocus = decTmp.DecLicenseDocus
	// decTmpXml.EdocRealations = decTmp.DecEdocRealations
	// 保存xml
	var xmlBodyBytes []byte
	if xmlBodyBytes, err = xml.Marshal(decTmpXml); err != nil {
		return
	}
	xmlBytes = []byte(xml.Header + string(xmlBodyBytes))
	return
}
