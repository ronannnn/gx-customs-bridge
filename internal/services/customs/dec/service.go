package dec

import (
	"encoding/xml"
	"fmt"

	"github.com/ronannnn/gx-customs-bridge/internal"
	"github.com/ronannnn/gx-customs-bridge/pkg/customs/decmodels"
)

type DecXmlService interface {
	GenDecTmpXml(decTmp decmodels.Dec, operType string, companyType string) ([]byte, error)
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

func (s *DecXmlServiceImpl) GenDecTmpXml(decModel decmodels.Dec, operType string, companyType string) (xmlBytes []byte, err error) {
	// 替换部分数据
	icCard, ok := s.customsCfg.IcCardMap[companyType]
	if !ok {
		err = fmt.Errorf("公司类型(%s)不存在", companyType)
		return
	}
	decModel.DecHead.InputerName = &icCard.DclErConc
	decModel.DecHead.DeclareName = &icCard.DclErConc
	decModel.DecHead.TypistNo = &icCard.IcCardNo
	decModel.DecSign.ICCode = &icCard.IcCardNo
	decModel.DecSign.OperName = &icCard.DclErConc
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
