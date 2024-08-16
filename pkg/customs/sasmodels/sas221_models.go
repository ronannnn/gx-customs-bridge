package sasmodels

import (
	"encoding/xml"

	"github.com/ronannnn/gx-customs-bridge/pkg/customs/commonmodels"
)

type Sas221Head struct {
	// 非业务相关字段
	ChgTmsCnt string `json:"chgTmsCnt"` // 变更次数

	// 编号
	SasPassportPreentNo string `json:"sasPassportPreentNo" xml:"sasPassportPreentNo"` // 核放单预录入编号(同Sas121的SeqNo)
	PassportNo          string `json:"passportNo" xml:"passportNo"`                   // 核放单编号

	// 关联信息
	RltTbTypecd string `json:"rltTbTypecd" xml:"rltTbTypecd"` // 关联单证类型代码
	RltNo       string `json:"rltNo" xml:"rltNo"`             // 关联单证编号

	// 代码
	DclTypecd      string `json:"dclTypecd" xml:"dclTypecd"`           // 申报类型代码
	IoTypecd       string `json:"ioTypecd" xml:"ioTypecd"`             // 进出标志代码
	PassportTypecd string `json:"passportTypecd" xml:"passportTypecd"` // 核放单类型代码
	BindTypecd     string `json:"bindTypecd" xml:"bindTypecd"`         // 绑定类型代码
	MasterCuscd    string `json:"masterCuscd" xml:"masterCuscd"`       // 主管关区代码

	// 区内企业
	AreainEtpsSccd string `json:"areainEtpsSccd" xml:"areainEtpsSccd"` // 区内企业社会信用代码
	AreainEtpsno   string `json:"areainEtpsno" xml:"areainEtpsno"`     // 区内企业编号
	AreainEtpsNm   string `json:"areainEtpsNm" xml:"areainEtpsNm"`     // 区内企业名称

	// 车辆信息
	VehicleNo      string `json:"vehicleNo" xml:"vehicleNo"`           // 承运车牌号
	VehicleIcNo    string `json:"vehicleIcNo" xml:"vehicleIcNo"`       // IC卡号(电子车牌）
	VehicleWt      string `json:"vehicleWt" xml:"vehicleWt"`           // 车自重
	VehicleFrameWt string `json:"vehicleFrameWt" xml:"vehicleFrameWt"` // 车架重
	VehicleFrameNo string `json:"vehicleFrameNo" xml:"vehicleFrameNo"` // 车架号

	// 集装箱信息
	ContainerNo   string `json:"containerNo" xml:"containerNo"`     // 集装箱号
	ContainerWt   string `json:"containerWt" xml:"containerWt"`     // 集装箱自重
	ContainerType string `json:"containerType" xml:"containerType"` // 集装箱箱型

	// 货物信息
	TotalGrossWt string `json:"totalGrossWt" xml:"totalGrossWt"` // 货物总毛重
	TotalNetWt   string `json:"totalNetWt" xml:"totalNetWt"`     // 货物总净重
	TotalWt      string `json:"totalWt" xml:"totalWt"`           // 总重量(包括车辆自重)

	// 卡口信息
	PassCollectWt  string `json:"passCollectWt" xml:"passCollectWt"`   // 卡口地磅采集重量
	WtError        string `json:"wtError" xml:"wtError"`               // 重量比对误差
	PassID         string `json:"passID" xml:"passID"`                 // 卡口ID1
	SecdPassId     string `json:"secPassId" xml:"secPassId"`           // 卡口ID2
	PassTime       string `json:"passTime" xml:"passTime"`             // 过卡时间1
	SecdPassTime   string `json:"secPassTime" xml:"secPassTime"`       // 过卡时间2
	Stucd          string `json:"stucd" xml:"stucd"`                   // 状态代码(0：已申请，1：已审批，2：已过卡，3：已过一卡，4：已过二卡，5：已删除，6：已作废。其中2单卡模式专用，3/4双卡口模式专用)
	EmapvMarkcd    string `json:"emapvMarkcd" xml:"emapvMarkcd"`       // 电子口岸标志代码(1：通过，2：转岗，3：退单)
	LogisticsStucd string `json:"logisticsStucd" xml:"logisticsStucd"` // 到检状态(一票多车类型核放单专用，用于标识哪些车已到达待检及判别最后过卡核放单。0：未到检，1：已待检，默认为未到检。)
	OwnerSystem    string `json:"ownerSystem" xml:"ownerSystem"`       // 所属系统(1-特殊区域，2-保税物流。此字段用于区域、物流系统筛选各自数据，界面不显示。字段值由交接入库时返填)

	// 申报信息
	DclErConc string `json:"dclErConc" xml:"dclErConc"` // 申请人及联系方式
	DclTime   string `json:"dclTime" xml:"dclTime"`     // 申请时间(核放单申请时间；核放单申请后系统按照发送时间自动生成)

	// 其他
	Rmk  string `json:"rmk" xml:"rmk"`   // 备注
	Col1 string `json:"col1" xml:"col1"` // 备用字段1(到货确认标志 1：是，0：否)
}

type Sas221List struct {
	// 非业务相关字段
	ChgTmsCnt string `json:"chgTmsCnt" xml:"chgTmsCnt"` // 变更次数

	// 编号
	PassportNo string `json:"passportNo" xml:"passportNo"` // 表头的核放单编号(备案时为空，变更时填写)

	// 序号
	PassportSeqNo string `json:"passportSeqNo" xml:"passportSeqNo"` // 自然序号，从1开始
	RltGdsSeqno   string `json:"rltGdsSeqno" xml:"rltGdsSeqno"`     // 关联商品序号

	// 商品信息
	GdsMtno   string `json:"gdsMtno" xml:"gdsMtno"`     // 商品料号
	Gdecd     string `json:"gdecd" xml:"gdecd"`         // 商品编码
	GdsNm     string `json:"gdsNm" xml:"gdsNm"`         // 商品名称
	DclUnitcd string `json:"dclUnitcd" xml:"dclUnitcd"` // 申报计量单位代码
	DclQty    string `json:"dclQty" xml:"dclQty"`       // 申报数量
	GrossWt   string `json:"grossWt" xml:"grossWt"`     // 毛重
	NetWt     string `json:"netWt" xml:"netWt"`         // 净重

	// others
	Rmk string `json:"rmk" xml:"rmk"` // 备注
}

type Sas221Acmp struct {
	PassportNo  string `json:"passportNo" xml:"passportNo"`   // 表头的核放单编号(备案时为空，变更时填写)
	RltTbTypecd string `json:"rltTbTypecd" xml:"rltTbTypecd"` // 关联单证类型代码(应该和表头的RltTbTypecd一致)
	RltNo       string `json:"rltNo" xml:"rltNo"`             // 关联单证编号
}

type Sas221 struct {
	HdeApprResult commonmodels.HdeApprResult `json:"hdeApprResult"`
	CheckInfo     commonmodels.CheckInfo     `json:"checkInfo"`
	Head          Sas221Head                 `json:"head"`
	List          []Sas221List               `json:"list"`
	Acmp          []Sas221Acmp               `json:"acmp"`
}

type Sas221Xml struct {
	XMLName     xml.Name `xml:"Package"`
	EnvelopInfo commonmodels.EnvelopInfo
	DataInfo    struct {
		PocketInfo   commonmodels.PocketInfo
		BusinessData struct {
			Sas221 struct {
				HdeApprResult  commonmodels.HdeApprResult
				CheckInfo      commonmodels.CheckInfo
				SasPassportBsc Sas221Head
				SasPassportDt  []Sas221List
				SasPassportRlt []Sas221Acmp
			} `xml:"SAS221"`
		} `xml:"BussinessData"`
	}
}
