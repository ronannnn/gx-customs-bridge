package sasmodels

import (
	"encoding/xml"

	"github.com/ronannnn/gx-customs-bridge/pkg/customs/commonmodels"
	"github.com/shopspring/decimal"
)

type Sas121Head struct {
	// 编号
	SeqNo        *string `json:"seqNo"`        // 预录入统一编号
	PassportNo   *string `json:"passportNo"`   // 核放单编号
	EtpsPreentNo *string `json:"etpsPreentNo"` // 企业内部编号

	// 关联信息
	RltTbTypecd *string `json:"rltTbTypecd"` // 关联单证类型代码
	RltNo       *string `json:"rltNo"`       // 关联单证编号

	// 代码
	DclTypecd      *string `json:"dclTypecd"`      // 申报类型代码
	IoTypecd       *string `json:"ioTypecd"`       // 进出标志代码
	PassportTypecd *string `json:"passportTypecd"` // 核放单类型代码
	BindTypecd     *string `json:"bindTypecd"`     // 绑定类型代码
	MasterCuscd    *string `json:"masterCuscd"`    // 主管关区代码

	// 区内信息
	AreainOriactNo *string `json:"areainOriactNo"` // 区内账册号
	// 区内企业
	AreainEtpsSccd *string `json:"areainEtpsSccd"` // 区内企业社会信用代码
	AreainEtpsno   *string `json:"areainEtpsno"`   // 区内企业编号
	AreainEtpsNm   *string `json:"areainEtpsNm"`   // 区内企业名称
	// 申报企业
	DclEtpsSccd *string `json:"dclEtpsSccd"` // 申报企业社会信用代码
	DclEtpsno   *string `json:"dclEtpsno"`   // 申报企业编号
	DclEtpsNm   *string `json:"dclEtpsNm"`   // 申报企业名称
	// 录入单位
	InputSccd *string `json:"inputSccd"` // 录入单位社会信用代码
	InputCode *string `json:"inputCode"` // 录入单位编号
	InputName *string `json:"inputName"` // 录入单位名称

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
	PassTime     *string `json:"passTime" xml:"passTime"`       // 过卡时间1
	SecdPassTime *string `json:"secPassTime" xml:"secPassTime"` // 过卡时间2
	Stucd        *string `json:"stucd" xml:"stucd"`             // 状态代码(0：已申请，1：已审批，2：已过卡，3：已过一卡，4：已过二卡，5：已删除，6：已作废。其中2单卡模式专用，3/4双卡口模式专用)

	// 申报信息
	DclErConc *string `json:"dclErConc"` // 申请人及联系方式
	DclTime   *string `json:"dclTime"`   // 申报日期

	// 其他
	Rmk  *string `json:"rmk"`  // 备注
	Col1 *string `json:"col1"` // 备用字段1(到货确认标志 1：是，0：否)
	Col2 *string `json:"col2"` // 备用字段2
	Col3 *string `json:"col3"` // 备用字段3
	Col4 *string `json:"col4"` // 备用字段4
}

type Sas121List struct {
	// 编号
	SeqNo      *string `json:"seqNo"`      // 表头的核放单预录入编号
	PassportNo *string `json:"passportNo"` // 表头的核放单编号(备案时为空，变更时填写)

	// 序号
	PassportSeqNo *string `json:"passportSeqNo"` // 自然序号，从1开始
	RltGdsSeqno   *string `json:"rltGdsSeqno"`   // 关联商品序号

	// 商品信息
	GdsMtno   *string          `json:"gdsMtno"`   // 商品料号
	Gdecd     *string          `json:"gdecd"`     // 商品编码
	GdsNm     *string          `json:"gdsNm"`     // 商品名称
	DclUnitcd *string          `json:"dclUnitcd"` // 申报计量单位代码
	DclQty    *decimal.Decimal `json:"dclQty"`    // 申报数量
	GrossWt   *decimal.Decimal `json:"grossWt"`   // 毛重
	NetWt     *decimal.Decimal `json:"netWt"`     // 净重

	// others
	Rmk  *string `json:"rmk"`  // 备注
	Col1 *string `json:"col1"` // 备用字段1
	Col2 *string `json:"col2"` // 备用字段2
	Col3 *string `json:"col3"` // 备用字段3
	Col4 *string `json:"col4"` // 备用字段4
}

type Sas121Acmp struct {
	SeqNo         string `json:"seqNo"`         // 表头的核放单预录入编号
	PassPortNo    string `json:"passPortNo"`    // 表头的核放单编号(备案时为空，变更时填写)
	RtlBillTypecd string `json:"rtlBillTypecd"` // 关联单证类型代码(应该和表头的RltTbTypecd一致)
	RtlBillNo     string `json:"rtlBillNo"`     // 关联单证编号
}

type Sas121 struct {
	Head Sas121Head   `json:"head"`
	List []Sas121List `json:"list"`
	Acmp []Sas121Acmp `json:"acmp"`
}

type Sas121Xml struct {
	XMLName xml.Name `xml:"Signature"`
	Object  struct {
		Package struct {
			EnvelopInfo commonmodels.EnvelopInfo
			DataInfo    struct {
				BusinessData struct {
					DeclareFlag     string `xml:"DelcareFlag"` // 这里xml的key和struct的key不一致，海关那个应该是拼写错误
					PassPortMessage struct {
						OperCusRegCode string
						PassportHead   Sas121Head
						PassportList   []Sas121List
						PassportAcmp   []Sas121Acmp
					}
				} `xml:"BussinessData"`
			}
		}
	}
}
