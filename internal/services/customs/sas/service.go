package sas

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/ronannnn/gx-customs-bridge/internal"
)

type SasXmlService interface {
	GenInv101Xml(inv101 Inv101, declareFlag string) ([]byte, error)
	GenSas121Xml(sas121 Sas121, declareFlag string) ([]byte, error)

	ParseInv201Xml([]byte) (Inv201, error)
	ParseInv202Xml([]byte) (Inv202, error)
	ParseInv211Xml([]byte) (Inv211, error)
	ParseSas221Xml([]byte) (Sas221, error)
	ParseSas223Xml([]byte) (Sas223, error)
	ParseSas224Xml([]byte) (Sas224, error)
}

func ProvideSasXmlService(
	customsCfg *internal.CustomsCfg,
) SasXmlService {
	return &SasXmlServiceImpl{
		customsCfg: customsCfg,
	}
}

type SasXmlServiceImpl struct {
	customsCfg *internal.CustomsCfg
}

func (s *SasXmlServiceImpl) GenInv101Xml(inv101 Inv101, declareFlag string) (xmlBytes []byte, err error) {
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

func (s *SasXmlServiceImpl) GenSas121Xml(sas121 Sas121, declareFlag string) (xmlBytes []byte, err error) {
	// 校验
	if declareFlag != "0" && declareFlag != "1" {
		err = fmt.Errorf("申报标志(declareFlag)必须是0或1")
		return
	}
	// 生成sas121Xml
	sas121Xml := Sas121Xml{}
	sas121Xml.Object.Package.EnvelopInfo.MessageType = "SAS121"
	sas121Xml.Object.Package.DataInfo.BusinessData.DeclareFlag = declareFlag
	sas121Xml.Object.Package.DataInfo.BusinessData.PassPortMessage.OperCusRegCode = s.customsCfg.OperCusRegCode
	sas121Xml.Object.Package.DataInfo.BusinessData.PassPortMessage.PassportHead = sas121.Head
	sas121Xml.Object.Package.DataInfo.BusinessData.PassPortMessage.PassportList = sas121.List
	sas121Xml.Object.Package.DataInfo.BusinessData.PassPortMessage.PassportAcmp = sas121.Acmp
	// 保存xml
	var xmlBodyBytes []byte
	if xmlBodyBytes, err = xml.Marshal(sas121Xml); err != nil {
		return
	}
	xmlBytes = []byte(xml.Header + string(xmlBodyBytes))
	return
}

func (s *SasXmlServiceImpl) ParseInv201Xml(xmlBytes []byte) (inv201 Inv201, err error) {
	inv201Xml := Inv201Xml{}
	if err = xml.Unmarshal(xmlBytes, &inv201Xml); err != nil {
		return
	}

	inv201.HdeApprResult = inv201Xml.DataInfo.BusinessData.Inv201.HdeApprResult
	inv201.Head = inv201Xml.DataInfo.BusinessData.Inv201.BondInvtBsc
	inv201.List = inv201Xml.DataInfo.BusinessData.Inv201.BondInvtDtl
	return
}

func (s *SasXmlServiceImpl) ParseInv202Xml(xmlBytes []byte) (inv202 Inv202, err error) {
	inv202Xml := Inv202Xml{}
	if err = xml.Unmarshal(xmlBytes, &inv202Xml); err != nil {
		return
	}
	inv202.InvApprResult = inv202Xml.DataInfo.BusinessData.Inv202.InvApprResult
	return
}

func (s *SasXmlServiceImpl) ParseInv211Xml(xmlBytes []byte) (inv211 Inv211, err error) {
	inv211Xml := Inv211Xml{}
	if err = xml.Unmarshal(xmlBytes, &inv211Xml); err != nil {
		return
	}
	inv211.Head = inv211Xml.DataInfo.BusinessData.Inv211.BondInvtBsc
	inv211.List = inv211Xml.DataInfo.BusinessData.Inv211.BwsDt
	return
}

func (s *SasXmlServiceImpl) ParseSas221Xml(xmlBytes []byte) (sas221 Sas221, err error) {
	sas221Xml := Sas221Xml{}
	if err = xml.Unmarshal(xmlBytes, &sas221Xml); err != nil {
		return
	}
	sas221.HdeApprResult = sas221Xml.DataInfo.BusinessData.Sas221.HdeApprResult
	sas221.Head = sas221Xml.DataInfo.BusinessData.Sas221.SasPassportBsc
	sas221.List = sas221Xml.DataInfo.BusinessData.Sas221.SasPassportDt
	sas221.Acmp = sas221Xml.DataInfo.BusinessData.Sas221.SasPassportRlt
	return
}

func (s *SasXmlServiceImpl) ParseSas223Xml(xmlBytes []byte) (sas223 Sas223, err error) {
	sas223Xml := Sas223Xml{}
	if err = xml.Unmarshal(xmlBytes, &sas223Xml); err != nil {
		return
	}
	sas223.HdeApprResult = sas223Xml.DataInfo.BusinessData.Sas223.HdeApprResult
	sas223.Head = sas223Xml.DataInfo.BusinessData.Sas223.SasPassportBsc
	return
}

func (s *SasXmlServiceImpl) ParseSas224Xml(xmlBytes []byte) (sas224 Sas224, err error) {
	sas224Xml := Sas224Xml{}
	if err = xml.Unmarshal(xmlBytes, &sas224Xml); err != nil {
		return
	}
	sas224.HdeApprResult = sas224Xml.DataInfo.BusinessData.Sas224.HdeApprResult
	sas224.Head = sas224Xml.DataInfo.BusinessData.Sas224.SasPassportBsc
	return
}
