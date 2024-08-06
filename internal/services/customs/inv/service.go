package inv

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/ronannnn/gx-customs-bridge/internal"
)

type ExportInv101Service interface {
	GenXml(inv101 Inv101, declareFlag string) error
}

func ProvideExportInv101Service(customsCfg internal.CustomsCfg) ExportInv101Service {
	return &ExportInv101ServiceImpl{
		customsCfg: customsCfg,
	}
}

type ExportInv101ServiceImpl struct {
	customsCfg internal.CustomsCfg
}

func (s *ExportInv101ServiceImpl) GenXml(inv101 Inv101, declareFlag string) (err error) {
	// 校验
	if declareFlag != "0" && declareFlag != "1" {
		err = fmt.Errorf("申报标志(declareFlag)必须是0或1")
		return
	}
	// 替换部分数据
	inv101.Head.IcCardNo = &(s.customsCfg.IcCardNo)
	// 生成inv101Xml
	inv101Xml := Inv101Xml{}
	inv101Xml.Object.Package.EnvelopInfo.MessageType = "INV101"
	inv101Xml.Object.Package.DataInfo.BusinessData.DeclareFlag = declareFlag
	inv101Xml.Object.Package.DataInfo.BusinessData.InvtMessage.SysId = s.customsCfg.Inv101SysId
	inv101Xml.Object.Package.DataInfo.BusinessData.InvtMessage.OperCusRegCode = s.customsCfg.OperCusRegCode
	inv101Xml.Object.Package.DataInfo.BusinessData.InvtMessage.InvtHeadType = inv101.Head
	inv101Xml.Object.Package.DataInfo.BusinessData.InvtMessage.InvtListType = inv101.List
	// 保存xml
	var output []byte
	if output, err = xml.Marshal(inv101Xml); err != nil {
		return
	}
	os.Stdout.Write([]byte(xml.Header)) //输出预定义的xml头  <?xml version="1.0" encoding="UTF-8"?>
	os.Stdout.Write(output)
	return
}
