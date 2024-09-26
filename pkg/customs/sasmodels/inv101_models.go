package sasmodels

import (
	"encoding/xml"

	"github.com/ronannnn/gx-customs-bridge/pkg/customs/commonmodels"
	"github.com/ronannnn/infra/models"
)

// Inv101Head 核注清单表头
type Inv101Head struct {
	// 非业务相关字段
	ChgTmsCnt *string `json:"chgTmsCnt"` // 变更次数

	// 编号
	SeqNo           *string `json:"seqNo" validate:"omitempty,len=18"`                                // 预录入统一编号
	EtpsInnerInvtNo *string `json:"etpsInnerInvtNo" validate:"required,not_blank" gorm:"uniqueIndex"` // 企业内部清单编号
	// 保税核注清单海关编号为18位，其中
	// 第1－2位为QD，表示核注清单
	// 第3－6位为接受申报海关的编号（海关规定的《关区代码表》中相应海关代码）
	// 第7-8位为海关接受申报的公历年份
	// 第9位为进出口标志（“I”为进口，“E”为出口）
	// 后9位为顺序编号
	BondInvtNo *string `json:"bondInvtNo" validate:"omitempty,len=18"` // 保税清单编号
	PutrecNo   *string `json:"putrecNo" validate:"required,not_blank"` // 备案编号，手(账)册编号

	// 清单信息
	DclTypecd      *string `json:"dclTypecd" validate:"required,oneof=1 2 3"` // 申报类型(1-备案申请 2-变更申请 3-删除申请)
	InvtType       *string `json:"invtType"`                                  // 清单类型代码
	BondInvtTypecd *string `json:"bondInvtTypecd"`                            // 清单类型代码
	ListStat       *string `json:"listStat"`                                  // 清单状态
	ListType       *string `json:"listType"`                                  // 流转类型

	// 核放相关
	InvtIochkptStucd   *string `json:"invtIochkptStucd"`   // 清单进出卡口状态代码
	PassportUsedTypeCd *string `json:"passportUsedTypeCd"` // 核放单生成标志代码
	VrfdedMarkcd       *string `json:"vrfdedMarkcd"`       // 核扣标记代码(0-未核扣 1-预核扣 2-已核扣 3-已核销)

	// 经营企业
	BizopEtpsSccd *string `json:"bizopEtpsSccd" validate:"required,len=18"`  // 社会信用代码
	BizopEtpsno   *string `json:"bizopEtpsno" validate:"required,len=10"`    // 编号
	BizopEtpsNm   *string `json:"bizopEtpsNm" validate:"required,not_blank"` // 名称
	// 申报企业
	DclEtpsSccd *string `json:"dclEtpsSccd" validate:"required,len=18"`  // 社会信用代码
	DclEtpsno   *string `json:"dclEtpsno" validate:"required,len=10"`    // 编号
	DclEtpsNm   *string `json:"dclEtpsNm" validate:"required,not_blank"` // 名称
	// 收发货企业
	RvsngdEtpsSccd *string `json:"rvsngdEtpsSccd" validate:"required,len=18"` // 社会信用代码
	RcvgdEtpsno    *string `json:"rcvgdEtpsno" validate:"required,len=10"`    // 编号
	RcvgdEtpsNm    *string `json:"rcvgdEtpsNm" validate:"required,not_blank"` // 名称
	// 录入单位
	InputCreditCode *string `json:"inputCreditCode" validate:"required,len=18"` // 社会信用代码
	InputCode       *string `json:"inputCode" validate:"required,len=10"`       // 编号
	InputName       *string `json:"inputName" validate:"required,not_blank"`    // 名称

	// 日期
	InputTime        *string `json:"inputTime" validate:"required,len=8"`   // 录入日期(格式：20240101)
	InvtDclTime      *string `json:"invtDclTime" validate:"required,len=8"` // 清单申报时间(格式：20240101)
	EntryDclTime     *string `json:"entryDclTime"`                          // 报关单申报时间(格式：20240101)
	PrevdTime        *string `json:"prevdTime"`                             // 预核扣时间
	FormalVrfdedTime *string `json:"formalVrfdedTime"`                      // 正式核扣时间

	// 代码
	ImpexpMarkcd      *string `json:"impexpMarkcd" validate:"required,oneof=I E"`          // 进出口标记代码(I-进口 E-出口)
	MtpckEndprdMarkcd *string `json:"mtpckEndprdMarkcd" validate:"required,oneof=I E"`     // 料件成品标记代码(I-料件 E-成品)
	SupvModecd        *string `json:"supvModecd" validate:"required,len=4,numeric"`        // 监管方式代码
	TrspModecd        *string `json:"trspModecd" validate:"required,len=1,numeric"`        // 运输方式代码
	ImpexpPortcd      *string `json:"impexpPortcd" validate:"required,len=4,numeric"`      // 进出境关别代码(进出口口岸代码)
	DclPlcCuscd       *string `json:"dclPlcCuscd" validate:"required,len=4,numeric"`       // 主管海关(申报地关区代码)
	StshipTrsarvNatcd *string `json:"stshipTrsarvNatcd" validate:"required,len=3,numeric"` // 起运/运抵国(地区)

	// 报关相关
	DclcusFlag        *string `json:"dclcusFlag" validate:"required,oneof=1 2"`   // 是否报关标志(1-报关 2-非报关)
	GenDecFlag        *string `json:"genDecFlag" validate:"required,oneof=1 2"`   // 是否生成报关单:1-生成 2-不生成
	DecType           *string `json:"decType" validate:"required"`                // 报关单类型
	DclcusTypecd      *string `json:"dclcusTypecd" validate:"required,oneof=1 2"` // 报关类型代码(1-关联报关 2-对应报关)
	NeedEntryModified *string `json:"needEntryModified"`                          // 报关单同步修改标志
	EntryStucd        *string `json:"entryStucd"`                                 // 报关状态(0：未放行，1：已放行) 该类型清单满足两个条件才能核扣：报关单被放行+货物全部过卡

	// 对应报关相关
	EntryNo *string `json:"entryNo"` // 对应报关单编号
	// 对应报关单申报单位
	CorrEntryDclEtpsSccd *string `json:"corrEntryDclEtpsSccd" validate:"required_if=DclcusTypecd 2,omitempty,len=18"`  // 社会信用代码
	CorrEntryDclEtpsNo   *string `json:"corrEntryDclEtpsNo" validate:"required_if=DclcusTypecd 2,omitempty,len=10"`    // 编号
	CorrEntryDclEtpsNm   *string `json:"corrEntryDclEtpsNm" validate:"required_if=DclcusTypecd 2,omitempty,not_blank"` // 名称

	// 关联报关相关
	RltInvtNo   *string `json:"rltInvtNo"`   // 关联清单编号
	RltPutrecNo *string `json:"rltPutrecNo"` // 关联备案编号，关联手(账)册编号
	RltEntryNo  *string `json:"rltEntryNo"`  // 关联报关单编号
	// 关联报关单境内收发货人
	RltEntryBizopEtpsSccd *string `json:"rltEntryBizopEtpsSccd" validate:"required_if=DclcusTypecd 1,omitempty,len=18"`  // 社会信用代码
	RltEntryBizopEtpsno   *string `json:"rltEntryBizopEtpsno" validate:"required_if=DclcusTypecd 1,omitempty,len=10"`    // 编号
	RltEntryBizopEtpsNm   *string `json:"rltEntryBizopEtpsNm" validate:"required_if=DclcusTypecd 1,omitempty,not_blank"` // 名称
	// 关联报关单生产销售(消费使用)单位
	RltEntryRvsngdEtpsSccd *string `json:"rltEntryRvsngdEtpsSccd" validate:"required_if=DclcusTypecd 1,omitempty,len=18"` // 社会信用代码
	RltEntryRcvgdEtpsno    *string `json:"rltEntryRcvgdEtpsno" validate:"required_if=DclcusTypecd 1,omitempty,len=10"`    // 编号
	RltEntryRcvgdEtpsNm    *string `json:"rltEntryRcvgdEtpsNm" validate:"required_if=DclcusTypecd 1,omitempty,not_blank"` // 名称
	// 关联报关单申报单位
	RltEntryDclEtpsSccd *string `json:"rltEntryDclEtpsSccd" validate:"required_if=DclcusTypecd 1,omitempty,len=18"`  // 社会信用代码
	RltEntryDclEtpsno   *string `json:"rltEntryDclEtpsno" validate:"required_if=DclcusTypecd 1,omitempty,len=10"`    // 编号
	RltEntryDclEtpsNm   *string `json:"rltEntryDclEtpsNm" validate:"required_if=DclcusTypecd 1,omitempty,not_blank"` // 名称

	// others
	IcCardNo  *string `json:"icCardNo"`  // 申报人IC卡号
	ApplyNo   *string `json:"applyNo"`   // 申请编号
	Rmk       *string `json:"rmk"`       // 备注
	LevyBlAmt *string `json:"levyBlAmt"` // 计征金额
}

// Inv101List 核注清单表体
type Inv101List struct {
	SeqNo         *string `json:"seqNo" validate:"omitempty,len=18"`         // 中心统一编号
	GdsSeqno      *string `json:"gdsSeqno" validate:"required,numeric"`      // 商品序号
	EntryGdsSeqno *string `json:"entryGdsSeqno" validate:"required,numeric"` // 报关单商品序号
	PutrecSeqno   *string `json:"putrecSeqno" validate:"omitempty,numeric"`  // 备案序号(对应底账序号）

	Gdecd            *string `json:"gdecd" validate:"required,numeric"`    // 商品编码
	GdsMtno          *string `json:"gdsMtno" validate:"required"`          // 商品料号
	GdsNm            *string `json:"gdsNm" validate:"required"`            // 商品名称
	GdsSpcfModelDesc *string `json:"gdsSpcfModelDesc" validate:"required"` // 商品规格型号

	LawfUnitcd     *string             `json:"lawfUnitcd" validate:"required,len=3,numeric"`                         // 法定计量单位代码
	LawfQty        *models.DecimalSafe `json:"lawfQty" validate:"required,d_gt=0"`                                   // 法定数量
	SecdLawfUnitcd *string             `json:"secdLawfUnitcd" validate:"omitempty,len=3,numeric"`                    // 第二法定计量单位代码
	SecdLawfQty    *models.DecimalSafe `json:"secdLawfQty" validate:"required_with=SecdLawfUnitcd,omitempty,d_gt=0"` // 第二法定数量

	DclUnitcd       *string             `json:"dclUnitcd" validate:"required,len=3,numeric"`                // 申报计量单位代码
	DclQty          *models.DecimalSafe `json:"dclQty" validate:"required,d_gt=0"`                          // 申报数量
	DclCurrcd       *string             `json:"dclCurrcd" validate:"required,len=3,numeric"`                // 申报币制代码
	DclUprcAmt      *models.DecimalSafe `json:"dclUprcAmt" validate:"required,d_gt=0,d_decimal_len_lte=4"`  // 企业申报单价
	DclTotalAmt     *models.DecimalSafe `json:"dclTotalAmt" validate:"required,d_gt=0,d_decimal_len_lte=2"` // 企业申报总价
	UsdStatTotalAmt *models.DecimalSafe `json:"usdStatTotalAmt"`                                            // 美元统计总金额

	Natcd            *string `json:"natcd" validate:"required,len=3,numeric"`                  // 原产国（地区）代码
	DestinationNatcd *string `json:"destinationNatcd" validate:"required,len=3,numeric"`       // 最终目的国（地区）代码
	LvyrlfModecd     *string `json:"lvyrlfModecd" validate:"required,oneof=1 2 3 4 5 6 7 8 9"` // 征免方式代码
	ClyMarkcd        *string `json:"clyMarkcd" validate:"required,oneof=0 1"`                  // 危化品标志(0-否 1-是)

	WtSfVal      *models.DecimalSafe `json:"wtSfVal"`      // 重量比例因子值
	FstSfVal     *models.DecimalSafe `json:"fstSfVal"`     // 第一比例因子值
	SecdSfVal    *models.DecimalSafe `json:"secdSfVal"`    // 第二比例因子值
	GrossWt      *models.DecimalSafe `json:"grossWt"`      // 毛重
	NetWt        *models.DecimalSafe `json:"netWt"`        // 净重
	UcnsVerno    *string             `json:"ucnsVerno"`    // 单耗版本号
	UseCd        *string             `json:"useCd"`        // 重点商品标识：0-非重点商品、1-目录重点商品、2-连带重点商品、3-过渡期重点商品
	ApplyTbSeqno *string             `json:"applyTbSeqno"` // 流转申报表序号

	ModfMarkcd *string `json:"modfMarkcd" validate:"required,oneof=0 1 2 3"` // 修改标志(0-未修改 1-修改 2-删除 3-增加)
	Rmk        *string `json:"rmk"`                                          // 备注
}

type Inv101 struct {
	Head Inv101Head   `json:"head"`
	List []Inv101List `json:"list"`
}

type Inv101Xml struct {
	XMLName xml.Name `xml:"Signature"`
	Object  struct {
		Package struct {
			EnvelopInfo commonmodels.EnvelopInfo
			DataInfo    struct {
				BusinessData struct {
					DeclareFlag string `xml:"DelcareFlag"` // 这里xml的key和struct的key不一致，海关那个应该是拼写错误
					InvtMessage struct {
						SysId          string
						OperCusRegCode string
						InvtHeadType   Inv101Head
						InvtListType   []Inv101List
					}
				} `xml:"BussinessData"`
			}
		}
	}
}
