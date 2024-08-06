package sas

import (
	"encoding/xml"

	"github.com/ronannnn/gx-customs-bridge/internal/services/customs/common"
	"github.com/shopspring/decimal"
)

type Sas224Head struct {
	// 非业务相关字段
	ChgTmsCnt *string `json:"chgTmsCnt"` // 变更次数

	// 编号
	SasPassportPreentNo *string `json:"sasPassportPreentNo"` // 核放单预录入编号(同Sas121的SeqNo)
	PassportNo          *string `json:"passportNo"`          // 核放单编号

	// 关联信息
	RltTbTypecd *string `json:"rltTbTypecd"` // 关联单证类型代码
	RltNo       *string `json:"rltNo"`       // 关联单证编号

	// 代码
	DclTypecd      *string `json:"dclTypecd"`      // 申报类型代码
	IoTypecd       *string `json:"ioTypecd"`       // 进出标志代码
	PassportTypecd *string `json:"passportTypecd"` // 核放单类型代码
	BindTypecd     *string `json:"bindTypecd"`     // 绑定类型代码
	MasterCuscd    *string `json:"masterCuscd"`    // 主管关区代码

	// 区内企业
	AreainEtpsSccd *string `json:"areainEtpsSccd"` // 区内企业社会信用代码
	AreainEtpsno   *string `json:"areainEtpsno"`   // 区内企业编号
	AreainEtpsNm   *string `json:"areainEtpsNm"`   // 区内企业名称

	// 车辆信息
	VehicleNo      *string          `json:"vehicleNo"`      // 承运车牌号
	VehicleIcNo    *string          `json:"vehicleIcNo"`    // IC卡号(电子车牌）
	VehicleWt      *decimal.Decimal `json:"vehicleWt"`      // 车自重
	VehicleFrameWt *decimal.Decimal `json:"vehicleFrameWt"` // 车架重
	VehicleFrameNo *string          `json:"vehicleFrameNo"` // 车架号

	// 集装箱信息
	ContainerNo   *string `json:"containerNo"`   // 集装箱号
	ContainerWt   *string `json:"containerWt"`   // 集装箱自重
	ContainerType *string `json:"containerType"` // 集装箱箱型

	// 货物信息
	TotalGrossWt *decimal.Decimal `json:"totalGrossWt"` // 货物总毛重
	TotalNetWt   *decimal.Decimal `json:"totalNetWt"`   // 货物总净重
	TotalWt      *decimal.Decimal `json:"totalWt"`      // 总重量(包括车辆自重)

	// 卡口信息
	PassCollectWt  *decimal.Decimal `json:"passCollectWt"`  // 卡口地磅采集重量
	WtError        *decimal.Decimal `json:"wtError"`        // 重量比对误差
	PassID         *string          `json:"passID"`         // 卡口ID1
	SecdPassId     *string          `json:"secPassId"`      // 卡口ID2
	PassTime       *string          `json:"passTime"`       // 过卡时间1
	SecdPassTime   *string          `json:"secPassTime"`    // 过卡时间2
	Stucd          *string          `json:"stucd"`          // 状态代码(0：已申请，1：已审批，2：已过卡，3：已过一卡，4：已过二卡，5：已删除，6：已作废。其中2单卡模式专用，3/4双卡口模式专用)
	EmapvMarkcd    *string          `json:"emapvMarkcd"`    // 电子口岸标志代码(1：通过，2：转岗，3：退单)
	LogisticsStucd *string          `json:"logisticsStucd"` // 到检状态(一票多车类型核放单专用，用于标识哪些车已到达待检及判别最后过卡核放单。0：未到检，1：已待检，默认为未到检。)
	OwnerSystem    *string          `json:"ownerSystem"`    // 所属系统(1-特殊区域，2-保税物流。此字段用于区域、物流系统筛选各自数据，界面不显示。字段值由交接入库时返填)

	// 申报信息
	DclErConc *string `json:"dclErConc"` // 申请人及联系方式
	DclTime   *string `json:"dclTime"`   // 申请时间(核放单申请时间；核放单申请后系统按照发送时间自动生成)

	// 其他
	Rmk  *string `json:"rmk"`  // 备注
	Col1 *string `json:"col1"` // 备用字段1(到货确认标志 1：是，0：否)
}

type Sas224 struct {
	HdeApprResult common.HdeApprResult `json:"hdeApprResult"`
	Head          Sas224Head           `json:"head"`
}

type Sas224Xml struct {
	XMLName     xml.Name `xml:"Package"`
	EnvelopInfo common.EnvelopInfo
	DataInfo    struct {
		PocketInfo   common.PocketInfo
		BusinessData struct {
			Sas224 struct {
				HdeApprResult  common.HdeApprResult
				SasPassportBsc Sas224Head
			} `xml:"SAS224"`
		}
	}
}
