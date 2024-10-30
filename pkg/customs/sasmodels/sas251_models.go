package sasmodels

import (
	"encoding/xml"

	"github.com/ronannnn/gx-customs-bridge/pkg/customs/commonmodels"
)

// 两步申报核放单审核回执/回执补发送

type Sas251Head struct {
	SasPassportPreentNo string `json:"sasPassportPreentNo" xml:"sasPassportPreentNo"` // 核放单预录入统一编号
	PassportNo          string `json:"passportNo" xml:"passportNo"`                   // 核放单编号
	ChgTmsCnt           string `json:"chgTmsCnt" xml:"chgTmsCnt"`                     // 变更次数

	DclTypecd   string `json:"dclTypecd" xml:"dclTypecd"`     // 申报类型(1-备案、2-变更、3-作废。目前核放单只允许备案--1)
	MasterCuscd string `json:"masterCuscd" xml:"masterCuscd"` // 主管关区代码
	EmapvMarkcd string `json:"emapvMarkcd" xml:"emapvMarkcd"` // 审批标志代码(1：通过、3：退单)

	// 区内企业
	AreainEtpsSccd string `json:"areainEtpsSccd" xml:"areainEtpsSccd"` // 社会信用代码
	AreainEtpsNo   string `json:"areainEtpsNo" xml:"areainEtpsNo"`     // 编号
	AreainEtpsNm   string `json:"areainEtpsNm" xml:"areainEtpsNm"`     // 名称

	// 车辆信息
	VehicleNo   string `json:"vehicleNo" xml:"vehicleNo"`     // 承运车牌号
	VehicleIcNo string `json:"vehicleIcNo" xml:"vehicleIcNo"` // IC卡号(电子车牌)

	// 集装箱信息
	ContainerNo string `json:"containerNo" xml:"containerNo"` // 集装箱号

	// 卡口信息
	LogisticsStucd string `json:"logisticsStucd" xml:"logisticsStucd"` // 到检状态(0：未到检、1：已待检。默认为未到检。)
	PassId         string `json:"passId" xml:"passId"`                 // 卡口唯一ID，系统自动返填。双卡模式下，为一卡ID
	SecdPassId     string `json:"secdPassId" xml:"secdPassId"`         // 二卡ID，系统自动返填。双卡模式专用
	PassTime       string `json:"passTime" xml:"passTime"`             // 过卡时间1(过卡时间，卡口抬杆后系统自动返填。双卡模式，为过一卡时间)
	SecdPassTime   string `json:"secdPassTime" xml:"secdPassTime"`     // 过卡时间2(过二卡时间，卡口抬杆后系统自动返填。双卡模式专用)
	Stucd          string `json:"stucd" xml:"stucd"`                   // 状态代码(0：已申请、1：已审批、2：已过卡、3：已过一卡、4：已过二卡、5：已删除。其中2单卡模式专用，3/4双卡口模式专用)

	// 其他
	DclErConc   string `json:"dclErConc" xml:"dclErConc"`     // 申请人及联系方式
	DclTime     string `json:"dclTime" xml:"dclTime"`         // 申报时间(核放单申请时间；核放单申请后系统按照发送时间自动生成)
	OwnerSystem string `json:"ownerSystem" xml:"ownerSystem"` // 业务系统(1-特殊区域、2-保税物流)
	Rmk         string `json:"rmk" xml:"rmk"`                 // 备注
	Col1        string `json:"col1" xml:"col1"`               // 备用1
	Col2        string `json:"col2" xml:"col2"`               // 备用2
	Col3        string `json:"col3" xml:"col3"`               // 备用3
	Col4        string `json:"col4" xml:"col4"`               // 备用4
}

type Sas251RltList struct {
	PassportNo string `json:"passportNo" xml:"passportNo"` // 核放单编号
	EntryId    string `json:"entryId" xml:"entryId"`       // 报关单号
}

type Sas251 struct {
	commonmodels.MqDataCommonPayload
	HdeApprResult commonmodels.HdeApprResult `json:"hdeApprResult"`
	Head          Sas251Head                 `json:"head"`
	RltList       []Sas251RltList            `json:"rltList"`
}

type Sas251Xml struct {
	XMLName     xml.Name `xml:"Package"`
	EnvelopInfo commonmodels.EnvelopInfo
	DataInfo    struct {
		PocketInfo   commonmodels.PocketInfo
		BusinessData struct {
			Sas251 struct {
				HdeApprResult       commonmodels.HdeApprResult
				CheckInfo           commonmodels.CheckInfo
				Sas2stepPassportBsc Sas251Head
				Sas2stepPassportRlt []Sas251RltList
			} `xml:"SAS251"`
		} `xml:"BussinessData"`
	}
}
