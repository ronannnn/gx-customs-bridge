package sasmodels

import (
	"encoding/xml"

	"github.com/ronannnn/gx-customs-bridge/pkg/customs/commonmodels"
	"github.com/ronannnn/infra/models"
)

type Sas121Head struct {
	// 编号
	SeqNo        *string `json:"seqNo" validate:"omitempty,len=18"`                                        // 预录入统一编号
	PassportNo   *string `json:"passportNo" validate:"omitempty,len=24"`                                   // 核放单编号
	EtpsPreentNo *string `json:"etpsPreentNo" validate:"required,not_blank" gorm:"uniqueIndex" mod:"trim"` // 企业内部编号

	// 关联信息
	RltTbTypecd *string `json:"rltTbTypecd" validate:"required,oneof=1 2 3"` // 关联单证类型代码(1-核注清单 2-出入库单 3-提运单)
	RltNo       *string `json:"rltNo" validate:"required" mod:"trim"`        // 关联单证编号

	// 代码
	DclTypecd      *string `json:"dclTypecd" validate:"required,oneof=1 3"`              // 申报类型代码(1-备案 3-作废)
	IoTypecd       *string `json:"ioTypecd" validate:"required,oneof=I E"`               // 进出标志代码(I-进口 E-出口)
	PassportTypecd *string `json:"passportTypecd" validate:"required,oneof=1 2 3 4 5 6"` // 核放单类型代码(1-先入区后报关 2-一线一体化进出区 3-二线进出区 4-非报关进出区 5-卡口登记货物 6-空车进出区)
	BindTypecd     *string `json:"bindTypecd" validate:"required,oneof=1 2 3"`           // 绑定类型代码(1-一车多票 2-一票一车3-一票多车)
	MasterCuscd    *string `json:"masterCuscd" validate:"required,len=4,numeric"`        // 主管关区代码

	// 区内信息
	AreainOriactNo *string `json:"areainOriactNo" mod:"trim"` // 区内账册号
	// 区内企业
	AreainEtpsSccd *string `json:"areainEtpsSccd" validate:"required,len=18"`             // 区内企业社会信用代码
	AreainEtpsno   *string `json:"areainEtpsno" validate:"required,len=10"`               // 区内企业编号
	AreainEtpsNm   *string `json:"areainEtpsNm" validate:"required,not_blank" mod:"trim"` // 区内企业名称
	// 申报企业
	DclEtpsSccd *string `json:"dclEtpsSccd" validate:"required,len=18"`             // 申报企业社会信用代码
	DclEtpsno   *string `json:"dclEtpsno" validate:"required,len=10"`               // 申报企业编号
	DclEtpsNm   *string `json:"dclEtpsNm" validate:"required,not_blank" mod:"trim"` // 申报企业名称
	// 录入单位
	InputSccd *string `json:"inputSccd" validate:"required,len=18"`               // 录入单位社会信用代码
	InputCode *string `json:"inputCode" validate:"required,len=10"`               // 录入单位编号
	InputName *string `json:"inputName" validate:"required,not_blank" mod:"trim"` // 录入单位名称

	// 车辆信息
	VehicleNo      *string             `json:"vehicleNo" validate:"required" mod:"trim"`  // 承运车牌号
	VehicleIcNo    *string             `json:"vehicleIcNo"`                               // IC卡号(电子车牌）
	VehicleWt      *models.DecimalSafe `json:"vehicleWt" validate:"required,d_gt=0"`      // 车自重
	VehicleFrameWt *models.DecimalSafe `json:"vehicleFrameWt" validate:"required,d_gt=0"` // 车架重
	VehicleFrameNo *string             `json:"vehicleFrameNo"`                            // 车架号

	// 集装箱信息
	ContainerNo   *string `json:"containerNo"`   // 集装箱号
	ContainerWt   *string `json:"containerWt"`   // 集装箱自重
	ContainerType *string `json:"containerType"` // 集装箱箱型

	// 货物信息
	TotalGrossWt *models.DecimalSafe `json:"totalGrossWt" validate:"required,d_gt=0"` // 货物总毛重
	TotalNetWt   *models.DecimalSafe `json:"totalNetWt" validate:"required,d_gt=0"`   // 货物总净重
	TotalWt      *models.DecimalSafe `json:"totalWt" validate:"required,d_gt=0"`      // 总重量(包括车辆自重)

	// 卡口信息
	PassTime     *string `json:"passTime" xml:"passTime"`         // 过卡时间1
	SecdPassTime *string `json:"secdPassTime" xml:"secdPassTime"` // 过卡时间2
	Stucd        *string `json:"stucd" xml:"stucd"`               // 状态代码(0：已申请，1：已审批，2：已过卡，3：已过一卡，4：已过二卡，5：已删除，6：已作废。其中2单卡模式专用，3/4双卡口模式专用)

	// 申报信息
	DclErConc *string `json:"dclErConc" validate:"required"` // 申请人及联系方式
	DclTime   *string `json:"dclTime" validate:"required"`   // 申报日期

	// 其他
	Rmk  *string `json:"rmk"`  // 备注
	Col1 *string `json:"col1"` // 备用字段1(到货确认标志 1：是，0：否)
	Col2 *string `json:"col2"` // 备用字段2
	Col3 *string `json:"col3"` // 备用字段3
	Col4 *string `json:"col4"` // 备用字段4
}

type Sas121List struct {
	// 编号
	SeqNo      *string `json:"seqNo" validate:"omitempty,len=18"`      // 表头的核放单预录入编号
	PassportNo *string `json:"passportNo" validate:"omitempty,len=24"` // 表头的核放单编号(备案时为空，变更时填写)

	// 序号
	PassportSeqNo *string `json:"passportSeqNo" validate:"required,numeric"` // 自然序号，从1开始
	RltGdsSeqno   *string `json:"rltGdsSeqno" validate:"required,numeric"`   // 关联商品序号

	// 商品信息
	GdsMtno   *string             `json:"gdsMtno" validate:"required" mod:"trim"`      // 商品料号
	Gdecd     *string             `json:"gdecd" validate:"required,numeric"`           // 商品编码
	GdsNm     *string             `json:"gdsNm" validate:"required" mod:"trim"`        // 商品名称
	DclUnitcd *string             `json:"dclUnitcd" validate:"required,len=3,numeric"` // 申报计量单位代码
	DclQty    *models.DecimalSafe `json:"dclQty" validate:"required,d_gt=0"`           // 申报数量
	GrossWt   *models.DecimalSafe `json:"grossWt" validate:"required,d_gt=0"`          // 毛重
	NetWt     *models.DecimalSafe `json:"netWt" validate:"required,d_gt=0"`            // 净重

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
