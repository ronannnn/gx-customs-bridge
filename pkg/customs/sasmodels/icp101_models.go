package sasmodels

import (
	"encoding/xml"

	"github.com/ronannnn/gx-customs-bridge/pkg/customs/commonmodels"
)

type Icp101Head struct {
	// 编号
	SasPassportPreentNo *string `json:"sasPassportPreentNo" validate:"omitempty,len=18"`                          // 核放单预录入统一编号
	PassportNo          *string `json:"passportNo" validate:"omitempty,len=24"`                                   // 核放单编号
	EtpsPreentNo        *string `json:"etpsPreentNo" validate:"required,not_blank" gorm:"uniqueIndex" mod:"trim"` // 企业内部编号

	// 代码
	DclTypecd   *string `json:"dclTypecd" validate:"required,oneof=1 2 3"`     // 申报类型(1-备案、2-变更、3-作废。目前核放单只允许备案--1)
	MasterCuscd *string `json:"masterCuscd" validate:"required,len=4,numeric"` // 主管关区代码

	// 区内企业
	AreainEtpsSccd *string `json:"areainEtpsSccd" validate:"required,len=18"`             // 社会信用代码
	AreainEtpsNo   *string `json:"areainEtpsNo" validate:"required,len=10"`               // 编号
	AreainEtpsNm   *string `json:"areainEtpsNm" validate:"required,not_blank" mod:"trim"` // 名称
	// 申报企业
	DclEtpsSccd *string `json:"dclEtpsSccd" validate:"required,len=18"`             // 社会信用代码
	DclEtpsNo   *string `json:"dclEtpsNo" validate:"required,len=10"`               // 编号
	DclEtpsNm   *string `json:"dclEtpsNm" validate:"required,not_blank" mod:"trim"` // 名称
	// 录入单位
	InputCreditCode *string `json:"inputCreditCode" validate:"required,len=18"`         // 社会信用代码
	InputCode       *string `json:"inputCode" validate:"required,len=10"`               // 代码
	InputName       *string `json:"inputName" validate:"required,not_blank" mod:"trim"` // 名称

	// 车辆信息
	VehicleNo   *string `json:"vehicleNo" validate:"required" mod:"trim"` // 承运车牌号
	VehicleIcNo *string `json:"vehicleIcNo"`                              // IC卡号(电子车牌)

	// 集装箱信息
	ContainerNo *string `json:"containerNo"` // 集装箱号

	// 卡口信息
	PassTime     *string `json:"passTime"`     // 过卡时间1(过卡时间，卡口抬杆后系统自动返填。双卡模式，为过一卡时间)
	SecdPassTime *string `json:"secdPassTime"` // 过卡时间2(过二卡时间，卡口抬杆后系统自动返填。双卡模式专用)

	// 其他
	DclErConc *string `json:"dclErConc"` // 申请人及联系方式
	Rmk       *string `json:"rmk"`       // 备注
}

type Icp101RltList struct {
	EntryId *string `json:"entryId" validate:"required,len=18"` // 报关单号
}

type Icp101 struct {
	Head    Icp101Head      `json:"head" validate:"required"`
	RltList []Icp101RltList `json:"rltList" validate:"required,dive,required"`
}

type Icp101Xml struct {
	XMLName xml.Name `xml:"Signature"`
	Object  struct {
		Package struct {
			EnvelopInfo commonmodels.EnvelopInfo
			DataInfo    struct {
				BusinessData struct {
					DeclareFlag          string `xml:"DelcareFlag"` // 这里xml的key和struct的key不一致，海关那个应该是拼写错误
					Sas2sPassPortMessage struct {
						SysId             string
						OperCusRegCode    string
						Sas2sPassportHead Icp101Head
						Sas2sPassportRlt  []Icp101RltList
					}
				} `xml:"BussinessData"`
			}
		}
	}
}
