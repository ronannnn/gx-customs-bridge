package inv

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/ronannnn/gx-customs-bridge/internal"
)

type InvService interface {
	GenXml(inv101 Inv101, declareFlag string) ([]byte, error)
	ParseInv201Xml([]byte) (Inv201, error)
	ParseInv202Xml([]byte) (Inv202, error)
	ParseInv211Xml([]byte) (Inv211, error)
}

func ProvideInvService(customsCfg internal.CustomsCfg) InvService {
	return &InvServiceImpl{
		customsCfg: customsCfg,
	}
}

type InvServiceImpl struct {
	customsCfg internal.CustomsCfg
}

func (s *InvServiceImpl) GenXml(inv101 Inv101, declareFlag string) (xmlBytes []byte, err error) {
	// 校验
	if declareFlag != "0" && declareFlag != "1" {
		err = fmt.Errorf("申报标志(declareFlag)必须是0或1")
		return
	}
	// 替换部分数据
	inv101.Head.IcCardNo = s.customsCfg.IcCardNo
	// 生成inv101Xml
	inv101Xml := Inv101Xml{}
	inv101Xml.Object.Package.EnvelopInfo.MessageType = "INV101"
	inv101Xml.Object.Package.DataInfo.BusinessData.DeclareFlag = declareFlag
	inv101Xml.Object.Package.DataInfo.BusinessData.InvtMessage.SysId = s.customsCfg.Inv101SysId
	inv101Xml.Object.Package.DataInfo.BusinessData.InvtMessage.OperCusRegCode = s.customsCfg.OperCusRegCode
	inv101Xml.Object.Package.DataInfo.BusinessData.InvtMessage.InvtHeadType = inv101.Head
	inv101Xml.Object.Package.DataInfo.BusinessData.InvtMessage.InvtListType = inv101.List
	// 保存xml
	var xmlBodyBytes []byte
	if xmlBodyBytes, err = xml.Marshal(inv101Xml); err != nil {
		return
	}
	xmlBytes = []byte(xml.Header + string(xmlBodyBytes))
	os.Stdout.Write(xmlBytes)
	return
}

func (s *InvServiceImpl) ParseInv201Xml(xmlBytes []byte) (inv201 Inv201, err error) {
	inv201Xml := Inv201Xml{}
	if err = xml.Unmarshal(xmlBytes, &inv201Xml); err != nil {
		return
	}

	inv201.HdeApprResult = inv201Xml.DataInfo.BusinessData.Inv201.HdeApprResult
	inv201.Head = inv201Xml.DataInfo.BusinessData.Inv201.BondInvtBsc
	inv201.List = inv201Xml.DataInfo.BusinessData.Inv201.BondInvtDtl
	return
}

func (s *InvServiceImpl) ParseInv202Xml(xmlBytes []byte) (inv202 Inv202, err error) {
	inv202Xml := Inv202Xml{}
	if err = xml.Unmarshal(xmlBytes, &inv202Xml); err != nil {
		return
	}
	inv202.InvApprResult = inv202Xml.DataInfo.BusinessData.Inv202.InvApprResult
	return
}

func (s *InvServiceImpl) ParseInv211Xml(xmlBytes []byte) (inv211 Inv211, err error) {
	inv211Xml := Inv211Xml{}
	if err = xml.Unmarshal(xmlBytes, &inv211Xml); err != nil {
		return
	}
	inv211.Head = inv211Xml.DataInfo.BusinessData.Inv211.BondInvtBsc
	inv211.List = inv211Xml.DataInfo.BusinessData.Inv211.BwsDt
	return
}
